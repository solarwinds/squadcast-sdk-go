package squadcastsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/solarwinds/squadcast-sdk-go/models/components"
)

// tokenResponse is the shape returned by the Squadcast token-refresh endpoint.
type tokenResponse struct {
	Data struct {
		AccessToken string `json:"access_token"`
		// ExpiresAt is seconds-since-epoch; fall back to a fixed TTL if absent.
		ExpiresAt int64 `json:"expires_at"`
	} `json:"data"`
}

// tokenTTL is the fallback validity window when the server does not return
// an explicit expiry.
const tokenTTL = 55 * time.Minute

// tokenRefreshTolerance is how early we treat the token as expired so we
// never ship a token that is about to expire mid-flight.
const tokenRefreshTolerance = 5 * time.Minute

// defaultRefreshURL is the Squadcast endpoint that exchanges a refresh token
// for a short-lived bearer token.
const defaultRefreshURL = "https://auth.squadcast.com/oauth/access-token"

// ──────────────────────────────────────────────────────────────
// TokenStore – pluggable persistence layer for the bearer token.
// ──────────────────────────────────────────────────────────────

// TokenStore is an interface that lets callers bring their own storage (e.g.
// Redis, Vault, an encrypted file) for the cached bearer token.  The default
// implementation keeps the token in-process.
type TokenStore interface {
	// Get returns the cached token and its expiry timestamp, or nil if not set.
	Get(ctx context.Context) (*CachedToken, error)
	// Set persists a new token together with its expiry timestamp.
	Set(ctx context.Context, token string, expires time.Time) error
}

// CachedToken holds the bearer token and the point in time after which it
// must be refreshed.
type CachedToken struct {
	Token   string
	Expires time.Time
}

// InMemoryTokenStore is a goroutine-safe, in-process TokenStore.
// It is the default used by WithRefreshToken.
type InMemoryTokenStore struct {
	mu    sync.RWMutex
	entry *CachedToken
}

func (s *InMemoryTokenStore) Get(_ context.Context) (*CachedToken, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.entry, nil
}

func (s *InMemoryTokenStore) Set(_ context.Context, token string, expires time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entry = &CachedToken{Token: token, Expires: expires}
	return nil
}

// ──────────────────────────────────────────────────────────────
// RefreshTokenSourceOptions – optional tunables.
// ──────────────────────────────────────────────────────────────

// RefreshTokenSourceOptions fine-tunes the behaviour of the token source.
// All fields are optional; sensible defaults are applied automatically.
type RefreshTokenSourceOptions struct {
	// TokenStore overrides the default in-memory store (e.g. for distributed
	// deployments that share a cache across SDK instances).
	TokenStore TokenStore
	// RefreshURL overrides the default Squadcast token-refresh endpoint.
	RefreshURL string
	// HTTPClient overrides the HTTP client used for the refresh call.
	HTTPClient *http.Client
}

// ──────────────────────────────────────────────────────────────
// refreshTokenSource – the core token-refresh machine.
// ──────────────────────────────────────────────────────────────

type refreshTokenSource struct {
	refreshToken string
	refreshURL   string
	store        TokenStore
	httpClient   *http.Client
}

// token returns a valid bearer token, fetching a new one when the cached
// entry has expired (or does not exist yet).
func (r *refreshTokenSource) token(ctx context.Context) (string, error) {
	cached, err := r.store.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("squadcastsdk: token store get: %w", err)
	}

	// Return the cached token if it is still valid.
	if cached != nil && time.Now().Before(cached.Expires) {
		return cached.Token, nil
	}

	return r.fetchAndStore(ctx)
}

// fetchAndStore calls the Squadcast refresh endpoint, parses the response and
// persists the new bearer token in the store.
func (r *refreshTokenSource) fetchAndStore(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.refreshURL, nil)
	if err != nil {
		return "", fmt.Errorf("squadcastsdk: build refresh request: %w", err)
	}

	// Squadcast uses the refresh token as the Bearer credential on this
	// endpoint to issue a short-lived access token.
	req.Header.Set("X-Refresh-Token", r.refreshToken)
	req.Header.Set("Accept", "application/json")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("squadcastsdk: refresh request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return "", fmt.Errorf("squadcastsdk: unexpected status %d from token endpoint: %s", resp.StatusCode, body)
	}

	var tr tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", fmt.Errorf("squadcastsdk: decode token response: %w", err)
	}

	if tr.Data.AccessToken == "" {
		return "", fmt.Errorf("squadcastsdk: empty access_token in refresh response")
	}

	var expires time.Time
	if tr.Data.ExpiresAt > 0 {
		expires = time.Unix(tr.Data.ExpiresAt, 0).Add(-tokenRefreshTolerance)
	} else {
		expires = time.Now().Add(tokenTTL)
	}

	if err := r.store.Set(ctx, tr.Data.AccessToken, expires); err != nil {
		return "", fmt.Errorf("squadcastsdk: token store set: %w", err)
	}

	return tr.Data.AccessToken, nil
}

// securitySource adapts the token source to the signature expected by
// WithSecuritySource.
func (r *refreshTokenSource) securitySource() func(context.Context) (components.Security, error) {
	return func(ctx context.Context) (components.Security, error) {
		token, err := r.token(ctx)
		if err != nil {
			return components.Security{}, err
		}
		return components.Security{BearerAuth: &token}, nil
	}
}

// ──────────────────────────────────────────────────────────────
// WithRefreshToken – public SDK option.
// ──────────────────────────────────────────────────────────────

// WithRefreshToken configures the SDK to authenticate using a refresh token.
//
// The SDK will transparently exchange the refresh token for a short-lived
// bearer token before the first request, cache it, and re-fetch it whenever
// it is about to expire.  The caller never needs to manage bearer tokens
// directly.
//
// Usage:
//
//	sdk := squadcastsdk.New(
//	    squadcastsdk.WithRefreshToken("your-refresh-token"),
//	)
//
// Optionally, supply custom options:
//
//	sdk := squadcastsdk.New(
//	    squadcastsdk.WithRefreshToken("your-refresh-token",
//	        squadcastsdk.RefreshTokenSourceOptions{
//	            TokenStore: myRedisStore,   // distributed cache
//	            RefreshURL: "https://...",  // custom endpoint
//	        },
//	    ),
//	)
func WithRefreshToken(refreshToken string, opts ...RefreshTokenSourceOptions) SDKOption {
	o := RefreshTokenSourceOptions{}
	if len(opts) > 0 {
		o = opts[0]
	}

	if o.TokenStore == nil {
		o.TokenStore = &InMemoryTokenStore{}
	}
	if o.RefreshURL == "" {
		o.RefreshURL = defaultRefreshURL
	}
	if o.HTTPClient == nil {
		o.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	}

	src := &refreshTokenSource{
		refreshToken: refreshToken,
		refreshURL:   o.RefreshURL,
		store:        o.TokenStore,
		httpClient:   o.HTTPClient,
	}

	return WithSecuritySource(src.securitySource())
}

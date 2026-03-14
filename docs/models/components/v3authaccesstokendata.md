# V3AuthAccessTokenData

Access token response returned by the OAuth endpoint.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AccessToken`                                                | `string`                                                     | :heavy_check_mark:                                           | JWT access token used as Bearer token for API requests.      |
| `ExpiresAt`                                                  | `int64`                                                      | :heavy_check_mark:                                           | Unix timestamp when the access token expires.                |
| `IssuedAt`                                                   | `int64`                                                      | :heavy_check_mark:                                           | Unix timestamp when the access token was issued.             |
| `RefreshToken`                                               | `string`                                                     | :heavy_check_mark:                                           | Refresh token that can be used to obtain a new access token. |
| `Type`                                                       | `string`                                                     | :heavy_check_mark:                                           | Token type, e.g. "Bearer".                                   |
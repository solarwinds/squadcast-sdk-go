<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	squadcastsdk "github.com/solarwinds/squadcast-sdk-go"
	"log"
)

func main() {
	ctx := context.Background()

	s := squadcastsdk.New()

	res, err := s.Auth.AuthGetAccessToken(ctx, "<value>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->
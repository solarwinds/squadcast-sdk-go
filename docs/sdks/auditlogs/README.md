# AuditLogs
(*AuditLogs*)

## Overview

### Available Operations

* [List](#list) - List all Audit Logs
* [Export](#export) - Initiate an asynchronous export of audit logs based on the provided filters. The export file will be generated and available for download. Use 'Get details of Audit Logs export history by ID' API to retrieve the download URL.
* [ListExportHistory](#listexporthistory) - List all Audit Logs export history
* [GetExportHistoryByID](#getexporthistorybyid) - Get details of Audit Logs export history by ID
* [GetByID](#getbyid) - Get audit log by ID

## List

List all Audit Logs
Returns array of audit logs for given team and filters

### Example Usage

<!-- UsageSnippet language="go" operationID="AuditLogs_listAuditLogs" method="get" path="/v3/audit-logs" -->
```go
package main

import(
	"context"
	"os"
	squadcastsdk "github.com/solarwinds/squadcast-sdk-go"
	"github.com/solarwinds/squadcast-sdk-go/types"
	"github.com/solarwinds/squadcast-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := squadcastsdk.New(
        squadcastsdk.WithSecurity(os.Getenv("SQUADCASTSDK_BEARER_AUTH")),
    )

    res, err := s.AuditLogs.List(ctx, operations.AuditLogsListAuditLogsRequest{
        PageSize: 832442,
        PageNumber: 555332,
        StartDate: types.MustDateFromString("2023-03-04"),
        EndDate: types.MustDateFromString("2024-08-07"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3AuditLogsListAuditLogsResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AuditLogsListAuditLogsRequest](../../models/operations/auditlogslistauditlogsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AuditLogsListAuditLogsResponse](../../models/operations/auditlogslistauditlogsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.BadRequestError          | 400                                | application/json                   |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.PaymentRequiredError     | 402                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.ConflictError            | 409                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.BadGatewayError          | 502                                | application/json                   |
| apierrors.ServiceUnavailableError  | 503                                | application/json                   |
| apierrors.GatewayTimeoutError      | 504                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## Export

Export Audit logs
Initiates export of audit logs based on provided filters

### Example Usage

<!-- UsageSnippet language="go" operationID="AuditLogs_exportAuditLogs" method="post" path="/v3/audit-logs/export" -->
```go
package main

import(
	"context"
	"os"
	squadcastsdk "github.com/solarwinds/squadcast-sdk-go"
	"github.com/solarwinds/squadcast-sdk-go/types"
	"github.com/solarwinds/squadcast-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := squadcastsdk.New(
        squadcastsdk.WithSecurity(os.Getenv("SQUADCASTSDK_BEARER_AUTH")),
    )

    res, err := s.AuditLogs.Export(ctx, components.V3AuditLogsExportAuditLogsRequest{
        Filters: components.Filters{
            StartDate: types.MustDateFromString("2025-07-29"),
            EndDate: types.MustDateFromString("2023-09-09"),
        },
        Name: "<value>",
        ExportType: components.ExportTypeJSON,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3AuditLogsExportAuditLogsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [components.V3AuditLogsExportAuditLogsRequest](../../models/components/v3auditlogsexportauditlogsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AuditLogsExportAuditLogsResponse](../../models/operations/auditlogsexportauditlogsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.BadRequestError          | 400                                | application/json                   |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.PaymentRequiredError     | 402                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.ConflictError            | 409                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.BadGatewayError          | 502                                | application/json                   |
| apierrors.ServiceUnavailableError  | 503                                | application/json                   |
| apierrors.GatewayTimeoutError      | 504                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## ListExportHistory

List all Audit Logs export history
Returns array of audit logs export history

### Example Usage

<!-- UsageSnippet language="go" operationID="AuditLogs_listAuditLogsExportHistory" method="get" path="/v3/audit-logs/export/history" -->
```go
package main

import(
	"context"
	"os"
	squadcastsdk "github.com/solarwinds/squadcast-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := squadcastsdk.New(
        squadcastsdk.WithSecurity(os.Getenv("SQUADCASTSDK_BEARER_AUTH")),
    )

    res, err := s.AuditLogs.ListExportHistory(ctx, 159672, 351281)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3AuditLogsListAuditLogsExportHistoryResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `pageSize`                                               | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `pageNumber`                                             | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AuditLogsListAuditLogsExportHistoryResponse](../../models/operations/auditlogslistauditlogsexporthistoryresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.BadRequestError          | 400                                | application/json                   |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.PaymentRequiredError     | 402                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.ConflictError            | 409                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.BadGatewayError          | 502                                | application/json                   |
| apierrors.ServiceUnavailableError  | 503                                | application/json                   |
| apierrors.GatewayTimeoutError      | 504                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## GetExportHistoryByID

Get details of Audit Logs export history by ID
Returns audit log export history details for the specified ID

### Example Usage

<!-- UsageSnippet language="go" operationID="AuditLogs_getAuditLogsExportHistoryById" method="get" path="/v3/audit-logs/export/history/{id}" -->
```go
package main

import(
	"context"
	"os"
	squadcastsdk "github.com/solarwinds/squadcast-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := squadcastsdk.New(
        squadcastsdk.WithSecurity(os.Getenv("SQUADCASTSDK_BEARER_AUTH")),
    )

    res, err := s.AuditLogs.GetExportHistoryByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.V3AuditLogsGetAuditLogExportHistoryByIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AuditLogsGetAuditLogsExportHistoryByIDResponse](../../models/operations/auditlogsgetauditlogsexporthistorybyidresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.BadRequestError          | 400                                | application/json                   |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.PaymentRequiredError     | 402                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.ConflictError            | 409                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.BadGatewayError          | 502                                | application/json                   |
| apierrors.ServiceUnavailableError  | 503                                | application/json                   |
| apierrors.GatewayTimeoutError      | 504                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## GetByID

Get audit log by ID
Returns audit log details for the specified ID

### Example Usage

<!-- UsageSnippet language="go" operationID="AuditLogs_getAuditLogById" method="get" path="/v3/audit-logs/{id}" -->
```go
package main

import(
	"context"
	"os"
	squadcastsdk "github.com/solarwinds/squadcast-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := squadcastsdk.New(
        squadcastsdk.WithSecurity(os.Getenv("SQUADCASTSDK_BEARER_AUTH")),
    )

    res, err := s.AuditLogs.GetByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.V3AuditLogsGetAuditLogByIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AuditLogsGetAuditLogByIDResponse](../../models/operations/auditlogsgetauditlogbyidresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.BadRequestError          | 400                                | application/json                   |
| apierrors.UnauthorizedError        | 401                                | application/json                   |
| apierrors.PaymentRequiredError     | 402                                | application/json                   |
| apierrors.ForbiddenError           | 403                                | application/json                   |
| apierrors.NotFoundError            | 404                                | application/json                   |
| apierrors.ConflictError            | 409                                | application/json                   |
| apierrors.UnprocessableEntityError | 422                                | application/json                   |
| apierrors.InternalServerError      | 500                                | application/json                   |
| apierrors.BadGatewayError          | 502                                | application/json                   |
| apierrors.ServiceUnavailableError  | 503                                | application/json                   |
| apierrors.GatewayTimeoutError      | 504                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |
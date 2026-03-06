# Workflows.Actions

## Overview

### Available Operations

* [Create](#create) - Create Action

## Create

Create an Action for a workflow

### Example Usage

<!-- UsageSnippet language="go" operationID="Workflows_createAction" method="post" path="/v3/workflows/{workflowID}/actions" -->
```go
package main

import(
	"context"
	"os"
	squadcastsdk "github.com/solarwinds/squadcast-sdk-go"
	"github.com/solarwinds/squadcast-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := squadcastsdk.New(
        squadcastsdk.WithSecurity(os.Getenv("SQUADCASTSDK_REFRESH_TOKEN_AUTH")),
    )

    res, err := s.Workflows.Actions.Create(ctx, "<id>", components.CreateV3WorkflowsActionRequestV3WorkflowsSqCreateStatusPageIssue(
        components.V3WorkflowsSqCreateStatusPageIssue{
            Name: components.V3WorkflowsSqCreateStatusPageIssueNameSqAddStatusPageIssue,
            Data: components.V3WorkflowsSqCreateStatusPageIssueData{
                ComponentAndImpact: []components.V3WorkflowsComponentAndImpact{},
                IssueTitle: "<value>",
                PageStatusID: 179034,
                StatusAndMessage: []components.V3WorkflowsIssueStatusAndMessage{
                    components.V3WorkflowsIssueStatusAndMessage{
                        Messages: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                        StatusID: 692068,
                    },
                },
                StatusPageID: 368871,
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        switch res.Object.Data.Type {
            case components.V3WorkflowsActionResponseTypeV3WorkflowsActionResponseSqAttachRunbooks:
                // res.Object.Data.V3WorkflowsActionResponseSqAttachRunbooks is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSqMarkIncidentSLOAffecting:
                // res.Object.Data.V3WorkflowsSqMarkIncidentSLOAffecting is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSqTriggerManualWebhook:
                // res.Object.Data.V3WorkflowsSqTriggerManualWebhook is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsUpdateIncidentPriority:
                // res.Object.Data.V3WorkflowsUpdateIncidentPriority is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSqCreateStatusPageIssue:
                // res.Object.Data.V3WorkflowsSqCreateStatusPageIssue is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSqAddIncidentNote:
                // res.Object.Data.V3WorkflowsSqAddIncidentNote is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSlackArchiveChannel:
                // res.Object.Data.V3WorkflowsSlackArchiveChannel is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSqAddCommunicationChannel:
                // res.Object.Data.V3WorkflowsSqAddCommunicationChannel is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSlackMessageChannel:
                // res.Object.Data.V3WorkflowsSlackMessageChannel is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSlackMessageUser:
                // res.Object.Data.V3WorkflowsSlackMessageUser is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSqMakeHTTPCall:
                // res.Object.Data.V3WorkflowsSqMakeHTTPCall is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSlackCreateIncidentChannel:
                // res.Object.Data.V3WorkflowsSlackCreateIncidentChannel is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsJiraCreateTicket:
                // res.Object.Data.V3WorkflowsJiraCreateTicket is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsMsTeamsMessageChannel:
                // res.Object.Data.V3WorkflowsMsTeamsMessageChannel is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsMsTeamsMessageUser:
                // res.Object.Data.V3WorkflowsMsTeamsMessageUser is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsSqSendEmail:
                // res.Object.Data.V3WorkflowsSqSendEmail is populated
            case components.V3WorkflowsActionResponseTypeV3WorkflowsMsTeamsCreateMeetingLink:
                // res.Object.Data.V3WorkflowsMsTeamsCreateMeetingLink is populated
            case components.V3WorkflowsActionResponseTypeAny:
                // res.Object.Data.Any is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `workflowID`                                                                               | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `v3WorkflowsActionRequest`                                                                 | [components.V3WorkflowsActionRequest](../../models/components/v3workflowsactionrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.WorkflowsCreateActionResponse](../../models/operations/workflowscreateactionresponse.md), error**

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
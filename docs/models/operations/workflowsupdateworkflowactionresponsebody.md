# WorkflowsUpdateWorkflowActionResponseBody

The request has succeeded.


## Supported Types

### SqAttachRunbooks

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodySqAttachRunbooks(operations.SqAttachRunbooks{/* values here */})
```

### V3WorkflowsSqMarkIncidentSLOAffecting

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSqMarkIncidentSLOAffecting(components.V3WorkflowsSqMarkIncidentSLOAffecting{/* values here */})
```

### V3WorkflowsSqTriggerManualWebhook

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSqTriggerManualWebhook(components.V3WorkflowsSqTriggerManualWebhook{/* values here */})
```

### V3WorkflowsUpdateIncidentPriority

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsUpdateIncidentPriority(components.V3WorkflowsUpdateIncidentPriority{/* values here */})
```

### V3WorkflowsSqCreateStatusPageIssue

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSqCreateStatusPageIssue(components.V3WorkflowsSqCreateStatusPageIssue{/* values here */})
```

### V3WorkflowsSqAddIncidentNote

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSqAddIncidentNote(components.V3WorkflowsSqAddIncidentNote{/* values here */})
```

### V3WorkflowsSlackArchiveChannel

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSlackArchiveChannel(components.V3WorkflowsSlackArchiveChannel{/* values here */})
```

### V3WorkflowsSqAddCommunicationChannel

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSqAddCommunicationChannel(components.V3WorkflowsSqAddCommunicationChannel{/* values here */})
```

### V3WorkflowsSlackMessageChannel

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSlackMessageChannel(components.V3WorkflowsSlackMessageChannel{/* values here */})
```

### V3WorkflowsSlackMessageUser

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSlackMessageUser(components.V3WorkflowsSlackMessageUser{/* values here */})
```

### V3WorkflowsSqMakeHTTPCall

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSqMakeHTTPCall(components.V3WorkflowsSqMakeHTTPCall{/* values here */})
```

### V3WorkflowsSlackCreateIncidentChannel

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSlackCreateIncidentChannel(components.V3WorkflowsSlackCreateIncidentChannel{/* values here */})
```

### V3WorkflowsJiraCreateTicket

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsJiraCreateTicket(components.V3WorkflowsJiraCreateTicket{/* values here */})
```

### V3WorkflowsMsTeamsMessageChannel

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsMsTeamsMessageChannel(components.V3WorkflowsMsTeamsMessageChannel{/* values here */})
```

### V3WorkflowsMsTeamsMessageUser

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsMsTeamsMessageUser(components.V3WorkflowsMsTeamsMessageUser{/* values here */})
```

### V3WorkflowsSqSendEmail

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsSqSendEmail(components.V3WorkflowsSqSendEmail{/* values here */})
```

### V3WorkflowsMsTeamsCreateMeetingLink

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyV3WorkflowsMsTeamsCreateMeetingLink(components.V3WorkflowsMsTeamsCreateMeetingLink{/* values here */})
```

### 

```go
workflowsUpdateWorkflowActionResponseBody := operations.CreateWorkflowsUpdateWorkflowActionResponseBodyAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch workflowsUpdateWorkflowActionResponseBody.Type {
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeSqAttachRunbooks:
		// workflowsUpdateWorkflowActionResponseBody.SqAttachRunbooks is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSqMarkIncidentSLOAffecting:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSqMarkIncidentSLOAffecting is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSqTriggerManualWebhook:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSqTriggerManualWebhook is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsUpdateIncidentPriority:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsUpdateIncidentPriority is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSqCreateStatusPageIssue:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSqCreateStatusPageIssue is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSqAddIncidentNote:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSqAddIncidentNote is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSlackArchiveChannel:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSlackArchiveChannel is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSqAddCommunicationChannel:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSqAddCommunicationChannel is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSlackMessageChannel:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSlackMessageChannel is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSlackMessageUser:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSlackMessageUser is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSqMakeHTTPCall:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSqMakeHTTPCall is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSlackCreateIncidentChannel:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSlackCreateIncidentChannel is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsJiraCreateTicket:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsJiraCreateTicket is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsMsTeamsMessageChannel:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsMsTeamsMessageChannel is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsMsTeamsMessageUser:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsMsTeamsMessageUser is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsSqSendEmail:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsSqSendEmail is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeV3WorkflowsMsTeamsCreateMeetingLink:
		// workflowsUpdateWorkflowActionResponseBody.V3WorkflowsMsTeamsCreateMeetingLink is populated
	case operations.WorkflowsUpdateWorkflowActionResponseBodyTypeAny:
		// workflowsUpdateWorkflowActionResponseBody.Any is populated
}
```

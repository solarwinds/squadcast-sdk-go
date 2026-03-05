# V3WorkflowsActionRequest


## Supported Types

### V3WorkflowsActionRequestSqAttachRunbooks

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsActionRequestSqAttachRunbooks(components.V3WorkflowsActionRequestSqAttachRunbooks{/* values here */})
```

### V3WorkflowsSqMarkIncidentSLOAffecting

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSqMarkIncidentSLOAffecting(components.V3WorkflowsSqMarkIncidentSLOAffecting{/* values here */})
```

### V3WorkflowsSqTriggerManualWebhook

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSqTriggerManualWebhook(components.V3WorkflowsSqTriggerManualWebhook{/* values here */})
```

### V3WorkflowsUpdateIncidentPriority

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsUpdateIncidentPriority(components.V3WorkflowsUpdateIncidentPriority{/* values here */})
```

### V3WorkflowsSqCreateStatusPageIssue

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSqCreateStatusPageIssue(components.V3WorkflowsSqCreateStatusPageIssue{/* values here */})
```

### V3WorkflowsSqAddIncidentNote

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSqAddIncidentNote(components.V3WorkflowsSqAddIncidentNote{/* values here */})
```

### V3WorkflowsSlackArchiveChannel

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSlackArchiveChannel(components.V3WorkflowsSlackArchiveChannel{/* values here */})
```

### V3WorkflowsSqAddCommunicationChannel

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSqAddCommunicationChannel(components.V3WorkflowsSqAddCommunicationChannel{/* values here */})
```

### V3WorkflowsSlackMessageChannel

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSlackMessageChannel(components.V3WorkflowsSlackMessageChannel{/* values here */})
```

### V3WorkflowsSlackMessageUser

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSlackMessageUser(components.V3WorkflowsSlackMessageUser{/* values here */})
```

### V3WorkflowsSqMakeHTTPCall

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSqMakeHTTPCall(components.V3WorkflowsSqMakeHTTPCall{/* values here */})
```

### V3WorkflowsSlackCreateIncidentChannel

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSlackCreateIncidentChannel(components.V3WorkflowsSlackCreateIncidentChannel{/* values here */})
```

### V3WorkflowsJiraCreateTicket

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsJiraCreateTicket(components.V3WorkflowsJiraCreateTicket{/* values here */})
```

### V3WorkflowsMsTeamsMessageChannel

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsMsTeamsMessageChannel(components.V3WorkflowsMsTeamsMessageChannel{/* values here */})
```

### V3WorkflowsMsTeamsMessageUser

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsMsTeamsMessageUser(components.V3WorkflowsMsTeamsMessageUser{/* values here */})
```

### V3WorkflowsSqSendEmail

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsSqSendEmail(components.V3WorkflowsSqSendEmail{/* values here */})
```

### V3WorkflowsMsTeamsCreateMeetingLink

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestV3WorkflowsMsTeamsCreateMeetingLink(components.V3WorkflowsMsTeamsCreateMeetingLink{/* values here */})
```

### 

```go
v3WorkflowsActionRequest := components.CreateV3WorkflowsActionRequestAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3WorkflowsActionRequest.Type {
	case components.V3WorkflowsActionRequestTypeV3WorkflowsActionRequestSqAttachRunbooks:
		// v3WorkflowsActionRequest.V3WorkflowsActionRequestSqAttachRunbooks is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSqMarkIncidentSLOAffecting:
		// v3WorkflowsActionRequest.V3WorkflowsSqMarkIncidentSLOAffecting is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSqTriggerManualWebhook:
		// v3WorkflowsActionRequest.V3WorkflowsSqTriggerManualWebhook is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsUpdateIncidentPriority:
		// v3WorkflowsActionRequest.V3WorkflowsUpdateIncidentPriority is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSqCreateStatusPageIssue:
		// v3WorkflowsActionRequest.V3WorkflowsSqCreateStatusPageIssue is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSqAddIncidentNote:
		// v3WorkflowsActionRequest.V3WorkflowsSqAddIncidentNote is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSlackArchiveChannel:
		// v3WorkflowsActionRequest.V3WorkflowsSlackArchiveChannel is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSqAddCommunicationChannel:
		// v3WorkflowsActionRequest.V3WorkflowsSqAddCommunicationChannel is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSlackMessageChannel:
		// v3WorkflowsActionRequest.V3WorkflowsSlackMessageChannel is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSlackMessageUser:
		// v3WorkflowsActionRequest.V3WorkflowsSlackMessageUser is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSqMakeHTTPCall:
		// v3WorkflowsActionRequest.V3WorkflowsSqMakeHTTPCall is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSlackCreateIncidentChannel:
		// v3WorkflowsActionRequest.V3WorkflowsSlackCreateIncidentChannel is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsJiraCreateTicket:
		// v3WorkflowsActionRequest.V3WorkflowsJiraCreateTicket is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsMsTeamsMessageChannel:
		// v3WorkflowsActionRequest.V3WorkflowsMsTeamsMessageChannel is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsMsTeamsMessageUser:
		// v3WorkflowsActionRequest.V3WorkflowsMsTeamsMessageUser is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsSqSendEmail:
		// v3WorkflowsActionRequest.V3WorkflowsSqSendEmail is populated
	case components.V3WorkflowsActionRequestTypeV3WorkflowsMsTeamsCreateMeetingLink:
		// v3WorkflowsActionRequest.V3WorkflowsMsTeamsCreateMeetingLink is populated
	case components.V3WorkflowsActionRequestTypeAny:
		// v3WorkflowsActionRequest.Any is populated
}
```

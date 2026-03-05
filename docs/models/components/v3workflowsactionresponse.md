# V3WorkflowsActionResponse


## Supported Types

### V3WorkflowsActionResponseSqAttachRunbooks

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsActionResponseSqAttachRunbooks(components.V3WorkflowsActionResponseSqAttachRunbooks{/* values here */})
```

### V3WorkflowsSqMarkIncidentSLOAffecting

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSqMarkIncidentSLOAffecting(components.V3WorkflowsSqMarkIncidentSLOAffecting{/* values here */})
```

### V3WorkflowsSqTriggerManualWebhook

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSqTriggerManualWebhook(components.V3WorkflowsSqTriggerManualWebhook{/* values here */})
```

### V3WorkflowsUpdateIncidentPriority

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsUpdateIncidentPriority(components.V3WorkflowsUpdateIncidentPriority{/* values here */})
```

### V3WorkflowsSqCreateStatusPageIssue

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSqCreateStatusPageIssue(components.V3WorkflowsSqCreateStatusPageIssue{/* values here */})
```

### V3WorkflowsSqAddIncidentNote

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSqAddIncidentNote(components.V3WorkflowsSqAddIncidentNote{/* values here */})
```

### V3WorkflowsSlackArchiveChannel

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSlackArchiveChannel(components.V3WorkflowsSlackArchiveChannel{/* values here */})
```

### V3WorkflowsSqAddCommunicationChannel

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSqAddCommunicationChannel(components.V3WorkflowsSqAddCommunicationChannel{/* values here */})
```

### V3WorkflowsSlackMessageChannel

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSlackMessageChannel(components.V3WorkflowsSlackMessageChannel{/* values here */})
```

### V3WorkflowsSlackMessageUser

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSlackMessageUser(components.V3WorkflowsSlackMessageUser{/* values here */})
```

### V3WorkflowsSqMakeHTTPCall

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSqMakeHTTPCall(components.V3WorkflowsSqMakeHTTPCall{/* values here */})
```

### V3WorkflowsSlackCreateIncidentChannel

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSlackCreateIncidentChannel(components.V3WorkflowsSlackCreateIncidentChannel{/* values here */})
```

### V3WorkflowsJiraCreateTicket

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsJiraCreateTicket(components.V3WorkflowsJiraCreateTicket{/* values here */})
```

### V3WorkflowsMsTeamsMessageChannel

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsMsTeamsMessageChannel(components.V3WorkflowsMsTeamsMessageChannel{/* values here */})
```

### V3WorkflowsMsTeamsMessageUser

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsMsTeamsMessageUser(components.V3WorkflowsMsTeamsMessageUser{/* values here */})
```

### V3WorkflowsSqSendEmail

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsSqSendEmail(components.V3WorkflowsSqSendEmail{/* values here */})
```

### V3WorkflowsMsTeamsCreateMeetingLink

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseV3WorkflowsMsTeamsCreateMeetingLink(components.V3WorkflowsMsTeamsCreateMeetingLink{/* values here */})
```

### 

```go
v3WorkflowsActionResponse := components.CreateV3WorkflowsActionResponseAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3WorkflowsActionResponse.Type {
	case components.V3WorkflowsActionResponseTypeV3WorkflowsActionResponseSqAttachRunbooks:
		// v3WorkflowsActionResponse.V3WorkflowsActionResponseSqAttachRunbooks is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSqMarkIncidentSLOAffecting:
		// v3WorkflowsActionResponse.V3WorkflowsSqMarkIncidentSLOAffecting is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSqTriggerManualWebhook:
		// v3WorkflowsActionResponse.V3WorkflowsSqTriggerManualWebhook is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsUpdateIncidentPriority:
		// v3WorkflowsActionResponse.V3WorkflowsUpdateIncidentPriority is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSqCreateStatusPageIssue:
		// v3WorkflowsActionResponse.V3WorkflowsSqCreateStatusPageIssue is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSqAddIncidentNote:
		// v3WorkflowsActionResponse.V3WorkflowsSqAddIncidentNote is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSlackArchiveChannel:
		// v3WorkflowsActionResponse.V3WorkflowsSlackArchiveChannel is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSqAddCommunicationChannel:
		// v3WorkflowsActionResponse.V3WorkflowsSqAddCommunicationChannel is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSlackMessageChannel:
		// v3WorkflowsActionResponse.V3WorkflowsSlackMessageChannel is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSlackMessageUser:
		// v3WorkflowsActionResponse.V3WorkflowsSlackMessageUser is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSqMakeHTTPCall:
		// v3WorkflowsActionResponse.V3WorkflowsSqMakeHTTPCall is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSlackCreateIncidentChannel:
		// v3WorkflowsActionResponse.V3WorkflowsSlackCreateIncidentChannel is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsJiraCreateTicket:
		// v3WorkflowsActionResponse.V3WorkflowsJiraCreateTicket is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsMsTeamsMessageChannel:
		// v3WorkflowsActionResponse.V3WorkflowsMsTeamsMessageChannel is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsMsTeamsMessageUser:
		// v3WorkflowsActionResponse.V3WorkflowsMsTeamsMessageUser is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsSqSendEmail:
		// v3WorkflowsActionResponse.V3WorkflowsSqSendEmail is populated
	case components.V3WorkflowsActionResponseTypeV3WorkflowsMsTeamsCreateMeetingLink:
		// v3WorkflowsActionResponse.V3WorkflowsMsTeamsCreateMeetingLink is populated
	case components.V3WorkflowsActionResponseTypeAny:
		// v3WorkflowsActionResponse.Any is populated
}
```

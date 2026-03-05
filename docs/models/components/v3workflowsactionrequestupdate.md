# V3WorkflowsActionRequestUpdate


## Supported Types

### V3WorkflowsActionRequestUpdateSqAttachRunbooks

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsActionRequestUpdateSqAttachRunbooks(components.V3WorkflowsActionRequestUpdateSqAttachRunbooks{/* values here */})
```

### V3WorkflowsSqMarkIncidentSLOAffectingUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSqMarkIncidentSLOAffectingUpdate(components.V3WorkflowsSqMarkIncidentSLOAffectingUpdate{/* values here */})
```

### V3WorkflowsSqTriggerManualWebhookUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSqTriggerManualWebhookUpdate(components.V3WorkflowsSqTriggerManualWebhookUpdate{/* values here */})
```

### V3WorkflowsUpdateIncidentPriorityUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsUpdateIncidentPriorityUpdate(components.V3WorkflowsUpdateIncidentPriorityUpdate{/* values here */})
```

### V3WorkflowsSqCreateStatusPageIssueUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSqCreateStatusPageIssueUpdate(components.V3WorkflowsSqCreateStatusPageIssueUpdate{/* values here */})
```

### V3WorkflowsSqAddIncidentNoteUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSqAddIncidentNoteUpdate(components.V3WorkflowsSqAddIncidentNoteUpdate{/* values here */})
```

### V3WorkflowsSlackArchiveChannelUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSlackArchiveChannelUpdate(components.V3WorkflowsSlackArchiveChannelUpdate{/* values here */})
```

### V3WorkflowsSqAddCommunicationChannelUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSqAddCommunicationChannelUpdate(components.V3WorkflowsSqAddCommunicationChannelUpdate{/* values here */})
```

### V3WorkflowsSlackMessageChannelUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSlackMessageChannelUpdate(components.V3WorkflowsSlackMessageChannelUpdate{/* values here */})
```

### V3WorkflowsSlackMessageUserUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSlackMessageUserUpdate(components.V3WorkflowsSlackMessageUserUpdate{/* values here */})
```

### V3WorkflowsSqMakeHTTPCallUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSqMakeHTTPCallUpdate(components.V3WorkflowsSqMakeHTTPCallUpdate{/* values here */})
```

### V3WorkflowsSlackCreateIncidentChannelUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSlackCreateIncidentChannelUpdate(components.V3WorkflowsSlackCreateIncidentChannelUpdate{/* values here */})
```

### V3WorkflowsJiraCreateTicketUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsJiraCreateTicketUpdate(components.V3WorkflowsJiraCreateTicketUpdate{/* values here */})
```

### V3WorkflowsMsTeamsMessageChannelUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsMsTeamsMessageChannelUpdate(components.V3WorkflowsMsTeamsMessageChannelUpdate{/* values here */})
```

### V3WorkflowsMsTeamsMessageUserUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsMsTeamsMessageUserUpdate(components.V3WorkflowsMsTeamsMessageUserUpdate{/* values here */})
```

### V3WorkflowsSqSendEmailUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsSqSendEmailUpdate(components.V3WorkflowsSqSendEmailUpdate{/* values here */})
```

### V3WorkflowsMsTeamsCreateMeetingLinkUpdate

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateV3WorkflowsMsTeamsCreateMeetingLinkUpdate(components.V3WorkflowsMsTeamsCreateMeetingLinkUpdate{/* values here */})
```

### 

```go
v3WorkflowsActionRequestUpdate := components.CreateV3WorkflowsActionRequestUpdateAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3WorkflowsActionRequestUpdate.Type {
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsActionRequestUpdateSqAttachRunbooks:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsActionRequestUpdateSqAttachRunbooks is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSqMarkIncidentSLOAffectingUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSqMarkIncidentSLOAffectingUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSqTriggerManualWebhookUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSqTriggerManualWebhookUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsUpdateIncidentPriorityUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsUpdateIncidentPriorityUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSqCreateStatusPageIssueUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSqCreateStatusPageIssueUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSqAddIncidentNoteUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSqAddIncidentNoteUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSlackArchiveChannelUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSlackArchiveChannelUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSqAddCommunicationChannelUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSqAddCommunicationChannelUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSlackMessageChannelUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSlackMessageChannelUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSlackMessageUserUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSlackMessageUserUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSqMakeHTTPCallUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSqMakeHTTPCallUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSlackCreateIncidentChannelUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSlackCreateIncidentChannelUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsJiraCreateTicketUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsJiraCreateTicketUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsMsTeamsMessageChannelUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsMsTeamsMessageChannelUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsMsTeamsMessageUserUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsMsTeamsMessageUserUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsSqSendEmailUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsSqSendEmailUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeV3WorkflowsMsTeamsCreateMeetingLinkUpdate:
		// v3WorkflowsActionRequestUpdate.V3WorkflowsMsTeamsCreateMeetingLinkUpdate is populated
	case components.V3WorkflowsActionRequestUpdateTypeAny:
		// v3WorkflowsActionRequestUpdate.Any is populated
}
```

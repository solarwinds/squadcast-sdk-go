# IncidentsMergeIncidentsRequest


## Supported Types

### V3IncidentsMergeIntoExistingParentRequest

```go
incidentsMergeIncidentsRequest := operations.CreateIncidentsMergeIncidentsRequestV3IncidentsMergeIntoExistingParentRequest(components.V3IncidentsMergeIntoExistingParentRequest{/* values here */})
```

### V3IncidentsMergeIntoNewParentRequest

```go
incidentsMergeIncidentsRequest := operations.CreateIncidentsMergeIncidentsRequestV3IncidentsMergeIntoNewParentRequest(components.V3IncidentsMergeIntoNewParentRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch incidentsMergeIncidentsRequest.Type {
	case operations.IncidentsMergeIncidentsRequestTypeV3IncidentsMergeIntoExistingParentRequest:
		// incidentsMergeIncidentsRequest.V3IncidentsMergeIntoExistingParentRequest is populated
	case operations.IncidentsMergeIncidentsRequestTypeV3IncidentsMergeIntoNewParentRequest:
		// incidentsMergeIncidentsRequest.V3IncidentsMergeIntoNewParentRequest is populated
}
```

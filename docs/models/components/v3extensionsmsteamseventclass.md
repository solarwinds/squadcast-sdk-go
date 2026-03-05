# V3ExtensionsMSTeamsEventClass

Represents the specific type of an incident-related event.


## Supported Types

### 

```go
v3ExtensionsMSTeamsEventClass := components.CreateV3ExtensionsMSTeamsEventClassStr(string{/* values here */})
```

### V3ExtensionsMSTeamsEventClassEnum

```go
v3ExtensionsMSTeamsEventClass := components.CreateV3ExtensionsMSTeamsEventClassV3ExtensionsMSTeamsEventClassEnum(components.V3ExtensionsMSTeamsEventClassEnum{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3ExtensionsMSTeamsEventClass.Type {
	case components.V3ExtensionsMSTeamsEventClassTypeStr:
		// v3ExtensionsMSTeamsEventClass.Str is populated
	case components.V3ExtensionsMSTeamsEventClassTypeV3ExtensionsMSTeamsEventClassEnum:
		// v3ExtensionsMSTeamsEventClass.V3ExtensionsMSTeamsEventClassEnum is populated
}
```

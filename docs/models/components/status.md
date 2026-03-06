# Status


## Supported Types

### 

```go
status := components.CreateStatusStr(string{/* values here */})
```

### 

```go
status := components.CreateStatusInteger(int64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch status.Type {
	case components.StatusTypeStr:
		// status.Str is populated
	case components.StatusTypeInteger:
		// status.Integer is populated
}
```

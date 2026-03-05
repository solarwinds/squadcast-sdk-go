# V3ServicesTaggingRulesExpressionBranchRHS


## Supported Types

### 

```go
v3ServicesTaggingRulesExpressionBranchRHS := components.CreateV3ServicesTaggingRulesExpressionBranchRHSStr(string{/* values here */})
```

### 

```go
v3ServicesTaggingRulesExpressionBranchRHS := components.CreateV3ServicesTaggingRulesExpressionBranchRHSBoolean(bool{/* values here */})
```

### 

```go
v3ServicesTaggingRulesExpressionBranchRHS := components.CreateV3ServicesTaggingRulesExpressionBranchRHSInt32(int{/* values here */})
```

### 

```go
v3ServicesTaggingRulesExpressionBranchRHS := components.CreateV3ServicesTaggingRulesExpressionBranchRHSFloat32(float32{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3ServicesTaggingRulesExpressionBranchRHS.Type {
	case components.V3ServicesTaggingRulesExpressionBranchRHSTypeStr:
		// v3ServicesTaggingRulesExpressionBranchRHS.Str is populated
	case components.V3ServicesTaggingRulesExpressionBranchRHSTypeBoolean:
		// v3ServicesTaggingRulesExpressionBranchRHS.Boolean is populated
	case components.V3ServicesTaggingRulesExpressionBranchRHSTypeInt32:
		// v3ServicesTaggingRulesExpressionBranchRHS.Int32 is populated
	case components.V3ServicesTaggingRulesExpressionBranchRHSTypeFloat32:
		// v3ServicesTaggingRulesExpressionBranchRHS.Float32 is populated
}
```

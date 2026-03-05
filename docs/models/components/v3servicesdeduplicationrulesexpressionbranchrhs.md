# V3ServicesDeduplicationRulesExpressionBranchRHS


## Supported Types

### 

```go
v3ServicesDeduplicationRulesExpressionBranchRHS := components.CreateV3ServicesDeduplicationRulesExpressionBranchRHSStr(string{/* values here */})
```

### 

```go
v3ServicesDeduplicationRulesExpressionBranchRHS := components.CreateV3ServicesDeduplicationRulesExpressionBranchRHSBoolean(bool{/* values here */})
```

### 

```go
v3ServicesDeduplicationRulesExpressionBranchRHS := components.CreateV3ServicesDeduplicationRulesExpressionBranchRHSInt32(int{/* values here */})
```

### 

```go
v3ServicesDeduplicationRulesExpressionBranchRHS := components.CreateV3ServicesDeduplicationRulesExpressionBranchRHSFloat32(float32{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3ServicesDeduplicationRulesExpressionBranchRHS.Type {
	case components.V3ServicesDeduplicationRulesExpressionBranchRHSTypeStr:
		// v3ServicesDeduplicationRulesExpressionBranchRHS.Str is populated
	case components.V3ServicesDeduplicationRulesExpressionBranchRHSTypeBoolean:
		// v3ServicesDeduplicationRulesExpressionBranchRHS.Boolean is populated
	case components.V3ServicesDeduplicationRulesExpressionBranchRHSTypeInt32:
		// v3ServicesDeduplicationRulesExpressionBranchRHS.Int32 is populated
	case components.V3ServicesDeduplicationRulesExpressionBranchRHSTypeFloat32:
		// v3ServicesDeduplicationRulesExpressionBranchRHS.Float32 is populated
}
```

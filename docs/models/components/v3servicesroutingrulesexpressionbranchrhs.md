# V3ServicesRoutingRulesExpressionBranchRHS


## Supported Types

### 

```go
v3ServicesRoutingRulesExpressionBranchRHS := components.CreateV3ServicesRoutingRulesExpressionBranchRHSStr(string{/* values here */})
```

### 

```go
v3ServicesRoutingRulesExpressionBranchRHS := components.CreateV3ServicesRoutingRulesExpressionBranchRHSBoolean(bool{/* values here */})
```

### 

```go
v3ServicesRoutingRulesExpressionBranchRHS := components.CreateV3ServicesRoutingRulesExpressionBranchRHSInt32(int{/* values here */})
```

### 

```go
v3ServicesRoutingRulesExpressionBranchRHS := components.CreateV3ServicesRoutingRulesExpressionBranchRHSFloat32(float32{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3ServicesRoutingRulesExpressionBranchRHS.Type {
	case components.V3ServicesRoutingRulesExpressionBranchRHSTypeStr:
		// v3ServicesRoutingRulesExpressionBranchRHS.Str is populated
	case components.V3ServicesRoutingRulesExpressionBranchRHSTypeBoolean:
		// v3ServicesRoutingRulesExpressionBranchRHS.Boolean is populated
	case components.V3ServicesRoutingRulesExpressionBranchRHSTypeInt32:
		// v3ServicesRoutingRulesExpressionBranchRHS.Int32 is populated
	case components.V3ServicesRoutingRulesExpressionBranchRHSTypeFloat32:
		// v3ServicesRoutingRulesExpressionBranchRHS.Float32 is populated
}
```

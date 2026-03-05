# OverlayUpdateDedupKeyOverlayResponseBody

The request has succeeded.


## Supported Types

### ResponseBody1

```go
overlayUpdateDedupKeyOverlayResponseBody := operations.CreateOverlayUpdateDedupKeyOverlayResponseBodyResponseBody1(operations.ResponseBody1{/* values here */})
```

### ResponseBody2

```go
overlayUpdateDedupKeyOverlayResponseBody := operations.CreateOverlayUpdateDedupKeyOverlayResponseBodyResponseBody2(operations.ResponseBody2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch overlayUpdateDedupKeyOverlayResponseBody.Type {
	case operations.OverlayUpdateDedupKeyOverlayResponseBodyTypeResponseBody1:
		// overlayUpdateDedupKeyOverlayResponseBody.ResponseBody1 is populated
	case operations.OverlayUpdateDedupKeyOverlayResponseBodyTypeResponseBody2:
		// overlayUpdateDedupKeyOverlayResponseBody.ResponseBody2 is populated
}
```

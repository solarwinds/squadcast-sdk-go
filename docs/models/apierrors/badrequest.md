# BadRequest

Represents a CircleCI error response for a 400 status code.


## Supported Types

### ResponseBodyError1

```go
badRequest := apierrors.CreateBadRequestResponseBodyError1(apierrors.ResponseBodyError1{/* values here */})
```

### ResponseBodyError2

```go
badRequest := apierrors.CreateBadRequestResponseBodyError2(apierrors.ResponseBodyError2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch badRequest.Type {
	case apierrors.BadRequestTypeResponseBodyError1:
		// badRequest.ResponseBodyError1 is populated
	case apierrors.BadRequestTypeResponseBodyError2:
		// badRequest.ResponseBodyError2 is populated
}
```

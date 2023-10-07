# grpcerrors

This package provides a convenient and performant way to use the standard gRPC error codes defined in https://grpc.github.io/grpc/core/md_doc_statuscodes.html.

Specifically, it provides:
* Static `var` instances of most status errors.
* A `struct` wrapper around all status errors.

## Benefits

`var` instances are helpful because they allow you to use `errors.Is` from the Go Standard Library. For example, say you write a package to interact with a database and you need to handle a case where a lookup is not found. You could `return grpcerrors.ErrNotFound` in your package. Higher level code orchestrating the database lookup can use `errors.Is(err, grpcerrors.ErrNotFound)` if it needs to understand what type of error occured. If the error is appropriate to return through an API, this error object can be returned directly through a gRPC API without any conversions.

Status `INVALID_ARGUMENT` is notably not provided as a static `var`. The reason is because this error should be accompanied with a message describing which argument is invalid. A static `var` would create a data race in a concurrent execution environment, so it's avoided.

The `struct` wrappers are helpful when a using other libraries in the grpc-go ecosystem. For example, [the gRPC status package implements a FromError function](https://github.com/grpc/grpc-go/blob/v1.58.2/status/status.go#L100C29-L100C41). This function declares an interface that requires `GRPCStatus()` be implemented. If it isn't, `FromError` converts a specific error into a vague unknown error. The [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) package is one such package where `FromError` is called. Without implementing `GRPCStatus()`, the JSON REST API returns an unknown error type - not very useful to the caller or for debugging purposes.

## Install
```
go get 'github.com/danielwchapman/grpcerrors'
```

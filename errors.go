package grpcerror

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gRPC errors defined as instances so they can efficiently
// be compared using errors.Is from the standard library.
var (
	ErrContextCancelled   = &CancelledError{}
	ErrUnknown            = &UnknownError{}
	ErrDeadlineExceeded   = &DeadlineExceededError{}
	ErrNotFound           = &NotFoundError{}
	ErrAlreadyExists      = &AlreadyExistsError{}
	ErrPermissionDenied   = &PermissionDeniedError{}
	ErrResourceExhausted  = &ResourceExhaustedError{}
	ErrFailedPrecondition = &FailedPreconditionError{}
	ErrAborted            = &AbortedError{}
	ErrOutOfRange         = &OutOfRangeError{}
	ErrUnimplemented      = &UnimplementedError{}
	ErrInternal           = &InternalError{}
	ErrUnavailable        = &UnavailableError{}
	ErrDataLoss           = &DataLossError{}
	ErrUnauthenticated    = &UnauthenticatedError{}
)

var (
	statusCancelled          = status.New(codes.Canceled, "context cancelled")
	statusUnknown            = status.New(codes.Unknown, "unknown error")
	statusInvalidArgument    = status.New(codes.InvalidArgument, "invalid argument")
	statusDeadlineExceeded   = status.New(codes.DeadlineExceeded, "deadline exceeded")
	statusNotFound           = status.New(codes.NotFound, "not found")
	statusAlreadyExists      = status.New(codes.AlreadyExists, "already exists")
	statusPermissionDenied   = status.New(codes.PermissionDenied, "permission denied")
	statusResourceExhausted  = status.New(codes.ResourceExhausted, "resource exhausted")
	statusFailedPrecondition = status.New(codes.FailedPrecondition, "failed precondition")
	statusAborted            = status.New(codes.Aborted, "aborted")
	statusOutOfRange         = status.New(codes.OutOfRange, "out of range")
	statusUnimplemented      = status.New(codes.Unimplemented, "unimplemented")
	statusInternal           = status.New(codes.Internal, "internal error")
	statusUnavailable        = status.New(codes.Unavailable, "unavailable")
	statusDataLoss           = status.New(codes.DataLoss, "data loss")
	statusUnauthenticated    = status.New(codes.Unauthenticated, "unauthenticated")
)

type CancelledError struct{}

func (e *CancelledError) Error() string {
	return "context cancelled"
}

func (e *CancelledError) GRPCStatus() *status.Status {
	return statusCancelled
}

func (e *CancelledError) Unwrap() error {
	return context.Canceled
}

type UnknownError struct{}

func (e *UnknownError) Error() string {
	return "unknown error"
}

func (e *UnknownError) GRPCStatus() *status.Status {
	return statusUnknown
}

type InvalidArgumentError struct {
	msg string
}

func MakeInvalidArgumentError(msg string) *InvalidArgumentError {
	return &InvalidArgumentError{msg: msg}
}

func (e *InvalidArgumentError) Error() string {
	return e.msg
}

func (e *InvalidArgumentError) GRPCStatus() *status.Status {
	if e.msg == "" {
		return statusInvalidArgument
	}
	return status.New(codes.InvalidArgument, e.msg)
}

func IsInvalidArgumentError(err error) (error, bool) {
	var invalidArgumentErr *InvalidArgumentError
	if errors.As(err, &invalidArgumentErr) {
		return invalidArgumentErr, true
	}
	return nil, false
}

type DeadlineExceededError struct{}

func (e *DeadlineExceededError) Error() string {
	return "deadline exceeded"
}

func (e *DeadlineExceededError) GRPCStatus() *status.Status {
	return statusDeadlineExceeded
}

func (e *DeadlineExceededError) Unwrap() error {
	return context.DeadlineExceeded
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "not found"
}

func (e *NotFoundError) GRPCStatus() *status.Status {
	return statusNotFound
}

type AlreadyExistsError struct{}

func (e *AlreadyExistsError) Error() string {
	return "already exists"
}

func (e *AlreadyExistsError) GRPCStatus() *status.Status {
	return statusAlreadyExists
}

type PermissionDeniedError struct {
	msg string
}

func MakePermissionDeniedError(msg string) *PermissionDeniedError {
	return &PermissionDeniedError{msg: msg}
}

func (e *PermissionDeniedError) Error() string {
	if e.msg == "" {
		return "permission denied"
	}
	return e.msg
}

func (e *PermissionDeniedError) GRPCStatus() *status.Status {
	if e.msg == "" {
		return statusPermissionDenied
	}
	return status.New(codes.PermissionDenied, e.msg)
}

type ResourceExhaustedError struct{}

func (e *ResourceExhaustedError) Error() string {
	return "resource exhausted"
}

func (e *ResourceExhaustedError) GRPCStatus() *status.Status {
	return statusResourceExhausted
}

type FailedPreconditionError struct {
	msg string
}

func MakeFailedPreconditionError(msg string) *FailedPreconditionError {
	return &FailedPreconditionError{msg: msg}
}

func (e *FailedPreconditionError) Error() string {
	if e.msg == "" {
		return "failed precondition"
	}
	return e.msg
}

func (e *FailedPreconditionError) GRPCStatus() *status.Status {
	if e.msg == "" {
		return statusFailedPrecondition
	}
	return status.New(codes.FailedPrecondition, e.msg)
}

type AbortedError struct{}

func (e *AbortedError) Error() string {
	return "aborted"
}

func (e *AbortedError) GRPCStatus() *status.Status {
	return statusAborted
}

type OutOfRangeError struct{}

func (e *OutOfRangeError) Error() string {
	return "out of range"
}

func (e *OutOfRangeError) GRPCStatus() *status.Status {
	return statusOutOfRange
}

type UnimplementedError struct{}

func (e *UnimplementedError) Error() string {
	return "unimplemented"
}

func (e *UnimplementedError) GRPCStatus() *status.Status {
	return statusUnimplemented
}

type InternalError struct {
	err error
}

func MakeInternalError(msg string) *InternalError {
	return &InternalError{err: errors.New(msg)}
}

func (e *InternalError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return "internal error"
}

func (e *InternalError) GRPCStatus() *status.Status {
	// always return "internal error" as message to avoid exposing internal details to clients
	return statusInternal
}

type UnavailableError struct{}

func (e *UnavailableError) Error() string {
	return "unavailable"
}

func (e *UnavailableError) GRPCStatus() *status.Status {
	return statusUnavailable
}

type DataLossError struct{}

func (e *DataLossError) Error() string {
	return "data loss"
}

func (e *DataLossError) GRPCStatus() *status.Status {
	return statusDataLoss
}

type UnauthenticatedError struct{}

func (e *UnauthenticatedError) Error() string {
	return "unauthenticated"
}

func (e *UnauthenticatedError) GRPCStatus() *status.Status {
	return statusUnauthenticated
}

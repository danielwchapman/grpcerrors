package grpcerrors

import (
	"context"
	"errors"
	"log/slog"
	"runtime"
)

// CleanError ensures sensitive data is not leaked into public API error messages.
// It also converts common error types, like ContextCancelled and DeadlineExceeded,
// to gRPC equivalents. In its current form, it does not translate all error types.
// When it is in doubt, it returns and logs an Internal Error.
func CleanError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	var invalidArgumentErr *InvalidArgumentError
	if errors.As(err, &invalidArgumentErr) {
		return invalidArgumentErr
	}
	if errors.Is(err, context.Canceled) {
		return ErrContextCancelled
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return ErrDeadlineExceeded
	}
	if errors.Is(err, ErrNotFound) {
		return ErrNotFound
	}
	if errors.Is(err, ErrAlreadyExists) {
		return ErrAlreadyExists
	}
	if errors.Is(err, ErrPermissionDenied) {
		return ErrPermissionDenied
	}
	if errors.Is(err, ErrResourceExhausted) {
		return ErrResourceExhausted
	}

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	// catch all
	slog.ErrorContext(ctx, f.Name()+" failed", "error", err)
	return ErrInternal
}

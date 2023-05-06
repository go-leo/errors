// Package errors provides simple error handling primitives.
package errors

import (
	"fmt"
	"io"
)

// New return a std error.
func New(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// NewWithStack return a error with stack.
func NewWithStack(format string, args ...interface{}) error {
	return &withStack{
		fmt.Errorf(format, args...),
		callers(),
	}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	if wc := new(withCode); As(err, &wc) {
		return &withCode{
			err:   wc.err,
			code:  wc.code,
			cause: err,
			stack: callers(),
		}
	}
	if e := new(withStack); As(err, &e) {
		return e
	}

	return &withStack{
		err,
		callers(),
	}
}

// can use this function instead WithMessage(WithStack(err), message).
// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func WrapStack(err error, message string) error {
	if err == nil {
		return nil
	}
	if wc := new(withCode); As(err, &wc) {
		return &withCode{
			err:   fmt.Errorf(message),
			code:  wc.code,
			cause: err,
			stack: callers(),
		}
	}

	messageErr := &withMessage{
		cause: err,
		msg:   message,
	}
	if ws := new(withStack); As(err, &ws) {
		return messageErr
	}

	return &withStack{
		messageErr,
		callers(),
	}
}

// can use this function instead WithMessagef(WithStack(err), format, args...).
// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func WrapStackf(err error, format string, args ...interface{}) error {
	return WrapStack(err, fmt.Sprintf(format, args...))
}

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error { return w.error }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withStack) Unwrap() error {
	if e, ok := w.error.(interface{ Unwrap() error }); ok {
		return e.Unwrap()
	}

	return w.error
}

// Format nolint: errcheck // WriteString could no check in pkg.
func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		if s.Flag('-') {
			fmt.Fprintf(s, "%-v", w.Cause())
			w.stack.Format(s, verb)
			return
		}

		fallthrough
	case 's':
		_, _ = io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

// option for withCode
type option func(*withCode)

// WithSkipDepth set skip depth.
func WithSkipDepth(skipDepth int) option {
	return func(w *withCode) {
		w.skipDepth = skipDepth
	}
}

// NewWithCode new error has default describe.
func NewWithCode(code int, format string, args ...interface{}) error {
	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		stack: callers(),
	}
}

// NewWithCodeX new error with code with options.
func NewWithCodeX(code int, message string, opts ...option) error {
	w := &withCode{
		err:   fmt.Errorf(message),
		code:  code,
		stack: callers(),
	}

	for _, opt := range opts {
		opt(w)
	}

	return w
}

// WrapC return an error annotating err with a stack trace and error code.
// Deprecated: Use WithCode instead.
func WrapC(err error, code int) error {
	if err == nil {
		return nil
	}

	return &withCode{
		err:   err,
		code:  code,
		cause: err,
		stack: callers(),
	}
}

// WithCode return an error annotating err with a stack trace and error code.
func WithCode(err error, code int) error {
	if err == nil {
		return nil
	}

	return &withCode{
		err:   err,
		code:  code,
		cause: err,
		stack: callers(),
	}
}

func WrapCode(err error, code int, message string) error {
	if err == nil {
		return nil
	}

	return &withCode{
		err:   fmt.Errorf(message),
		code:  code,
		cause: err,
		stack: callers(),
	}
}

func WrapCodef(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		cause: err,
		stack: callers(),
	}
}

type withCode struct {
	err   error
	code  int
	cause error
	*stack
	// jumpDepth is the number of stack frames to skip when reporting
	skipDepth int
}

// Error return the externally-safe error message.
func (w *withCode) Error() string { return fmt.Sprintf("%v", w) }

// Cause return the cause of the WithCode error.
func (w *withCode) Cause() error { return w.cause }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withCode) Unwrap() error { return w.cause }

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}

	return &withMessage{
		cause: err,
		msg:   message,
	}
}

// WithMessagef annotates err with the format specifier.
// If err is nil, WithMessagef returns nil.
func WithMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg }
func (w *withMessage) Cause() error  { return w.cause }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withMessage) Unwrap() error { return w.cause }

//nolint: errcheck // WriteString could no check in pkg
func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			io.WriteString(s, w.msg)

			return
		}

		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}

		if cause.Cause() == nil {
			break
		}

		err = cause.Cause()
	}

	return err
}

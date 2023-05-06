package code

//go:generate codegen -type=int
//go:generate codegen -type=int -doc -output ../docs/error_code_generated.md

// base: base errors.
const (
	// ErrUnknown - 500: Internal server error.
	ErrUnknown int = iota + 100001

	// ErrBind - 400: Error occurred while binding the request body to the struct.
	ErrBind

	// ErrValidation - 400: Validation failed.
	ErrValidation
)

// Account-server: Account errors.
const (
	// ErrAccountAuthTypeInvalid - 400: Account AuthType not support.
	ErrAccountAuthTypeInvalid int = iota + 110001

	// ErrUserNotFound - 400: User Not Found.
	ErrUserNotFound

	// ErrUserDisabled - 400: User disabled.
	ErrUserDisabled
)

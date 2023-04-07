package code

//go:generate codegen -type=int
//go:generate codegen -type=int -doc -output ../docs/error_code_generated.md

// Account-server: Account errors.
const (
	// ErrAccountAuthTypeInvalid - 400: Account AuthType not support.
	ErrAccountAuthTypeInvalid int = iota + 110001

	// ErrAccountGenerateTokenFailed - 500: Account generate token failed.
	ErrAccountGenerateTokenFailed

	// ErrAccountAlreadyLogin - 200: Account already login, logout to login other account.
	ErrAccountAlreadyLogin
)

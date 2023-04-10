// Code generated by "codegen -type=int  "; DO NOT EDIT.
package code

import "github.com/go-leo/errors"
import "fmt"

// init register error codes defines in this source code to `github.com/go-leo/errors`
func init() {
	register(ErrAccountAuthTypeInvalid, 400, "Account AuthType not support")
	register(ErrAccountGenerateTokenFailed, 500, "Account generate token failed")
	register(ErrAccountAlreadyLogin, 200, "Account already login, logout to login other account")
}

// Account AuthType not support
func IsErrAccountAuthTypeInvalid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.ParseCoder(err)
	return e.Code() == 110001 && e.HTTPStatus() == 400
}

// Account AuthType not support
func NewErrAccountAuthTypeInvalid(format string, args ...interface{}) error {
	return errors.NewWithCodeX(ErrAccountAuthTypeInvalid, fmt.Sprintf(format, args...), errors.WithSkipDepth(1))
}

// Account generate token failed
func IsErrAccountGenerateTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.ParseCoder(err)
	return e.Code() == 110002 && e.HTTPStatus() == 500
}

// Account generate token failed
func NewErrAccountGenerateTokenFailed(format string, args ...interface{}) error {
	return errors.NewWithCodeX(ErrAccountGenerateTokenFailed, fmt.Sprintf(format, args...), errors.WithSkipDepth(1))
}

// Account already login, logout to login other account
func IsErrAccountAlreadyLogin(err error) bool {
	if err == nil {
		return false
	}
	e := errors.ParseCoder(err)
	return e.Code() == 110003 && e.HTTPStatus() == 200
}

// Account already login, logout to login other account
func NewErrAccountAlreadyLogin(format string, args ...interface{}) error {
	return errors.NewWithCodeX(ErrAccountAlreadyLogin, fmt.Sprintf(format, args...), errors.WithSkipDepth(1))
}

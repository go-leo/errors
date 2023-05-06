package code

import (
	"net/http"

	"github.com/go-leo/errors"
	"golang.org/x/exp/slices"
)

// ErrCode implements `panda/pkg/errors`.Coder interface.
type ErrCode struct {
	// C refers to the code of the ErrCode.
	C int `json:"code,omitempty"`

	// HTTP status that should be used for the associated error code.
	HTTP int `json:"http,omitempty"`

	// External (user) facing error text.
	Ext string `json:"msg,omitempty"`

	// Ref specify the reference document.
	Ref string `json:"ref,omitempty"`
}

var _ errors.Coder = &ErrCode{}

// Code returns the integer code of ErrCode.
func (coder ErrCode) Code() int {
	return coder.C
}

// String implements stringer. String returns the external error message,
// if any.
func (coder ErrCode) String() string {
	return coder.Ext
}

// Reference returns the reference document.
func (coder ErrCode) Reference() string {
	return coder.Ref
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder ErrCode) HTTPStatus() int {
	if coder.HTTP == 0 {
		return http.StatusInternalServerError
	}

	return coder.HTTP
}

//nolint: unparam // .
func register(code int, httpStatus int, message string, refs ...string) {
	found := slices.Contains([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}

	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}

	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  reference,
	}

	errors.MustRegister(coder)
}

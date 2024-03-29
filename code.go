package errors

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"google.golang.org/grpc/status"
)

// global default error codes
var (
	UnknownCoder    Coder = defaultCoder{1, http.StatusInternalServerError, "An internal server error occurred", "http://github.com/go-leo/errors/README.md"}
	BindCoder       Coder = defaultCoder{2, http.StatusBadRequest, "Error occurred while binding the request params to the struct", ""}
	ValidationCoder Coder = defaultCoder{3, http.StatusBadRequest, "Request params validate failed", ""}
)

// Coder defines an interface for an error code detail information.
type Coder interface {
	// HTTP status that should be used for the associated error code.
	HTTPStatus() int

	// External (user) facing error text.
	String() string

	// Reference returns the detail documents for user.
	Reference() string

	// Code returns the code of the coder
	Code() int
}

type defaultCoder struct {
	// C refers to the integer code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Ext string

	// Ref specify the reference document.
	Ref string
}

// Code returns the integer code of the coder.
func (coder defaultCoder) Code() int {
	return coder.C
}

// String implements stringer. String returns the external error message,
// if any.
func (coder defaultCoder) String() string {
	return coder.Ext
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder defaultCoder) HTTPStatus() int {
	if coder.HTTP == 0 {
		return http.StatusInternalServerError
	}

	return coder.HTTP
}

// Reference returns the reference document.
func (coder defaultCoder) Reference() string {
	return coder.Ref
}

// codes contains a map of error codes to metadata.
var (
	codes   = map[int]Coder{}
	codeMux = &sync.Mutex{}
)

// Register register a user define error code.
// It will overrid the exist code.
func Register(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved by `github.com/panda/errors` as unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	codes[coder.Code()] = coder
}

// MustRegister register a user define error code.
// It will panic when the same Code already exist.
func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("code '0' is reserved by 'github.com/panda/errors' as ErrUnknown error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	codes[coder.Code()] = coder
}

// ParseCoder parse any error into *WithCode.
// nil error will return nil direct.
// None withStack error will be parsed as ErrUnknown.
func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}

	v := new(withCode)

	if errors.As(err, &v) {
		if coder, ok := codes[v.code]; ok {
			return coder
		}
	}

	ge, ok := status.FromError(err)
	if !ok {
		return UnknownCoder
	}

	for _, detail := range ge.Details() {
		switch d := detail.(type) {
		case *Status:
			return codes[int(d.Code)]
		}
	}

	return UnknownCoder
}

// GRPCErr convert error to grpc error.
// if err no register Coder, return unknown grpc error.
// func GRPCErr(err error) error {
// 	s := GRPCStatus(err)
// 	if s == nil {
// 		return nil
// 	}
// 	return s.Err()
// }

// GRPCStatus convert error to grpc *status.Status.
// if err no register Coder, return unknown grpc error.
func GRPCStatus(err error) *status.Status {
	if err == nil {
		return nil
	}

	var c Coder = UnknownCoder
	if v := new(withCode); errors.As(err, &v) {
		coder, ok := codes[v.code]
		if ok {
			c = coder
		}
	}

	s, _ := status.New(ToGRPCCode(c.HTTPStatus()), c.String()).
		WithDetails(&Status{
			Code: int32(c.Code()),
			Http: int32(c.HTTPStatus()),
			Ref:  c.Reference(),
		})

	return s
}

// GRPCCodeStatus convert code to grpc *status.Status.
// if err no register Coder, return unknown grpc error.
// If code is known, it is more efficient to use this method than GRPCStatus.
func GRPCCodeStatus(code int) *status.Status {
	c := GetCoder(code)
	s, _ := status.New(ToGRPCCode(c.HTTPStatus()), c.String()).
		WithDetails(&Status{
			Code: int32(c.Code()),
			Http: int32(c.HTTPStatus()),
			Ref:  c.Reference(),
		})

	return s
}

// GetCoder get Coder with code
// not found return ErrUnknown
// note: can not be change
func GetCoder(code int) Coder {
	if coder, ok := codes[code]; ok {
		return coder
	}

	return UnknownCoder
}

// IsCode reports whether any error in err's chain contains the given error code.
func IsCode(err error, code int) bool {
	if coder := ParseCoder(err); coder != nil {
		return coder.Code() == code
	}

	return false
}

func init() {
	codes[UnknownCoder.Code()] = UnknownCoder
}

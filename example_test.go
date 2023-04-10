package errors

import (
	"errors"
	"fmt"
)

func ExampleNewWithCode() {
	err := NewWithCode(ErrInvalidJSON, "id 1000")
	fmt.Println(err)
	fmt.Printf("%-v\n", err)
	fmt.Printf("%+v\n", err)

	// Output: Data is not valid JSON
	//#0 (1001) Data is not valid JSON, id 1000 [/Users/litao/code/errors/example_test.go:9 (github.com/go-leo/errors.ExampleNewWithCode)]
	//#0 (1001) Data is not valid JSON, id 1000 [/Users/litao/code/errors/example_test.go:9 (github.com/go-leo/errors.ExampleNewWithCode)]
}

func ExampleWrapC() {
	var err error
	err = NewWithCode(ConfigurationNotValid, "this is an error message")
	fmt.Println(err)
	err = WrapC(err, ErrInvalidJSON)
	fmt.Println(err)
	fmt.Println()

	err = errors.New("this is an error message")
	fmt.Println(err)
	err = WrapC(err, ErrInvalidJSON)
	fmt.Println(err)

	// Output:
	// ConfigurationNotValid error
	// Data is not valid JSON
	//
	// this is an error message
	// Data is not valid JSON
}

func ExampleWithCode_code() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed)
	}

	fmt.Println(err.(*withCode).code)
	// Output: 1003
}

func ExampledefaultCoder_HTTPStatus() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed)
	}

	fmt.Println(codes[err.(*withCode).code].HTTPStatus())
	// Output: 500
}

func ExampleCoder_HTTPStatus() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed)
	}

	coder := ParseCoder(err)
	fmt.Println(coder.HTTPStatus())
	// Output: 500
}

func ExampleString() {
	err := loadConfig()
	if nil != err {
		err = WrapC(err, ErrLoadConfigFailed)
	}

	fmt.Println(codes[err.(*withCode).code].String())
	// Output: Load configuration file failed
}

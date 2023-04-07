package example

import (
	"fmt"

	"github.com/leo/errors"
	"github.com/leo/errors/example/code"
)

func Example() {
	err := getUser("")
	fmt.Println(err)
	fmt.Printf("%-v\n", err)
	fmt.Printf("%+v\n", err)

	err = getUser("1")
	fmt.Println(err)
	fmt.Printf("%-v\n", err)
	fmt.Printf("%+v\n", err)

	err = getUser("2")
	fmt.Println(err)
	fmt.Printf("%-v\n", err)
	fmt.Printf("%+v\n", err)

	//Output:
	//Account AuthType not support
	// #0 (110001) Account AuthType not support, token  [/Users/litao/code/errors/example/examp_test.go:64 (github.com/leo/errors/example.getUser)]
	// #0 (110001) Account AuthType not support, token  [/Users/litao/code/errors/example/examp_test.go:64 (github.com/leo/errors/example.getUser)]
	// datebase, connection error!
	// datebase, connection error!
	// github.com/leo/errors/example.getUserByID
	// 	/Users/litao/code/errors/example/examp_test.go:82
	// github.com/leo/errors/example.getUser
	// 	/Users/litao/code/errors/example/examp_test.go:71
	// github.com/leo/errors/example.Example
	// 	/Users/litao/code/errors/example/examp_test.go:16
	// datebase, connection error!
	// github.com/leo/errors/example.getUserByID
	// 	/Users/litao/code/errors/example/examp_test.go:82
	// github.com/leo/errors/example.getUser
	// 	/Users/litao/code/errors/example/examp_test.go:71
	// github.com/leo/errors/example.Example
	// 	/Users/litao/code/errors/example/examp_test.go:16
	// testing.runExample
	// 	/Users/litao/.go/current/src/testing/run_example.go:63
	// testing.runExamples
	// 	/Users/litao/.go/current/src/testing/example.go:44
	// testing.(*M).Run
	// 	/Users/litao/.go/current/src/testing/testing.go:1721
	// main.main
	// 	_testmain.go:49
	// runtime.main
	// 	/Users/litao/.go/current/src/runtime/proc.go:250
	// runtime.goexit
	// 	/Users/litao/.go/current/src/runtime/asm_amd64.s:1571
	// Account AuthType not support
	// #0 (110001) Account AuthType not support, uid  [/Users/litao/code/errors/example/examp_test.go:78 (github.com/leo/errors/example.getUserByID)]
	// #0 (110001) Account AuthType not support, uid  [/Users/litao/code/errors/example/examp_test.go:78 (github.com/leo/errors/example.getUserByID)]
}

func getUser(token string) error {
	if token == "" {
		return code.NewErrAccountAuthTypeInvalid("token %s", token)
		// return errors.NewWithCode(code.ErrAccountAuthTypeInvalid, "token %s", token)
	}
	var uid string
	if token == "1" {
		uid = "1"
	}
	return getUserByID(uid)
}

// mock repository select user
func getUserByID(uid string) error {
	if uid == "" {
		// return errors.NewWithStack("uid can not be empty!")
		return code.NewErrAccountAuthTypeInvalid("uid %s", uid)
	}
	err := dbSelectUser(uid)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// mock db select user
func dbSelectUser(uid string) error {
	return errors.New("datebase, connection error!")
}

package example

import (
	"testing"

	"github.com/leo/errors"
	"github.com/leo/errors/example/code"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	err := code.NewErrAccountAlreadyLogin("uid %d", 23333)
	assert.True(t, code.IsErrAccountAlreadyLogin(err))

	withMessageErr := errors.WithMessage(err, "warp a message")
	assert.True(t, code.IsErrAccountAlreadyLogin(withMessageErr))

	stackErr := errors.WithStack(err)
	assert.True(t, code.IsErrAccountAlreadyLogin(stackErr))

	// warpC 如果err没有code，则新增一个，如果有则会覆盖
	otherErr := errors.WrapC(err, code.ErrAccountAuthTypeInvalid)
	assert.False(t, code.IsErrAccountAlreadyLogin(otherErr))
}

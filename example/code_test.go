package example

import (
	"testing"

	"github.com/go-leo/errors"
	"github.com/go-leo/errors/example/code"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	err := code.NewErrUserNotFound("uid %d", 23333)
	assert.True(t, code.IsErrUserNotFound(err))

	withMessageErr := errors.WithMessage(err, "warp a message")
	assert.True(t, code.IsErrUserNotFound(withMessageErr))

	stackErr := errors.WithStack(err)
	assert.True(t, code.IsErrUserNotFound(stackErr))

	// warpC 如果err没有code，则新增一个，如果有则会覆盖
	otherErr := errors.WrapC(err, code.ErrUserDisabled)
	assert.False(t, code.IsErrUserNotFound(otherErr))
}

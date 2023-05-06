package data

import (
	"github.com/go-leo/errors"

	"github.com/go-leo/errors/example/code"
)

type UserRepo struct{}

func (ur *UserRepo) GetUser(uid int) (name string, err error) {
	if uid <= 10 {
		return "", code.NewErrUserDisabled("uid: %d", uid)
	} else if 10 < uid && uid <= 100 {
		return "", code.NewErrUserNotFound("uid: %d", uid)
	} else {
		e := errors.New("database conn failed!")
		return "", errors.WithStack(e)
	}
}

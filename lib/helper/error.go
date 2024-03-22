package helper

import (
	"fmt"
	"net/http"

	"gocrudssample/lib/constant"
)

var commonErrorMap = map[error]int{
	constant.ErrNotFound: http.StatusNotFound,
	constant.ErrConflict: http.StatusConflict,
}

// CommonError is
func CommonError(err error) (int, error) {
	switch err {
	case constant.ErrNotFound:
		return commonErrorMap[constant.ErrNotFound], constant.ErrNotFound
	case constant.ErrConflict:
		return commonErrorMap[constant.ErrConflict], constant.ErrConflict
	}
	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}

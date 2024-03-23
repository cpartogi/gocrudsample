package utils

import "gocrudsample/lib/helper"

func errorType(err error) (int, error) {
	return helper.CommonError(err)
}

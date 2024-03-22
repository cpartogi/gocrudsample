package utils

import "gocrudssample/lib/helper"

func errorType(err error) (int, error) {
	return helper.CommonError(err)
}

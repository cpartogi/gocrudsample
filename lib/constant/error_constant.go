package constant

import "fmt"

var (
	ErrNotFound     = fmt.Errorf("Data not found")
	ErrConflict     = fmt.Errorf("conflict, data already exist")
	ErrBadRequest   = fmt.Errorf("Bad request")
	ErrInvalidUuid  = fmt.Errorf("Invalid Id format (uuid required)")
	ErrTitle        = fmt.Errorf("Title required")
	ErrTypeNotFound = fmt.Errorf("Tutorial Type not found")
)

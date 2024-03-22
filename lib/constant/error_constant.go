package constant

import "fmt"

var (
	// ErrNotFound is
	ErrNotFound = fmt.Errorf("data not found")
	// ErrConflict is
	ErrConflict = fmt.Errorf("conflict, data already exist")
)

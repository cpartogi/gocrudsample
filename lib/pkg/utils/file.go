package utils

import "os"

func IsFileorDirExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

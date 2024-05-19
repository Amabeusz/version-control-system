package common

import (
	"os"
	"path/filepath"
)

func GetRootPath() string {
	ex, err := os.Executable()
	Check(err)
	return filepath.Dir(ex)
}

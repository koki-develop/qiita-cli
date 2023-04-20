package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(filename string) (*os.File, error) {
	for i := 0; ; i++ {
		ex, err := Exists(filename)
		if err != nil {
			return nil, err
		}
		if !ex {
			break
		}

		base := filepath.Base(filename)
		ext := filepath.Ext(filename)
		base = base[:len(base)-len(ext)]
		filename = fmt.Sprintf("%s(%d)%s", base, i, ext)
	}

	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func Exists(p string) (bool, error) {
	if _, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

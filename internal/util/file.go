package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func CreateFile(filename string) (*os.File, error) {
	out := filename
	dir := filepath.Dir(filename)
	ext := filepath.Ext(filename)
	base := filepath.Base(filename)
	base = base[:len(base)-len(ext)]

	for i := 0; ; i++ {
		ex, err := Exists(out)
		if err != nil {
			return nil, err
		}
		if !ex {
			break
		}

		out = path.Join(dir, fmt.Sprintf("%s(%d)%s", base, i, ext))
	}

	f, err := os.Create(out)
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

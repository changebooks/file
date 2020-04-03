package file

import (
	"path/filepath"
	"strings"
)

func PathInfo(path string) (base string, name string, extension string) {
	base = BasePath(path)
	p := strings.LastIndex(base, ".")
	if p >= 0 {
		name = base[:p]
		extension = base[p+1:]
	} else {
		name = base
	}
	return
}

func Abs(path string) (string, error) {
	return filepath.Abs(path)
}

func BasePath(path string) string {
	return filepath.Base(path)
}

func Directory(path string) string {
	return filepath.Dir(path)
}

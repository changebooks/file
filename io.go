package file

import (
	"io"
	"io/ioutil"
	"os"
)

func Rename(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func Copy(newPath string, oldPath string, perm os.FileMode) (written int64, err error) {
	src, err := os.Open(oldPath)
	defer func() {
		if src != nil {
			_ = src.Close()
		}
	}()
	if err != nil {
		return 0, err
	}

	dst, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE, perm)
	defer func() {
		if dst != nil {
			_ = dst.Close()
		}
	}()
	if err != nil {
		return 0, err
	}

	return io.Copy(dst, src)
}

func Read(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func Write(fileName string, s string, append bool, perm os.FileMode) error {
	if append {
		return Append(fileName, []byte(s), perm)
	} else {
		return ioutil.WriteFile(fileName, []byte(s), perm)
	}
}

func Append(fileName string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}

	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}

	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

func Remove(fileName string) error {
	return os.Remove(fileName)
}

func Mkdir(path string, perm os.FileMode) error {
	if IsDirExist(path) {
		return nil
	}

	if err := os.MkdirAll(path, perm); err != nil {
		return err
	}

	return nil
}

func IsDirExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

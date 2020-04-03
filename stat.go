package file

import "os"

func Size(path string) int64 {
	exist, info := IsExist(path)
	if exist && info != nil {
		return info.Size()
	} else {
		return -1
	}
}

func ModTime(path string) int64 {
	exist, info := IsExist(path)
	if exist && info != nil {
		return info.ModTime().Unix()
	} else {
		return -1
	}
}

func IsDir(path string) bool {
	exist, info := IsExist(path)
	if exist && info != nil {
		return info.IsDir()
	} else {
		return false
	}
}

func IsExist(path string) (bool, os.FileInfo) {
	info, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false, nil
	} else {
		return true, info
	}
}

func Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

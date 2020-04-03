package file

import "os"

const Perm = 0666

// 是否可读，尝试只读方式打开文件，判断后关闭文件
func IsReadable(path string) bool {
	f, err := os.OpenFile(path, os.O_RDONLY, Perm)
	defer func() {
		if f != nil {
			_ = f.Close()
		}
	}()

	if err != nil {
		if os.IsPermission(err) {
			return false
		}
	}

	return true
}

// 是否可写，尝试读写方式打开文件，判断后关闭文件
func IsWritable(path string) bool {
	f, err := os.OpenFile(path, os.O_WRONLY, Perm)
	defer func() {
		if f != nil {
			_ = f.Close()
		}
	}()

	if err != nil {
		if os.IsPermission(err) {
			return false
		}
	}

	return true
}

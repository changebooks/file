# file
文件
==

<pre>
path := ""

fileInfo, err := file.Stat(path)
fmt.Println(fileInfo, err)

isExist, fileInfo := file.IsExist(path)
fmt.Println(isExist, fileInfo)

isDir := file.IsDir(path)
fmt.Println(isDir)

modTime := file.ModTime(path)
fmt.Println(modTime)

size := file.Size(path)
fmt.Println(size)

isReadable := file.IsReadable(path)
fmt.Println(isReadable)

isWritable := file.IsWritable(path)
fmt.Println(isWritable)
</pre>
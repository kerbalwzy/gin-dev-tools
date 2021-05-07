package kerbalwzygo

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// 判断文件或者文件夹路径是否存在
func PathOk(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}

// 按行读取文件内容, 并判断是否为UTF8编码
func ValidFileUTF8(filepath string, checkLines int) (bool, error) {
	if checkLines <= 0 {
		checkLines = 10
	}
	isUTF8 := true
	fp, err := os.Open(filepath)
	if nil != err {
		return false, err
	}
	defer fp.Close()
	buf := bufio.NewReader(fp)
	for i := 0; i < checkLines; {
		line, err := buf.ReadBytes('\n')
		if len(line) > 0 {
			isUTF8 = utf8.Valid(line)
			i++
		}
		if io.EOF == err {
			break
		}
		if nil != err {
			return false, err
		}
	}
	return isUTF8, nil
}

// 获取目录下是所有文件的绝对路径(不含文件夹, 并且可以通过suffix过滤, 当suffix为空字符串或者"*"时表示配匹所有文件尾缀)
func ListDirFiles(dirPath, suffix string) ([]string, error) {
	res := make([]string, 0)
	_, err := PathOk(dirPath)
	if nil != err {
		return res, err
	}
	filterFunc := func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			res = append(res, path)
		}
		return err
	}
	if suffix != "*" && suffix != "" {
		filterFunc = func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() && strings.HasSuffix(path, suffix) {
				res = append(res, path)
			}
			return err
		}
	}
	err = filepath.WalkDir(dirPath, filterFunc)
	return res, err
}

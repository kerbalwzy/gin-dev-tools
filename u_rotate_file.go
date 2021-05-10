package kerbalwzygo

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Custom rotate file writer.
type RotateFileWriter struct {
	fileName string     // The file name, but the finally file name would be `<obj.fileName>.<timestamp>`.
	dirPath  string     // The dir path for saving the files.
	count    int        // The count of existed files.
	size     int64      // The size of the file which is used now.
	maxCount int        // The max value of the <obj.count>, meaning how much file can be saved totally.
	maxSize  int64      // Meaning how much bytes data one file can saving.
	fp       *os.File   // The current file object that for writing data
	wt       sync.Mutex // The lock for operate the 'count'
	once     sync.Once  // Keep initial just one times
}

// Check the target dir if existed which for saving files.
// When the result is false, will create the dir, if error happened there, panic.
func (obj *RotateFileWriter) checkDir() {
	info, err := os.Stat(obj.dirPath)
	if nil != err || !info.IsDir() {
		err := os.Mkdir(obj.dirPath, os.ModePerm)
		if nil != err {
			panic(err)
		}
	}
}

// Search the files under the target dir, which matching the target file format.
func (obj *RotateFileWriter) searchExistedFiles() map[int64]string {
	fileNameRe := regexp.MustCompile(fmt.Sprintf(`^%s.\d+$`, obj.fileName))
	exitedFiles := make(map[int64]string)
	files, _ := ioutil.ReadDir(obj.dirPath)
	for _, file := range files {
		if !file.IsDir() && fileNameRe.MatchString(file.Name()) {
			splitSlice := strings.Split(file.Name(), ".")
			timestampStr := splitSlice[len(splitSlice)-1]
			if timestamp, err := strconv.ParseInt(timestampStr, 10, 64); nil == err {
				exitedFiles[timestamp] = file.Name()
			}

		}
	}
	return exitedFiles
}

// Initial method, check whether the log dir existed and have history log files
func (obj *RotateFileWriter) Init() {
	obj.once.Do(func() {
		obj.checkDir()

		exitedFiles := obj.searchExistedFiles()
		exitedFilesCount := len(exitedFiles)

		// check weather have tag files here
		if exitedFilesCount > 0 {
			// if haven, find the newest one and use it
			var newEstLogFile int64
			for K := range exitedFiles {
				if K > newEstLogFile {
					newEstLogFile = K
				}
			}
			fileName := exitedFiles[newEstLogFile]
			fp, err := os.OpenFile(obj.dirPath+"/"+fileName, os.O_APPEND|os.O_WRONLY, 0666)
			if nil != err {
				panic(err)
			}
			fileInfo, err := fp.Stat()
			if nil != err {
				panic(err)
			}
			obj.fp = fp
			obj.count = exitedFilesCount
			obj.size = fileInfo.Size()

		} else {
			// if not haven, create one new log file
			obj.createNewFile()
		}
	})

}

func (obj *RotateFileWriter) createNewFile() {
	obj.wt.Lock()
	defer obj.wt.Unlock()
	obj.count++
	if obj.count > obj.maxCount {
		// if the count is large than maxCount, remove the oldest log file
		fileDateNums := make([]int64, 0)
		existedFiles := obj.searchExistedFiles()
		for k := range existedFiles {
			fileDateNums = append(fileDateNums, k)
		}
		oldEstFile := fileDateNums[0]
		for _, item := range fileDateNums {
			if item < oldEstFile {
				oldEstFile = item
			}
		}
		_ = os.Remove(obj.dirPath + "/" + existedFiles[oldEstFile])
		obj.count--
	}

	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%s.%d", obj.fileName, timeStamp)
	fp, _ := os.OpenFile(obj.dirPath+"/"+fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	obj.fp = fp
	obj.size = 0
}

func (obj *RotateFileWriter) Write(p []byte) (n int, err error) {
	addSize := int64(len(p))
	newSize := obj.size + addSize
	if newSize > obj.maxSize {
		obj.createNewFile()
	}
	obj.size += addSize

	return obj.fp.Write(p)
}

// 循环文件写入器: fileName基本文件名, dirPath文件夹路径, maxCount最大文件数量, maxSize最大文件体积
func NewRotateFileWriter(fileName, dirPath string, maxCount int, maxSize int64) *RotateFileWriter {
	writer := &RotateFileWriter{
		fileName: fileName,
		dirPath:  dirPath,
		maxCount: maxCount,
		maxSize:  maxSize}
	writer.Init()
	return writer
}

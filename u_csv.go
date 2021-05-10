package kerbalwzygo

import (
	"encoding/csv"
	"io"
	"os"
)

// 读取csv文件(PS: 去除了空行)
func ReadCSV(filepath string) ([][]string, error) {
	res := make([][]string, 0)
	// 获取文件句柄
	fp, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	r := csv.NewReader(fp)
	for {
		row, err := r.Read()
		if len(row) > 0 {
			res = append(res, row)
		}
		if err == io.EOF {
			break
		}
	}
	return res, nil
}

package kerbalwzygo

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"math"
	"regexp"
	"strconv"
)

var ExcelIllegalCharactersRe = regexp.MustCompile(`[\000-\010]|[\013-\014]|[\016-\037]`)

//Excel单个sheet最多只能有1048576行,超出的行数据将保存到复制了名称的sheet
const ExcelMaxRowCount = 1048576

//const ExcelMaxRowCount = 2

type ExcelSheet struct {
	Name      string
	Content   [][]interface{}
	safeLimit int
}

func (obj *ExcelSheet) Len() int {
	return len(obj.Content)
}

func (obj *ExcelSheet) SetSafeLimit(n int) {
	if n == 0 || n > ExcelMaxRowCount {
		n = ExcelMaxRowCount
	}
	obj.safeLimit = n
}

func (obj *ExcelSheet) Safe() []ExcelSheet {
	if obj.safeLimit == 0 || obj.safeLimit > ExcelMaxRowCount {
		obj.safeLimit = ExcelMaxRowCount
	}
	res := make([]ExcelSheet, 1)
	if obj.Len() <= obj.safeLimit {
		res[0] = *obj
		return res
	}
	// 超出了安全上限
	for i := 0; i < int(math.Ceil(float64(obj.Len())/float64(obj.safeLimit))); i++ {
		if i == 0 {
			res[0] = ExcelSheet{
				Name:    obj.Name,
				Content: obj.Content[0:obj.safeLimit],
			}
		} else {
			res = append(res, ExcelSheet{
				Name:    obj.Name + "-" + strconv.Itoa(i),
				Content: obj.Content[obj.safeLimit*i : obj.safeLimit*(i+1)],
			})
		}
	}
	return res
}

func MakeExcelFp(data ...ExcelSheet) (*excelize.File, error) {
	fp := excelize.NewFile()
	for index, item := range data {
		if index == 0 {
			fp.SetSheetName("Sheet1", item.Name)
		} else {
			fp.NewSheet(item.Name)
		}
		streamWriter, err := fp.NewStreamWriter(item.Name)
		if nil != err {
			fmt.Println(err)
			return nil, err
		}
		for rowIndex, row := range item.Content {
			cell, _ := excelize.CoordinatesToCellName(1, rowIndex+1)
			err = streamWriter.SetRow(cell, row)
			if nil != err {
				return nil, err
			}
		}
		err = streamWriter.Flush()
		if nil != err {
			return nil, err
		}
	}
	return fp, nil
}

func SafeMakeExcelFp(data ...ExcelSheet) (*excelize.File, error) {
	tempData := make([]ExcelSheet, 0, len(data))
	for _, item := range data {
		tempData = append(tempData, item.Safe()...)
	}
	return MakeExcelFp(tempData...)
}

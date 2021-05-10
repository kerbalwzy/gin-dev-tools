package kerbalwzygo

import (
	"testing"
	"time"
)

func TestSheetData(t *testing.T) {
	sheet := ExcelSheet{
		Name: "测试",
	}
	sheet.SetSafeLimit(2)
	sheet.Content = append(sheet.Content, []interface{}{1, 3})
	sheet.Content = append(sheet.Content, []interface{}{1.2, 3})
	sheet.Content = append(sheet.Content, []interface{}{"hao", 3})
	sheet.Content = append(sheet.Content, []interface{}{true, 3})
	sheet.Content = append(sheet.Content, []interface{}{time.Now(), 3})
	t.Log(sheet.Len())
	t.Log(len(sheet.Safe()))
}

func TestMakeExcelFp(t *testing.T) {
	sheet := ExcelSheet{
		Name: "测试",
	}
	sheet.SetSafeLimit(2)
	sheet.Content = append(sheet.Content, []interface{}{1, 3})
	sheet.Content = append(sheet.Content, []interface{}{1.2, 3})
	sheet.Content = append(sheet.Content, []interface{}{"hao", 3})
	sheet.Content = append(sheet.Content, []interface{}{true, 3})
	sheet.Content = append(sheet.Content, []interface{}{time.Now().Format("2006-01-02 15:04:05"), 3})
	t.Log(sheet.Len())
	t.Log(len(sheet.Safe()))
	fp, err := MakeExcelFp(sheet.Safe()...)
	if nil != err {
		t.Fatal(err)
	} else {
		t.Log(&fp)
		fp.SaveAs("u_excel_test.xlsx")
	}
}

func TestSafeMakeExcelFp(t *testing.T) {
	sheet := ExcelSheet{
		Name: "测试",
	}
	sheet.SetSafeLimit(2)
	sheet.Content = append(sheet.Content, []interface{}{1, 3})
	sheet.Content = append(sheet.Content, []interface{}{1.2, 3})
	sheet.Content = append(sheet.Content, []interface{}{"hao", 3})
	sheet.Content = append(sheet.Content, []interface{}{true, 3})
	sheet.Content = append(sheet.Content, []interface{}{time.Now().Format("2006-01-02 15:04:05"), 3})
	t.Log(sheet.Len())
	fp, err := SafeMakeExcelFp(sheet)
	if nil != err {
		t.Fatal(err)
	} else {
		t.Log(&fp)
		fp.SaveAs("u_excel_test.xlsx")
	}
}

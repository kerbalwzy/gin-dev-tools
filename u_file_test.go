package kerbalwzygo

import "testing"

func TestValidFileUTF8(t *testing.T) {
	filepath := "C:\\Users\\admin\\Desktop\\202001-202103放款数据\\202101放款-mars\\test.txt"
	yes, err := ValidFileUTF8(filepath, 100)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(yes)
	}
}

func TestListDirFiles(t *testing.T) {
	dirPath := "C:\\Users\\admin\\Desktop\\202001-202103放款数据"
	suffix := "csv"
	res, err := ListDirFiles(dirPath, suffix)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(len(res))
	}
}

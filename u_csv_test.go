package kerbalwzygo

import "testing"

func TestReadCSV(t *testing.T) {
	filepath := "C:\\Users\\admin\\Desktop\\202001-202103放款数据\\202101放款-mars\\dorius-us-1.csv"
	res, err := ReadCSV(filepath)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(len(res))
	}
}

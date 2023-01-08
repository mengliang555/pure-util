package csv_json_util

import (
	"testing"
)

func TestCsvToJson(t *testing.T) {
	v, er := CsvFileToJson("test_file.csv")
	if er != nil {
		t.Errorf(er.Error())
	}
	t.Log(string(v))
}

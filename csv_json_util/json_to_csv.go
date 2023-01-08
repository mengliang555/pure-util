package csv_json_util

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonToCsv(filepath string) ([]byte, error) {
	var erAns error
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	deco := json.NewDecoder(file)
	ans := make([]map[string]any, 0)
	er := deco.Decode(&ans)
	if er != nil {
		erAns = fmt.Errorf("faild to decode%s", filepath)
	}
	//TODO title and info
	return nil, erAns
}

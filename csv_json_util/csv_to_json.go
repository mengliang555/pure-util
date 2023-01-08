package csv_json_util

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var _once sync.Once

func doOnce(behave func()) {
	_once.Do(behave)
}

func CsvFileToJson(filepath string) ([]byte, error) {
	var erAns error
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	csv := csv.NewReader(file)
	successCount, totalCount := 0, 0

	res := new(bytes.Buffer)
	keyList := make([]string, 0)

	res.WriteByte('[')
	for {
		info, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("failed to read the file:%s when line:%d\n", filepath, successCount)
			erAns = fmt.Errorf("failed to read the file%s when line:%d", filepath, successCount)
		}
		if totalCount == 0 {
			keyList = append(keyList, info...)
		} else {
			if val, err := generateSingleItem(keyList, info, successCount < 2); err != nil {
				log.Printf("failed to generate the json with error:%s\n", err.Error())
				doOnce(func() {
					erAns = fmt.Errorf("%w and failed to generate the json with error:%s", erAns, err.Error())
				})
			} else {
				if _, err = res.Write(val); err != nil {
					log.Printf("failed to write the json with error:%s\n", err.Error())
				} else {
					successCount++
				}
			}
		}
		totalCount++
	}
	res.WriteByte(']')
	return res.Bytes(), erAns
}

func generateSingleItem(keyList []string, info []string, isFirst bool) ([]byte, error) {
	if len(keyList) != len(info) {
		return nil, errors.New("the length of csv is not match with title,pls check")
	}
	res := new(bytes.Buffer)
	if !isFirst {
		res.WriteByte(',')
	}
	res.WriteByte('{')
	for i, v := range keyList {
		res.WriteByte('"')
		res.WriteString(v)
		res.WriteString("\":\"")
		res.WriteString(info[i])
		res.WriteByte('"')
		if i != len(keyList)-1 {
			res.WriteByte(',')
		}
	}
	res.WriteByte('}')
	return res.Bytes(), nil
}

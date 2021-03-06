package scanfile

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/* 单个文件扫描后返回对象 */
type FileRes struct {
	File    string
	Strings StringArray
	Num     int
}

func InitFileRes(file string) *FileRes {
	p := new(FileRes)
	fileInfo, _ := os.Stat(file)
	p.File = fileInfo.Name()
	return p
}

func (res *FileRes) Add(str string) {
	res.Num++
	res.Strings.Set(strings.TrimRight(str, "\r"))
}

/* ScanFile 扫描后返回对象 */
type ScanResult struct {
	ResList []*FileRes
	Counter *Counter
}

func (result *ScanResult) AddCounter(counter *Counter) {
	result.Counter = counter
}

func (result *ScanResult) AddFileRes(res *FileRes) {
	if len(res.Strings) > 0 {
		result.ResList = append(result.ResList, res)
	}
}

func (result *ScanResult) ToJson() string {

	r, err := json.Marshal(result)

	if err != nil {
		fmt.Println(err)
		return "{}"
	}
	return string(r)
}

type StringArray []string

func (a *StringArray) Set(s string) {
	*a = append(*a, s)
}

func (a *StringArray) String() string {
	return fmt.Sprint(*a)
}

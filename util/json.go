package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func IndentJson(jsonDir string, jsonIndentedDir string) {
	//jsonDir := "/Users/fuyiwei/Downloads/shuju/"
	//jsonIndentedDir := "/Users/fuyiwei/Downloads/Indentedshuju/"
	jsons, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(jsons))
	for _, info := range jsons {
		open, err := os.Open(jsonDir + info.Name())
		if err != nil {
			fmt.Println(err)
		}
		all, err := ioutil.ReadAll(open)
		if err != nil {
			fmt.Println(err)
		}
		var out bytes.Buffer
		err = json.Indent(&out, all, "", "\t")
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(jsonIndentedDir+info.Name(), out.Bytes(), 0666)
		if err != nil {
			fmt.Println(err)
		}
	}
}

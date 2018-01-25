package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		Tree(os.Args[1], 1, map[int]bool{1: true})
	}
}

func Tree(dirname string, curHier int, hierMap map[int]bool) error {
	dirAbs, err := filepath.Abs(dirname)
	if err != nil {
		return err
	}

	fileInfos, err := ioutil.ReadDir(dirAbs)
	if err != nil {
		return err
	}
	fileNum := len(fileInfos)
	for i, fileInfo := range fileInfos {
		for j := 1; j < curHier; j++ {
			if hierMap[j] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(strings.Repeat(" ", 3))
		}

		// map是引用类型，所以新建一个map
		tmpMap := map[int]bool{}
		for k, v := range hierMap {
			tmpMap[k] = v
		}
		if i+1 == fileNum {
			fmt.Print("`")
			delete(tmpMap, curHier)
		} else {
			fmt.Print("|")
			tmpMap[curHier] = true
		}
		fmt.Print("-- ")
		fmt.Println(fileInfo.Name())
		if fileInfo.IsDir() {
			Tree(filepath.Join(dirAbs, fileInfo.Name()), curHier+1, tmpMap)
		}
	}
	return nil
}

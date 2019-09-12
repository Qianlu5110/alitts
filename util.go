/*
#
# @Time : 2019/9/10 16:34
# @Author : Qian Lu
*/

package main

import (
	"fmt"
	"os"
)

func InitWorkDir() bool {
	exist, err := PathExists(dir)
	if err != nil {
		fmt.Println(fmt.Sprintf("find work dir failed... [%v]\n", err))
		return false
	}

	if !exist {
		fmt.Printf("init work dir...\n")
		// 创建文件夹
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("init work dir failed... [%v]\n", err)
			return false
		} else {
			fmt.Printf("init work dir success\n")
		}
	}

	return true
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

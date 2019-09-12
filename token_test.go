/*
#
# @Time : 2019/9/10 16:20
# @Author : Qian Lu
*/

package main

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := GetAliToken()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Testing token:%s", token))
}

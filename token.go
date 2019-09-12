/*
#
# @Time : 2019/4/22 13:54
# @Author : Qian Lu
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func GetAliToken() (string, error) {
	client, err := sdk.NewClientWithAccessKey("cn-shanghai", AccessKeyId, AccessKeySecret)
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Domain = "nls-meta.cn-shanghai.aliyuncs.com"
	request.ApiName = "CreateToken"
	request.Version = "2019-02-28"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}

	if response.GetHttpStatus() != 200 {
		return "", errors.New(fmt.Sprintf("get ali token failed code : %d", response.GetHttpStatus()))
	}

	res := new(AccessTokenRes)
	err = json.Unmarshal(response.GetHttpContentBytes(), res)
	if err != nil {
		fmt.Println(fmt.Sprintf("get ali token res unmarshal err:%v, content : %s", err, response.GetHttpContentBytes()))
		return "", err
	}

	fmt.Println("get ali token success.")
	return res.Token.Id, nil
}

type AccessTokenRes struct {
	NlsRequestId string `json:"NlsRequestId"`
	RequestId    string `json:"RequestId"`
	ErrMsg       string `json:"ErrMsg"`
	Token        Token  `json:"Token"`
}

type Token struct {
	Id         string `json:"Id"`
	UserId     string `json:"UserId"`
	ExpireTime int64  `json:"ExpireTime"`
}

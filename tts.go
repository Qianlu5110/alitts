/*
#
# @Time : 2019/4/22 13:52
# @Author : Qian Lu
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func processPOSTRequest(
	token string,
	text string,
	audioSaveFile string,
	format string,
	sampleRate int,
	voice string) {

	/**
	 * 设置HTTPS POST请求
	 * 1.使用HTTPS协议
	 * 2.语音合成服务域名：nls-gateway.cn-shanghai.aliyuncs.com
	 * 3.语音合成接口请求路径：/stream/v1/tts
	 * 4.设置必须请求参数：appkey、token、text、format、sample_rate
	 * 5.设置可选请求参数：voice、volume、speech_rate、pitch_rate
	 */

	var url = "https://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/tts"
	bodyContent := make(map[string]interface{})
	bodyContent["appkey"] = AppKey
	bodyContent["text"] = text
	bodyContent["token"] = token
	bodyContent["format"] = format
	bodyContent["sample_rate"] = sampleRate
	// voice 发音人，可选，默认是xiaoyun
	bodyContent["voice"] = voice
	// volume 音量，范围是0~100，可选，默认50
	bodyContent["volume"] = 50
	// speech_rate 语速，范围是-500~500，可选，默认是0
	bodyContent["speech_rate"] = 0
	// pitch_rate 语调，范围是-500~500，可选，默认是0
	bodyContent["pitch_rate"] = 0

	bodyJson, err := json.Marshal(bodyContent)
	if err != nil {
		panic(nil)
	}
	fmt.Println("ready to execute tts request to ali.")
	/**
	 * 发送HTTPS POST请求，处理服务端的响应
	 */
	response, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte(bodyJson)))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	contentType := response.Header.Get("Content-Type")
	body, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		panic(fmt.Sprintf("get tts response code not right.Code : %d, Res: %s", response.StatusCode, string(body)))
	}

	fmt.Println("execute tts request succeed.")

	if "audio/mpeg" == contentType {
		file, _ := os.Create(audioSaveFile)
		defer file.Close()
		file.Write([]byte(body))
		fmt.Println("tts voice file save succeed.")
		fmt.Println(fmt.Sprintf("file path:%s", audioSaveFile))
	} else {
		fmt.Println(fmt.Sprintf("tts server return error. ContentType:%s", contentType))
	}
}

/*
#
# @Time : 2019/9/10 16:37
# @Author : Qian Lu
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

const AppVersion = "1.0.1 release"

var dir string
var voice string
var format string
var sample int
var txt string

func main() {
	// 查询版本
	version := flag.Bool("v", false, "show program version")

	// 获取执行参数
	flag.StringVar(&dir, "d", "",
		"Windows default dir C:/work/tts/ \n"+
			"Linux default dir /home/devops/tts/")
	flag.IntVar(&sample, "s", 8000,
		"tts sample, values:[8000,16000,24000] not all voice support 24000\n"+
			"default 8000")
	flag.StringVar(&voice, "voice", "Siyue",
		"ali tts voice, you can choose Xiaoyun,Xiaogang,Xiaomeng,Xiaowei,Ruoxi, \n"+
			"Siqi,Sijia,Sicheng,Aiqi,Aijia,Aicheng,Aida,Ninger,Ruilin,Amei,Xiaoxue, \n"+
			"Siyue,Aiya,Aixia,Aimei,Aiyu,Aiyue,Aijing,Xiaomei,Yina,Sijing,Sitong,Xiaobei, \n"+
			"Aitong,Aiwei,Aibao,Halen,Harry,Eric,Emily,Luna,Luca,Wendy,William,Olivia,Shanshan \n"+
			"default Siyue")
	flag.StringVar(&format, "format", "wav",
		"tts file type [wav or mp3] \n"+
			"default wav")
	flag.StringVar(&txt, "txt", "您好",
		"tts content waiting to compose \n"+
			"default to compose \"您好\"")
	flag.Parse()

	flag.Parse()
	if *version {
		fmt.Println(fmt.Sprintf("ali tts version:%s", AppVersion))
		os.Exit(0)
	}

	fmt.Printf("load config...\n")

	if !(format == "wav" || format == "mp3") {
		panic("format err")
	}

	if dir == "" || &dir == nil {
		switch runtime.GOOS {
		case "darwin":
			dir = "/home/devops/tts/"
		case "linux":
			dir = "/home/devops/tts/"
		case "windows":
			dir = "C:/work/tts/"
		}
	}
	fmt.Println(dir)

	// 准备工作目录
	workDirState := InitWorkDir()
	if !workDirState {
		panic("work dir init err")
	}

	// 获取阿里接口token
	token, err := GetAliToken()
	if err != nil {
		panic("ali token get err")
	}

	// 进行阿里TTs
	filePath := dir + txt + "_" + strconv.Itoa(sample) + "." + format
	processPOSTRequest(token, txt, filePath, format, sample, voice)
}

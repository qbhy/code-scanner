package main

import (
	"fmt"
	"github.com/bieber/barcode"
	"os"
	"flag"
	"image"
	"image/png"
	"image/jpeg"
	"image/gif"
	"strings"
	"encoding/json"
	"github.com/96qbhy/go-utils"
)

type Encode struct {
	Msg  string `json:"msg"`
	Type string `json:"type"`
	Text string `json:"text"`
}

func main() {
	// 获取命令行参数
	path := flag.String("path", "path", "file path")
	flag.Parse()

	if *path == "path" {
		outPut(Encode{
			Msg: "not found file",
		})
		return
	}

	exists, _ := utils.PathExists(*path)
	if exists == false {
		outPut(Encode{
			Msg: "not found file",
		})
		return
	}

	img := barcode.NewImage(DecodeImage(*path))

	scanner := barcode.NewScanner().SetEnabledAll(true)

	symbols, _ := scanner.ScanImage(img)
	if len(symbols) >= 1 {
		outPut(Encode{
			Text: symbols[0].Data,
			Type: symbols[0].Type.Name(),
		})
	} else {
		outPut(Encode{
			Msg: "not found code",
		})
	}
}

// 解析图片
func DecodeImage(imagePath string) image.Image {
	fin, _ := os.Open(imagePath)
	defer fin.Close()

	var imageType string
	indexOf := strings.LastIndex(imagePath, ".") + 1
	imageType = utils.SubStr(imagePath, indexOf, 5)
	if imageType == "png" {
		src, _ := png.Decode(fin)
		return src
	} else if imageType == "gif" {
		src, _ := gif.Decode(fin)
		return src
	} else {
		src, _ := jpeg.Decode(fin)
		return src
	}
}

// 统一输出格式
func outPut(encode Encode) {
	if encode.Msg == "" {
		encode.Msg = "success"
	}

	j, errs := json.Marshal(&encode) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	fmt.Print(string(j)) //byte[]转换成string 输出
}

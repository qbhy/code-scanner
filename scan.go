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
	"github.com/ricardolonga/jsongo"
	"encoding/json"
)

type Encode struct {
	msg  string `json:"msg"`
	text string `json:"text"`
}

func main() {
	// 获取命令行参数
	path := flag.String("path", "path", "file path")
	flag.Parse()

	if *path == "path" {
		fmt.Println("请输入文件地址")
		return
	}
	fin, _ := os.Open(*path)
	defer fin.Close()

	img := barcode.NewImage(decodeImage(*path))

	scanner := barcode.NewScanner().SetEnabledAll(true)

	symbols, _ := scanner.ScanImage(img)
	for _, s := range symbols {
		data := jsongo.Object()
		data.Put("text", s.Data).
			Put("type", s.Type.Name()).
			Put("quality", s.Quality)
		outPut(Encode{
			text: s.Data,
		})
	}
}

// 解析图片
func decodeImage(imagePath string) image.Image {
	fin, _ := os.Open(imagePath)
	var imageType string = "png"
	indexOf := strings.LastIndex(imagePath, ".") + 1
	imageType = Substr(imagePath, indexOf, 5)
	defer fin.Close()
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

// 截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// 统一输出格式
func outPut(encode Encode) {
	if encode.msg == "" {
		encode.msg = "success"
	}

	j, errs := json.Marshal(encode) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	fmt.Println(string(j)) //byte[]转换成string 输出
	fmt.Println(encode)

}
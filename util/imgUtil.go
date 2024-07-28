package util

import (
	"image/png"
	"os"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

func resizeImg(imgPath, out string) {
	// 打开 PNG 图片文件
	file, err := os.Open(imgPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 解码 PNG 图片
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	// 缩放图片到指定尺寸
	newImg := imaging.Resize(img, 59, 0, imaging.Lanczos)

	// 创建输出文件
	outFile, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, newImg)
	if err != nil {
		panic(err)
	}
}

func resizeImg2(imgPath, out string) {
	// 打开 PNG 图片文件
	file, err := os.Open(imgPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 解码 PNG 图片
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	// 缩放图片到指定尺寸
	newImg := resize.Resize(200, 0, img, resize.Lanczos3)

	// 创建输出文件
	outFile, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// 将缩放后的图片编码为 PNG 格式并保存
	err = png.Encode(outFile, newImg)
	if err != nil {
		panic(err)
	}
}

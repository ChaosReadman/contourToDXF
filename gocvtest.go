package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
)

func main() {
	// ファイル読み込み
	imgPath := "PC250001.jpg"
	img := gocv.IMRead(imgPath, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Could not read image %s\n", imgPath)
		os.Exit(1)
	}
	// グレースケール作成
	imgGray := gocv.NewMat()
	gocv.CvtColor(img, &imgGray, gocv.ColorBGRToGray)
	// ここまでを保存
	outPath := "gray.jpeg"
	if ok := gocv.IMWrite(outPath, imgGray); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
}

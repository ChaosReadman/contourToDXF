package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
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
	grayPath := "gray.jpeg"
	if ok := gocv.IMWrite(grayPath, imgGray); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
	imgTreshold := gocv.NewMat()
	gocv.Threshold(imgGray, &imgTreshold, 127, 255, 0)
	// ここまでを保存
	thresholdPath := "Threshold.jpeg"
	if ok := gocv.IMWrite(thresholdPath, imgTreshold); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
	// 輪郭を抽出
	contours := gocv.FindContours(imgTreshold, gocv.RetrievalTree, gocv.ChainApproxSimple)
	// 元のイメージに書き込む
	gocv.DrawContours(&img, contours, -1, color.RGBA{255, 0, 0, 0}, 3)
	// ここまでを保存
	contourdPath := "Contour.jpeg"
	if ok := gocv.IMWrite(contourdPath, img); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
}

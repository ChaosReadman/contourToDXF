package main

import (
	"fmt"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func writeImg(file string, img *gocv.Mat) {
	if ok := gocv.IMWrite(file, *img); !ok {
		fmt.Printf("Failed to write image: %s\n", file)
		os.Exit(1)
	}
}

func readImg(file string) gocv.Mat {
	img := gocv.IMRead(file, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Could not read image %s\n", file)
		os.Exit(1)
	}
	return img
}

func main() {
	// ファイル読み込み
	img := readImg("PC250001.jpg")

	// グレースケール作成
	imgGray := gocv.NewMat()
	gocv.CvtColor(img, &imgGray, gocv.ColorBGRToGray)
	// ここまでを保存
	writeImg("gray.jpg", &imgGray)

	imgTreshold := gocv.NewMat()
	gocv.Threshold(imgGray, &imgTreshold, 127, 255, 0)
	// ここまでを保存
	writeImg("Threshold.jpg", &imgTreshold)

	// 輪郭を抽出
	contours := gocv.FindContours(imgTreshold, gocv.RetrievalTree, gocv.ChainApproxSimple)
	// 元のイメージに書き込む
	gocv.DrawContours(&img, contours, -1, color.RGBA{255, 0, 0, 0}, 3)
	// ここまでを保存
	writeImg("Contour.jpg", &img)
}

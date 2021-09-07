package main

import (
	"image"
	"image/color"

	"image/jpeg"
	"log"
	"os"

	"gocv.io/x/gocv"
)

var w = 640
var h = 480

func main() {
	// 描画領域作成
	atom := gocv.NewMatWithSize(h, w, gocv.MatTypeCV8UC3)
	defer atom.Close()

	// 色定義
	black := color.RGBA{0, 0, 0, 0}
	blue := color.RGBA{0, 0, 255, 0}
	red := color.RGBA{255, 0, 0, 0}
	white := color.RGBA{255, 255, 255, 0}
	yellow := color.RGBA{255, 255, 0, 0}

	// 短径描画
	gocv.Rectangle(&atom, image.Rect(0, 0, 640, 480), white, -1)
	gocv.Rectangle(&atom, image.Rect(10, 10, 100, 100), black, 5)
	gocv.Rectangle(&atom, image.Rect(200, 15, 300, 350), blue, 5)
	gocv.Rectangle(&atom, image.Rect(10, 400, 600, 470), yellow, -1) // -1の場合は塗りつぶし
	gocv.Rectangle(&atom, image.Rect(40, 350, 550, 450), red, -1)    // -1の場合は塗りつぶし

	// 多角形描画
	// ps := make([][]image.Point, 1)
	// p := make([]image.Point, 4)
	// p[0] = image.Pt(int(0.1*float64(w)), int(0.1*float64(h)))
	// p[1] = image.Pt(int(0.1*float64(w)), int(0.4*float64(h)))
	// p[2] = image.Pt(int(0.3*float64(w)), int(0.2*float64(h)))
	// p[3] = image.Pt(int(0.3*float64(w)), int(0.1*float64(h)))
	// ps[0] = p
	// gocv.DrawContours(&atom, ps, 0, red, -1) // -1の場合は塗りつぶし

	file_path := "aiueo.jpg"

	// 保存
	//gocv.IMWrite(file_path, atom)

	// 画像に変換
	image, err := atom.ToImage()
	if err != nil {
		log.Println(err)
		return
	}

	// file出力
	file, errOs := os.Create(file_path)
	if errOs != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// jpeg形式で出力
	if err := jpeg.Encode(file, image, &jpeg.Options{100}); err != nil {
		log.Println(err)
		return
	}

	/*
		// bytes変換
		var jpegBytes []byte
		buf := bytes.NewBuffer(jpegBytes)
		if err := jpeg.Encode(buf, image, nil); err != nil {
			log.Println(err)
		}
		jpegBytes = buf.Bytes()
	*/

}

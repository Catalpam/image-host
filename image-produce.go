package main

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
	"path"
	"strings"
)

func WaterMark(filepath string)  {
	imageBaseName := path.Base(filepath)
	imageExt := path.Ext(filepath)
	imagePrefix := strings.TrimSuffix(imageBaseName, imageExt)

	println("----------WaterMark----------")
	buffer, err := bimg.Read(filepath)
	if err != nil {
		println("error1")
		fmt.Fprintln(os.Stderr, err)
	}

	watermark := bimg.Watermark{
		Text:       Config.Watermark.Word.Content,
		Opacity:    0.3,
		Width:      200,
		DPI:        100,
		Margin:     150,
		Font:       "sans bold 12",
		Background: bimg.Color{130, 130, 130},
	}

	if bimg.NewImage(buffer).Type() == "png"{
		buffer = Convert(buffer)
	}

	newImage, err := bimg.NewImage(buffer).Watermark(watermark)
	if err != nil {
		println("error2")
		fmt.Fprintln(os.Stderr, err)
	}

	errWrite := bimg.Write(Config.Dir+imagePrefix+"-w."+bimg.NewImage(newImage).Type(), newImage)
	if errWrite != nil {
		println("error3")
		panic("WaterMark Error")
	}
}

func Slim(filepath string)  {
	buffer, err := bimg.Read(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	imageBaseName := path.Base(filepath)
	imageExt := path.Ext(filepath)
	imagePrefix := strings.TrimSuffix(imageBaseName, imageExt)
	println(imagePrefix)

	size, err := bimg.NewImage(buffer).Size()
	if size.Width * size.Height > Config.Slim.Pixels {
		fmt.Println("The number of pixels is over", Config.Slim.Pixels, ".")
		newImage, err := bimg.NewImage(buffer).Resize(800, 600)
		if err != nil {
			println("size_error_2")
			fmt.Fprintln(os.Stderr, err)
		}
		bimg.Write(Config.Dir+imagePrefix+"-slim."+imageExt, newImage)
	}
}

func Convert(image []byte) []byte{

	newImage, err := bimg.NewImage(image).Convert(bimg.JPEG)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if bimg.NewImage(newImage).Type() == "jpeg" {
		fmt.Fprintln(os.Stderr, "The image was converted into jpeg")
	}

	return newImage
}
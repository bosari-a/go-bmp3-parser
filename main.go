package main

import (
	"encoding/binary"
	"fmt"

	"github.com/bosari-a/go-bmp3-parser/bmp3"
	"github.com/bosari-a/go-bmp3-parser/parse24bit"
)

func main() {
	var bh bmp3.BITMAPHEADER
	var bi bmp3.BITMAPINFOHEADER
	parse24bit.Parse24bitData("./FLAG_B24.BMP", &bh, &bi)
	sig := fmt.Sprintf("%c%c", (*bh.HEADER["signature"])[0], (*bh.HEADER["signature"])[1])
	fmt.Printf("%s\n", sig)
	fmt.Printf("%v\n", binary.LittleEndian.Uint16(*bi.INFOHEADER["bitsPerPixel"]))

	// image := *res.ImageData
	// for i := int(res.Height) - 1; i > 0; i-- {
	// 	for j := 0; j < int(res.Width); j++ {
	// 		fmt.Printf("\x1b[38;2;%v;%v;%vm%%\x1b[0m", (*image[i])[j].Red, (*image[i])[j].Green, (*image[i])[j].Blue)
	// 	}
	// 	println()
	// }


}

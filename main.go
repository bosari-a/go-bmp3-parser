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
	parse24bit.Parse24bitData("./assets/FLAG_B24.BMP", &bh, &bi)
	fmt.Printf("%c%c\n", (*bh.HEADER["signature"])[0], (*bh.HEADER["signature"])[1])
	fmt.Printf("%v\n", binary.LittleEndian.Uint16(*bi.INFOHEADER["bitsPerPixel"]))
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bosari-a/go-bmp3-parser/bmp3"
)

func main() {
	var bh bmp3.BITMAPHEADER
	bh.HEADER = make(map[string]*[]byte)
	bh.HEADERBYTES = map[string]uint8{
		"signature":  2,
		"fileSize":   4,
		"reserved":   4,
		"dataOffset": 4,
	}
	fd, err := os.OpenFile("./assets/FLAG_B24.BMP", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	var offset int64
	bh.ParseHeader(fd, &offset)
	fmt.Printf("%c%c\n", (*bh.HEADER["signature"])[0], (*bh.HEADER["signature"])[1])
}

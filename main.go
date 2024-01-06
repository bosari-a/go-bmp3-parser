package main

import (
	"log"
	"os"

	"github.com/bosari-a/go-bmp3-parser/bmp3"
)

func main() {
	fd, err := os.OpenFile("./assets/FLAG_B24.BMP", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	var bh bmp3.BITMAPHEADER = make(map[string]*[]byte)
	var offset uint8
	for k, v := range bmp3.HEADERBYTES {
		var b = make([]byte, v)
		fd.ReadAt(b, int64(offset))
		bh[k] = &b
		offset += v
	}
	
	
}

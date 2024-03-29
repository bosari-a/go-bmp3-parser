package bmp3

import (
	"os"
)

// constants
const HEADERLENGTH = 4
const INFOHEADERLENGTH = 11

// Bitmap header format
type BITMAPHEADER struct {
	HEADER      map[string]*[]byte
	HEADERBYTES [4]int64
	HEADERPROPS [4]string
}

// function to parse header
func (bh *BITMAPHEADER) ParseHeader(fd *os.File) {
	for i := 0; i < HEADERLENGTH; i++ {
		k := bh.HEADERPROPS[i]
		v := bh.HEADERBYTES[i]
		b := make([]byte, v)
		fd.Read(b)
		bh.HEADER[k] = &b
	}
}

// bitmap info header format
type BITMAPINFOHEADER struct {
	INFOHEADER      map[string]*[]byte
	INFOHEADERBYTES [INFOHEADERLENGTH]int64
	INFOHEADERPROPS [INFOHEADERLENGTH]string
}

// function to parse info header
func (bi *BITMAPINFOHEADER) ParseInfoHeader(fd *os.File) {
	for i := 0; i < INFOHEADERLENGTH; i++ {
		k := bi.INFOHEADERPROPS[i]
		v := bi.INFOHEADERBYTES[i]
		b := make([]byte, v)
		fd.Read(b)
		bi.INFOHEADER[k] = &b
	}
}

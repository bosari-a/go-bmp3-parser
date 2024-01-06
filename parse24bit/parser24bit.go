package parse24bit

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/bosari-a/go-bmp3-parser/bmp3"
)

// RGB colors according to the specification
// colors are stored backwards: BGR
type RGB struct {
	Blue  uint8
	Green uint8
	Red   uint8
}

// This will be the return value of Parse24bitData function
type Result struct {
	ImageData *([](*[]RGB))
	Width     int16
	Height    int16
}

// Parses the BMP file
func Parse24bitData(file string, bh *bmp3.BITMAPHEADER, bi *bmp3.BITMAPINFOHEADER) (Result, error) {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		return Result{}, err
	}
	// Parsing BITMAPHEADER
	bh.HEADER = make(map[string]*[]byte)
	bh.HEADERBYTES = [bmp3.HEADERLENGTH]int64{2, 4, 4, 4}
	bh.HEADERPROPS = [bmp3.HEADERLENGTH]string{"signature", "fileSize", "reserved", "dataOffset"}
	bh.ParseHeader(fd)

	// Parsing BITMAPINFOHEADER
	bi.INFOHEADERBYTES = [bmp3.INFOHEADERLENGTH]int64{
		4,
		4,
		4,
		2,
		2,
		4,
		4,
		4,
		4,
		4,
		4,
	}
	bi.INFOHEADERPROPS = [bmp3.INFOHEADERLENGTH]string{
		"size",
		"width",
		"height",
		"planes",
		"bitsPerPixel",
		"compression",
		"imageSize",
		"xpixelsPerM",
		"ypixelsPerM",
		"colorsUsed",
		"importantColors",
	}
	bi.INFOHEADER = make(map[string]*[]byte)
	bi.ParseInfoHeader(fd)

	// checking if bitmap is 24bit RGB
	// Signature must be "BM"
	// compression must be 0
	// Data offset must be 54
	// Bits per pixel must be 24
	signature := fmt.Sprintf("%c%c", (*bh.HEADER["signature"])[0], (*bh.HEADER["signature"])[1])
	compression := int16(binary.LittleEndian.Uint16(*bi.INFOHEADER["compression"]))
	dataOffset := int16(binary.LittleEndian.Uint16(*bh.HEADER["dataOffset"]))
	bitsPerPixel := int16(binary.LittleEndian.Uint16(*bi.INFOHEADER["bitsPerPixel"]))

	if signature != "BM" ||
		compression != 0 ||
		dataOffset != 54 ||
		bitsPerPixel != 24 {
		return Result{}, errors.New("unsupported format")
	}

	w := int16(binary.LittleEndian.Uint16(*bi.INFOHEADER["width"]))
	h := int16(binary.LittleEndian.Uint16(*bi.INFOHEADER["height"]))
	if h < 0 {
		h = -1 * h
	}
	padding := (4 - ((int(w) * 3) % 4)) % 4

	imageData := make([](*[]RGB), h)

	for i := 0; i < int(h); i++ {
		bytesRow := make([]byte, int(w)*3)
		fd.Read(bytesRow)
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, bytesRow)
		row := make([]RGB, w)
		binary.Read(buf, binary.LittleEndian, &row)
		imageData[i] = &row
		fd.Seek(int64(padding), io.SeekCurrent)
	}
	return Result{
		ImageData: &imageData,
		Height:    h,
		Width:     w,
	}, nil
}

package parse24bit

import (
	"log"
	"os"

	"github.com/bosari-a/go-bmp3-parser/bmp3"
)

func Parse24bitData(file string, bh *bmp3.BITMAPHEADER, bi *bmp3.BITMAPINFOHEADER) {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	bh.HEADER = make(map[string]*[]byte)
	bh.HEADERBYTES = [bmp3.HEADERLENGTH]int64{2, 4, 4, 4}
	bh.HEADERPROPS = [bmp3.HEADERLENGTH]string{"signature", "fileSize", "reserved", "dataOffset"}
	bh.ParseHeader(fd)

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
}

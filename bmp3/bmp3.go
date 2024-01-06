package bmp3

type BITMAPHEADER = map[string]*[]byte

var HEADERBYTES = map[string]uint8{
	"signature":  2,
	"fileSize":   4,
	"reserved":   4,
	"dataOffset": 4,
}

type BITMAPINFOHEADER = map[string]*[]byte

var INFOHEADERBYTES = map[string]uint8{
	"size":            4,
	"width":           4,
	"height":          4,
	"planes":          2,
	"bitsPerPixel":    2,
	"compression":     4,
	"imageSize":       4,
	"xpixelsPerM":     4,
	"ypixelsPerM":     4,
	"colorsUsed":      4,
	"importantColors": 4,
}

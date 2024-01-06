package bmp3

type BITMAPHEADER struct {
	Signature  [2]uint8
	FileSize   uint32
	Reserved   uint32
	DataOffset uint32
}

type BITMAPINFOHEADER struct {
	Size            uint32
	Width           int32
	Height          int32
	Planes          int16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XpixelsPerM     uint32
	YpixelsPerM     uint32
	ColorsUsed      uint32
	ImportantColors uint32
}

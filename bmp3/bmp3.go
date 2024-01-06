package bmp3

import "os"

type BITMAPHEADER struct {
	HEADER      map[string]*[]byte
	HEADERBYTES map[string]uint8
}

func (bh *BITMAPHEADER) ParseHeader(fd *os.File, offset *int64) {
	for k, v := range (*bh).HEADERBYTES {
		b := make([]byte, v)
		fd.ReadAt(b, *offset)
		bh.HEADER[k] = &b
		*offset += int64(v)
	}
}

type BITMAPINFOHEADER struct {
	INFOHEADER      map[string]*[]byte
	INFOHEADERBYTES map[string]uint8
}

func (bi *BITMAPINFOHEADER) ParseInfoHeader(fd *os.File, offset *int64) {
	for k, v := range bi.INFOHEADERBYTES {
		b := make([]byte, v)
		fd.ReadAt(b, *offset)
		bi.INFOHEADER[k] = &b
		*offset += int64(v)
	}
}

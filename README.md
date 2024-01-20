# Windows 3.x 24 bit BMP Parser for Golang

## Description

This library provides a function to parse 24 bit Windows 3.x `BMP` files.

## Usage

If you want a little more detail regarding the exported functions you can read more in [exports](#exports)

Usage is simple and goes as follows:

```go
import (
	"github.com/bosari-a/go-bmp3-parser/bmp3"
	"github.com/bosari-a/go-bmp3-parser/parser24bit"
)
    // initialize metadata structs
	var bh bmp3.BITMAPHEADER
	var bi bmp3.BITMAPINFOHEADER
    // parse!
	res, err := parser24bit.Parse24bitData("myimage.bmp", &bh, &bi)
	if err != nil {
		log.Fatal(err)
	}

	w := res.Width
	h := res.Height
	imageData := *(res.ImageData)
```

## Exports

Image files usually have some metadata "offset" before the actual color/rgba/pixel values that represent what you see.

Within this package is a module `bmp3` that defines the structure of that metadata. You will need to initialize the bitmap header and infoheader before using the parser so this module is important.

The following is exported by the `bmp3` module:
This also assumes knowledge about the `BMP` specification since for normal usage (when you only want RGB rows of data) you don't need to go such lengths.

```go
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
func (bh *BITMAPHEADER) ParseHeader(fd *os.File)

// bitmap info header format
type BITMAPINFOHEADER struct {
	INFOHEADER      map[string]*[]byte
	INFOHEADERBYTES [INFOHEADERLENGTH]int64
	INFOHEADERPROPS [INFOHEADERLENGTH]string
}

// function to parse info header
func (bi *BITMAPINFOHEADER) ParseInfoHeader(fd *os.File)
```


The parser module function is defined as follows:

```go
func Parse24bitData(file string, bh *bmp3.BITMAPHEADER, bi *bmp3.BITMAPINFOHEADER) (Result, error)
```
It takes in `file` which is a path to the `BMP` file you want to parse, `bh` which is a BITMAPHEADER pointer, and `bi` which is a BITMAPINFOHEADER pointer. On success it returns `Result` with the error being `nil`.

`return` value `Result` is defined as (after `RGB` struct):
```go
// RGB colors according to the specification
// colors are stored backwards: BGR
type RGB struct {
	Blue  uint8
	Green uint8
	Red   uint8
}

type Result struct {
	ImageData *([](*[]RGB))
	Width     int16
	Height    int16
}
```

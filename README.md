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

This code snippet is from one of my other projects that uses this parser to generate terminal ansi art: [github.com/bosari-a/image-to-ansi-go](github.com/bosari-a/image-to-ansi-go)

That ansi art project I mentioned is a good example to illustrate how this parser work. The art generation is quite simple, the rest of the code looks like this:

```go
	for i := int(h) - 1; i > 0; i-- {
		fmt.Print("\t")
		for j := 0; j < int(w); j++ {
			row := (*imageData[i])
			pixel := row[j]
			fmt.Printf("\x1b[38;2;%v;%v;%vm%%\x1b[0m", pixel.Red, pixel.Green, pixel.Blue)
		}
		fmt.Println("")
	}
```

Ignoring the ansi aspect for a second, notice that you need to start looping in reverse height. That's because according to the `BMP` specification scan lines (rows of pixels) are stored bottom to top. My parser adheres to the specification in that regard.

Another thing to keep in mind: you need to pass in `bh` and `bi` as parameters because I believe a proper parser should provide all data. This includes meta data which is available in those variables that you are required to initialize. They are mutated by the parser function as it reads data from the `BMP` file.


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

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint16(19778))

	d := make([]byte, 2)
	buf.Read(d)

	fmt.Printf("%c%c\n", d[0], d[1])
}

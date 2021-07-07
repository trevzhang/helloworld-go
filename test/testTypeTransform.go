package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
)

func strToByte(src string) []byte {
	// strSlice := []string{src}
	buffer := &bytes.Buffer{}
	err := gob.NewEncoder(buffer).Encode(src)
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

func uint64ToByte(src uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, src)
	return buf
}

func main() {
	var u uint64
	u = 128
	byteValue := uint64ToByte(u)
	fmt.Println(byteValue)
	str := "128"
	byteValue = strToByte(str)
	fmt.Println(byteValue)
}

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Little-endian

// Value types
const I32 = 0x7F
const I64 = 0x7E
const F32 = 0x7D
const F64 = 0x7C

// Module sections
const CUSTOM_SECTION = 0x00 // Debugging. Ignored by wasm, so ignored by us.
const TYPE_SECTION = 0x01
const IMPORT_SECTION = 0x02
const FUNCTION_SECTION = 0x03
const TABLE_SECTION = 0x04
const MEMORY_SECTION = 0x05
const GLOBAL_SECTION = 0x06
const EXPORT_SECTION = 0x07
const START_SECTION = 0x08
const ELEMENT_SECTION = 0x09
const CODE_SECTION = 0x10
const DATA_SECTION = 0x11

// TODO: Validate float type - IEEE 754-2008 vs go bytes.

func Magic() []byte {
	return []byte{0x00, 0x61, 0x73, 0x6D}
}

func Version() []byte {
	return []byte{0x01, 0x00, 0x00, 0x00}
}

func EncodeI32(i uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, i)
	return bs
}

func EncodeI64(i uint64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, i)
	return bs
}

// Accepts bool, int32, int64, etc.
func EncodeNumber(i interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		fmt.Println("Binary encoding of number failed", err)
	}
	fmt.Printf("%x", buf.Bytes())
	return buf.Bytes()
}

func EncodeVector(data []byte) []byte {
	// TODO - verify order in output.
	size := EncodeI32(uint32(len(data)))
	return append(size[:], data[:]...)
}

func main() {
	var a = 0x0F
	var b = uint32(32)
	fmt.Println(EncodeNumber(uint32(32)))
	fmt.Printf("%x", EncodeI32(b))
	fmt.Printf("Go wasm bin {%d}\n", a)
}

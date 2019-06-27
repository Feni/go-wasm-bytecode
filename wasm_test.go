package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMagic(t *testing.T) {
	got := Magic()
	correct := []byte{0x00, 0x61, 0x73, 0x6D}
	if !reflect.DeepEqual(got, correct) {
		fmt.Println("Wrong magic")
	} else {
		fmt.Println("Right magic")
	}
}

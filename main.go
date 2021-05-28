package main

import (
	"log"
	"fmt"
	"github.com/tarm/serial"
)
func main() {
	c := &serial.Config{Name: "COM5", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
			log.Fatal(err)
	}
	_, err = s.Write([]byte{0xe3})
	if err != nil {
			log.Fatal(err)
	}

	// データ受信のプログラム
	buf := make([]byte, 128)
	n, err := s.Read(buf)
	fmt.Printf("受信データ：%v\n", n)
	if err != nil {
			log.Fatal(err)
	}
}
package main

import (
	"fmt"
	"log"
	"github.com/tarm/serial"
)

var data int

func main() {
	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
			log.Fatal(err)
	}

	go SerialDataInput()
	go SerialWrite(s)

	for {}
}

func SerialDataInput() {
	for {
		fmt.Print("データ入力：")
		fmt.Scan(&data)
	}
}

func SerialWrite(s *serial.Port) {
	for {
		_, err := s.Write([]byte{byte(data)})
		if err != nil {
				log.Fatal(err)
		}
	}
}
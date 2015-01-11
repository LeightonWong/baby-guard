package main

import (
	//"bytes"
	//"encoding/binary"
	"bufio"
	"log"

	"github.com/samofly/serial"
)

func main() {
	// Open the port.
	port, err := serial.Open("/dev/ttyACM0", 9600)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	r := bufio.NewReader(port)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Fatalf("Read error : %v", err)
		}
		log.Printf("Read line: %s\n", line)
	}
	// Write 4 bytes to the port.
	/*
		b := []byte{0x00, 0x01, 0x02, 0x03}
		n, err := port.Write(b)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}

		fmt.Println("Wrote", n, "bytes.")
	*/
	/*
		b := make([]byte, 16)
		for {
			n, err := port.Read(b)
			if err != nil {
				log.Fatalf("Read error: %v", err)
			}
			if n > 0 {
				log.Printf("Read from serial num : %d, data: %v", n, b)
				tem := string(b[:n-1])
				log.Printf("Read serial: %s", tem)
				b = make([]byte, 16)
			}
		}
	*/
}

package main

import (
	//"bytes"
	//"encoding/binary"
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/lnmx/serial"
)

func main() {
	// Open the port.
	port, err := serial.Open("/dev/ttyACM0", 9600)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	client := &http.Client{}
	r := bufio.NewReader(port)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Fatalf("Read error : %v", err)
		}
		//log.Printf("Received line: %s\n", line)
		r, err := client.PostForm("http://127.0.0.1:8080/record", url.Values{"temp": {string(line)}})
		if err != nil {
			log.Printf("Send request error: %s\n", err)
		}
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		log.Printf("Send line: %s result %s\n", line, body)
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

package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"math/rand"
	"net/http" //package for http based web programs
	"strings"
	"time"
)

var s io.ReadWriteCloser

func init() {
	var err error
	c := &serial.Config{
		Name: "COM3",
		Baud: 9600,
		//		ReadTimeout: time.Millisecond * 5,
	}

	s, err = serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
	}
	//  s.Close()

	rand.Seed(time.Now().UnixNano())

}

func SendToArduinoHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Send to Ardy")

	fmt.Fprintln(w, "<html>"+SendToArduino()+"</html>")
}

func ReceiveFromArduinoHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Receive From Ardy")

	fmt.Fprintln(w, "<html>"+ReceiveFromArduino()+"</html>")
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SendToArduino() string {
	randomID := randSeq(8)
	fmt.Fprintln(s, randomID)
	return "Sent data to Arduino. Id=" + randomID
}

func ReceiveFromArduino() string {
	// Receive reply
	reply := ReadArduino()

	return reply
}

func ReadArduino() string {
	keepreading := true
	data := ""

	scanner := bufio.NewScanner(s)
	for keepreading {
		scanner.Scan()
		line := scanner.Text()
		fmt.Println(line)

		if strings.Contains(line, "END") {
			keepreading = false
			fmt.Println("End of reading")
		}

		data += line + "<br>\n"

	}

	return data
}

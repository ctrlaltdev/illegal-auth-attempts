package main

import (
	"fmt"

	"github.com/nxadm/tail"
)

var (
	file = "/var/log/auth.log"
)

func tailer(c chan<- string) {
	t, err := tail.TailFile(file, tail.Config{Follow: true, ReOpen: true, MustExist: true})
	CheckErr(err)

	for line := range t.Lines {
		c <- line.Text
	}
}

func dispatcher(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	done := make(chan bool)
	c := make(chan string, 2)

	go tailer(c)
	go dispatcher(c)

	<-done
}

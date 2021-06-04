package main

import (
	"github.com/nxadm/tail"
)

var (
	file = "./auth.log"
	// file = "/var/log/auth.log"

	done = make(chan bool)

	logs = make(chan string, 12)

	pamLogs          = make(chan string, 2)
	failedPwdLogs    = make(chan string, 2)
	disconnectLogs   = make(chan string, 2)
	disconnectedLogs = make(chan string, 2)
	connClosedLogs   = make(chan string, 2)
	invalidUserLogs  = make(chan string, 2)

	events = make(chan *Event, 12)
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
		logType := ParseType(msg)

		switch logType {
		case "pam":
			pamLogs <- msg
			break
		case "failedpwd":
			failedPwdLogs <- msg
			break
		case "disconnect":
			disconnectLogs <- msg
			break
		case "connclosed":
			connClosedLogs <- msg
			break
		case "invaliduser":
			invalidUserLogs <- msg
			break
		default:
			break
		}
	}
}

func main() {

	go tailer(logs)
	go dispatcher(logs)

	go PamParser(pamLogs, events)
	go FailedPwdParser(failedPwdLogs, events)
	go DisconnectParser(disconnectLogs, events)
	go DisconnectedParser(disconnectedLogs, events)
	go ConnClosedParser(connClosedLogs, events)
	go InvalidUserParser(invalidUserLogs, events)

	go Printer(events)

	<-done
}

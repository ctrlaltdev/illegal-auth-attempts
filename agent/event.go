package main

import (
	"fmt"
	"regexp"
)

type Event struct {
	Type    string
	GroupID string
	User    string
	Source  string
}

type LogTypeParser struct {
	logType string
	re      *regexp.Regexp
}

func ParseType(log string) string {
	var types = []LogTypeParser{
		{"pam", regexp.MustCompile(`pam_unix\(sshd\:auth\)\:`)},
		{"failedpwd", regexp.MustCompile(`Failed password for invalid user`)},
		{"disconnect", regexp.MustCompile(`Received disconnect from`)},
		{"disconnected", regexp.MustCompile(`Disconnected from authenticating user`)},
		{"connclosed", regexp.MustCompile(`Connection closed by invalid user`)},
		{"invaliduser", regexp.MustCompile(`Invalid user`)},
	}

	for _, t := range types {
		matches := t.re.MatchString(log)
		if matches {
			return t.logType
		}
	}

	return ""
}

func PamParser(c <-chan string, d chan<- *Event) {
	var (
		event = Event{Type: "pam"}

		reGroupID = regexp.MustCompile(`sshd\[(\d+)]\:`)
		reUser    = regexp.MustCompile(`user\=([\w\d\-\_]+)`)
		reSource  = regexp.MustCompile(`rhost\=([\w\d\_\-\.\:]+)`)
	)

	for {
		msg := <-c

		mGroupID := reGroupID.FindStringSubmatch(msg)
		mUser := reUser.FindStringSubmatch(msg)
		mSource := reSource.FindStringSubmatch(msg)

		if len(mGroupID) == 2 {
			event.GroupID = mGroupID[1]
		}

		if len(mUser) == 2 {
			event.User = mUser[1]
		}

		if len(mSource) == 2 {
			event.Source = mSource[1]
		}

		d <- &event
	}
}

func FailedPwdParser(c <-chan string, d chan<- *Event) {
	var (
		event = Event{Type: "failedpwd"}

		reGroupID = regexp.MustCompile(`sshd\[(\d+)]\:`)
		reUser    = regexp.MustCompile(`Failed password for invalid user ([\w\d\-\_]+) from`)
		reSource  = regexp.MustCompile(`from ([\w\d\_\-\.\:]+) port`)
	)

	for {
		msg := <-c

		mGroupID := reGroupID.FindStringSubmatch(msg)
		mUser := reUser.FindStringSubmatch(msg)
		mSource := reSource.FindStringSubmatch(msg)

		if len(mGroupID) == 2 {
			event.GroupID = mGroupID[1]
		}

		if len(mUser) == 2 {
			event.User = mUser[1]
		}

		if len(mSource) == 2 {
			event.Source = mSource[1]
		}

		d <- &event
	}
}

func DisconnectParser(c <-chan string, d chan<- *Event) {
	var (
		event = Event{Type: "disconnect"}

		reGroupID = regexp.MustCompile(`sshd\[(\d+)]\:`)
		reSource  = regexp.MustCompile(`Received disconnect from ([\w\d\_\-\.\:]+) port`)
	)

	for {
		msg := <-c

		mGroupID := reGroupID.FindStringSubmatch(msg)
		mSource := reSource.FindStringSubmatch(msg)

		if len(mGroupID) == 2 {
			event.GroupID = mGroupID[1]
		}

		if len(mSource) == 2 {
			event.Source = mSource[1]
		}

		d <- &event
	}
}

func DisconnectedParser(c <-chan string, d chan<- *Event) {
	var (
		event = Event{Type: "disconnected"}

		reGroupID = regexp.MustCompile(`sshd\[(\d+)]\:`)
		reUser    = regexp.MustCompile(`Disconnected from invalid user ([\w\d\-\_]+)`)
		reSource  = regexp.MustCompile(`([\w\d\_\-\.\:]+) port \d+ \[preauth\]`)
	)

	for {
		msg := <-c

		mGroupID := reGroupID.FindStringSubmatch(msg)
		mUser := reUser.FindStringSubmatch(msg)
		mSource := reSource.FindStringSubmatch(msg)

		if len(mGroupID) == 2 {
			event.GroupID = mGroupID[1]
		}

		if len(mUser) == 2 {
			event.User = mUser[1]
		}

		if len(mSource) == 2 {
			event.Source = mSource[1]
		}

		d <- &event
	}
}

func ConnClosedParser(c <-chan string, d chan<- *Event) {
	var (
		event = Event{Type: "connclosed"}

		reGroupID = regexp.MustCompile(`sshd\[(\d+)]\:`)
		reUser    = regexp.MustCompile(`Connection closed by invalid user ([\w\d\-\_]+)`)
		reSource  = regexp.MustCompile(`([\w\d\_\-\.\:]+) port \d+ \[preauth\]`)
	)

	for {
		msg := <-c

		mGroupID := reGroupID.FindStringSubmatch(msg)
		mUser := reUser.FindStringSubmatch(msg)
		mSource := reSource.FindStringSubmatch(msg)

		if len(mGroupID) == 2 {
			event.GroupID = mGroupID[1]
		}

		if len(mUser) == 2 {
			event.User = mUser[1]
		}

		if len(mSource) == 2 {
			event.Source = mSource[1]
		}

		d <- &event
	}
}

func InvalidUserParser(c <-chan string, d chan<- *Event) {
	var (
		event = Event{Type: "invaliduser"}

		reGroupID = regexp.MustCompile(`sshd\[(\d+)]\:`)
		reUser    = regexp.MustCompile(`Invalid user ([\w\d\-\_]+) from`)
		reSource  = regexp.MustCompile(`from ([\w\d\_\-\.\:]+) port`)
	)

	for {
		msg := <-c

		mGroupID := reGroupID.FindStringSubmatch(msg)
		mUser := reUser.FindStringSubmatch(msg)
		mSource := reSource.FindStringSubmatch(msg)

		if len(mGroupID) == 2 {
			event.GroupID = mGroupID[1]
		}

		if len(mUser) == 2 {
			event.User = mUser[1]
		}

		if len(mSource) == 2 {
			event.Source = mSource[1]
		}

		d <- &event
	}
}

func Printer(c <-chan *Event) {
	for {
		event := <-c
		fmt.Printf("%+v\n", *event)
	}
}

func ParseEntry(line string) *Event {
	var event Event

	reGroupID := regexp.MustCompile(`sshd\[(\d+)]\:`)
	resGroupID := reGroupID.FindStringSubmatch(line)
	if len(resGroupID) == 2 {
		event.GroupID = resGroupID[1]
	}

	reUser := regexp.MustCompile(`user\=([\w\d\-\_]+)`)
	resUser := reUser.FindStringSubmatch(line)
	if len(resUser) == 2 {
		event.User = resUser[1]
	}

	reSource := regexp.MustCompile(`rhost\=([\w\d\_\-\.\:]+)`)
	resSource := reSource.FindStringSubmatch(line)
	if len(resSource) == 2 {
		event.Source = resSource[1]
	}

	if event.User == "" && event.Source == "" {
		return nil
	}

	return &event
}

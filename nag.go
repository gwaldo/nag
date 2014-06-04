// #! `which go`

// nag - a command-line reminder and timer

// Usage:
// 	$ nag <integer> [ s|m|h ] [ -m <message> ] [ -s,  --silent ]

// Options:
// 	time (in [ seconds | minutes | hours ])
// 	message (optional)
// 	silent (do not `say` the command)

package main

import (
	"fmt"
	"log"
	"flag"
	"time"
	"runtime"
	"os/exec"
)

var seconds	= flag.Int("sec", 0, "Number of seconds to wait before Nagging.")
var minutes	= flag.Int("min", 0, "Number of minutes to wait before Nagging.")
var message	= flag.String("msg", "HEY!", "Message to Nag you with.")
var silent	= flag.Bool("silent", false, "Do not speak the string (Mac Only.)")
var s		= flag.Bool("s", false, "Do not speak the string (Mac Only.)")

func speak(m *string) {
	if runtime.GOOS == "darwin" {
		speak := exec.Command("say", *message)
		err := speak.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	flag.Parse()

	/*
	// Arch Details
	fmt.Println("GOARCH: ", runtime.GOARCH)
	fmt.Println("GOOS:   ", runtime.GOOS)
	*/

	/*
	// Diag
	fmt.Println("Nagging in this many seconds: ", *seconds)
	fmt.Println("Nagging in this many minutes: ", *minutes)
	fmt.Println("Nagging with this message: ", *message)
	*/
	
	time.Sleep(time.Duration(*seconds) * time.Second)
	time.Sleep(time.Duration(*minutes) * time.Minute)
	fmt.Println()
	
	if !*silent && !*s {
		speak(message)
	}
	fmt.Println(*message)
}

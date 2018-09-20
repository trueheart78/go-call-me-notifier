package main

import (
	"log"
	"time"

	"github.com/deckarep/gosx-notifier"
)

func main() {
	// check if MacOS 10.9+ and exit with proper output if not
	// else continue
	for i := 0; i < 5; i++ {
		notify2(true)
		time.Sleep(2 * time.Second)
	}
	for i := 0; i < 2; i++ {
		notify2(false)
		time.Sleep(4 * time.Second)
	}
}

func notify2(emergency bool) {
	t := "emergency"
	note := gosxnotifier.NewNotification("[" + t + "]")
	note.Group = "com.unique.alexa.alerter"
	note.Sender = "com.apple.Safari"
	if emergency {
		note.Title = "Lynda Alert!"
		note.Subtitle = "It's an emergency"
		note.AppIcon = "assets/icons/rotating_light.png"
		note.ContentImage = "assets/icons/rotating_light.png"
		note.Sound = gosxnotifier.Sosumi
	} else {
		note.Title = "Lynda Ping"
		note.Subtitle = "When you get time..."
		note.ContentImage = "assets/icons/sparkling_heart.png"
		note.AppIcon = "assets/icons/sparkling_heart.png"
	}
	err := note.Push()

	//If necessary, check error
	if err != nil {
		log.Println("Uh oh!")
	}
}

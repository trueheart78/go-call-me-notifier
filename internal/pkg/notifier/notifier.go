package notifier

import (
	"log"

	"github.com/deckarep/gosx-notifier"
)

// Emergency creates and pushes a new emergency notification
func Emergency() (bool, error) {
	notification := new(true)
	return push(notification)
}

// NonEmergent creates and pushes a new non-emergent notification
func NonEmergent() (bool, error) {
	notification := new(false)
	return push(notification)
}

func push(notification *gosxnotifier.Notification) (bool, error) {
	err := notification.Push()

	//If necessary, check error
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return false, err
	}
	return true, nil
}

func new(emergency bool) *gosxnotifier.Notification {
	note := gosxnotifier.NewNotification("[notifier]")
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

	return note
}

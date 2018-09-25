package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/trueheart78/go-call-me-notifier/internal/pkg/config"
	"github.com/trueheart78/go-call-me-notifier/internal/pkg/notifier"
)

var cfg config.Config

func init() {
	if runtime.GOOS != "darwin" {
		fmt.Printf("Unsupported operating system: %v\n", runtime.GOOS)
		os.Exit(1)
	}
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Connecting to %v\n", cfg.RedisURL)
}

func main() {
	fmt.Println(runtime.GOOS)
}

func emergency() {
	for i := 0; i < 5; i++ {
		notifier.Emergency()

		time.Sleep(2 * time.Second)
	}
}

func nonemergent() {
	for i := 0; i < 2; i++ {
		notifier.NonEmergent()
		time.Sleep(4 * time.Second)
	}
}

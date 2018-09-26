package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/go-redis/redis"
	"github.com/trueheart78/go-call-me-notifier/internal/pkg/config"
	"github.com/trueheart78/go-call-me-notifier/internal/pkg/notifier"
)

func init() {
	if runtime.GOOS != "darwin" {
		fmt.Printf("Unsupported operating system: %v\n", runtime.GOOS)
		os.Exit(1)
	}
	_, err := config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	monitor()
}

func monitor() {
	cfg, _ := config.NewConfig()
	redisdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: cfg.RedisPassword,
		DB:       0, // use default DB
	})

	_, err := redisdb.Ping().Result()
	if err != nil {
		hasPass := (cfg.RedisPassword != "")
		fmt.Printf("Error: Unable to connect to redis [password: %v]\n  url: %v\n", hasPass, cfg.RedisURL)
		os.Exit(1)
	}

	fmt.Printf("Connected to redis @ %v\n", cfg.RedisURL)

	pubsub := redisdb.Subscribe(cfg.RedisChannels.Emergency, cfg.RedisChannels.NonEmergent)

	// Wait for confirmation that subscription is created before publishing anything.
	_, err = pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// When pubsub is closed channel is closed too.
	defer pubsub.Close()

	fmt.Printf("Listening on '%v' and '%v'...\n", cfg.RedisChannels.Emergency, cfg.RedisChannels.NonEmergent)
	// Consume messages.
	var t time.Time
	timeFormat := "Jan 2, 2006 at 3:04pm (MST)"
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		t = time.Now()

		fmt.Printf("Message received on '%v': %v [%v]\n", msg.Channel, msg.Payload, t.Format(timeFormat))
		if msg.Channel == cfg.RedisChannels.Emergency {
			emergency()
		} else if msg.Channel == cfg.RedisChannels.NonEmergent {
			nonemergent()
		}
	}
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

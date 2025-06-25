package main

import (
	"log"
	"time"

	"github.com/kigland/FlashPoint/sdk"
)

func main() {
	client := sdk.NewClient("http://127.0.0.1:8080", "1234567890")
	err := client.Set("test", "testValue", 10*time.Second, "txt", "text/plain")
	if err != nil {
		log.Fatal(err)
	}
}

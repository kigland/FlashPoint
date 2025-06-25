package main

import (
	"log"
	"os"
	"time"

	"github.com/kigland/FlashPoint/sdk"
)

const imgPath = "/Users/kevin/Downloads/Screenshot 2025-06-24 at 12.32.21.png"

func main() {
	client := sdk.NewClient("http://127.0.0.1:8080", "1234567890")

	bs, err := os.ReadFile(imgPath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.SetBinary("test", bs, 30*time.Second, "image/png")
	if err != nil {
		log.Fatal(err)
	}
}

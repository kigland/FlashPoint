package main

import (
	"log"
	"time"

	"github.com/kigland/FlashPoint/sdk"
)

type TestObj struct {
	Test    string `json:"test"`
	ObjName string `json:"obj_name"`
}

func main() {
	client := sdk.NewClient("http://127.0.0.1:8080", "1234567890")

	_, err := client.SetJSON("test", TestObj{
		Test:    "test",
		ObjName: "obj_name",
	}, 30*time.Second, "application/json")
	if err != nil {
		log.Fatal(err)
	}
}

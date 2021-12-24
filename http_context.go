package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	argument := os.Args[1]

	if len(argument) <= 2 {
		return
	}
	client := http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	req, err := http.NewRequestWithContext(ctx, "GET", argument, nil)

	if err != nil {
		log.Fatal("error timeout: %v", err)
	}

	defer cancel()

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("%v", err)
	}

	log.Println("StatusCode: ", res.StatusCode)

}

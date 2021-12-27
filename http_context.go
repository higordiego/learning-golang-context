package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	website := os.Args[1]

	timeout, err := strconv.ParseInt(os.Args[2], 10, 64)

	if len(website) <= 2 || err != nil {
		return
	}
	client := http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	req, err := http.NewRequestWithContext(ctx, "GET", website, nil)

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

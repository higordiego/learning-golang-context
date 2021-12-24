package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

func main() {
	// context

	ctx, cancel := context.WithCancel(context.Background())

	log.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				log.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)

	log.Println("num gortins 2:", runtime.NumGoroutine())
	log.Println("about cancel")

	cancel()

	time.Sleep(time.Second * 2)
	log.Println("num gortins 3:", runtime.NumGoroutine())
	log.Println("cancel sender")
}

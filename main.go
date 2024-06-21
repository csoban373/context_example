package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	sleepAndPrint(ctx, 2*time.Second, "hello world")
}

func sleepAndPrint(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err().Error())
	}

	// c := make(chan bool, 1)

	// go func() {
	// 	time.Sleep(d)
	// 	fmt.Println("example")
	// 	c <- true
	// }()

	// select {
	// case <-ctx.Done():
	// 	fmt.Println(ctx.Err())
	// case <-c:
	// 	fmt.Println("here")
	// }
}

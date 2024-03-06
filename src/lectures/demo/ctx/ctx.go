package main

import (
	"context"
	"fmt"
	"time"
)

func smapleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed", msg)
				return
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted", msg)
			}
		}
	}()
	return out
}

func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := smapleOperation(ctx, "webserver", 200)
	microservice := smapleOperation(ctx, "microservice", 500)
	database := smapleOperation(ctx, "database", 900)

MainLoop:
	for {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microservice:
			fmt.Println(val)
			fmt.Println("cancel context")
			cancelCtx()
			break MainLoop
		case val := <-database:
			fmt.Println(val)
		}

	}
	fmt.Println(<-database)
}

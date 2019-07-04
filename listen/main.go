package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// fin := make(chan struct{}, 1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		onePort(ctx)
		wg.Done()
	}()
	go func() {
		twoPorts(ctx)
		wg.Done()
	}()

	<-c
	cancel()
	wg.Wait()
}

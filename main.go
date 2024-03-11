package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"go-chiserver-demo/domain/newsFeed"
	"go-chiserver-demo/server/chiserver"
)

func main() {
	exitChannel := make(chan os.Signal)
	signal.Notify(exitChannel, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverAddress := "localhost:3000"

	feed := newsFeed.New()

	server := chiserver.NewServer(serverAddress, feed)
	server.Start(ctx)

	fmt.Println("Press Ctrl + C to shutdown")
	<-exitChannel
}

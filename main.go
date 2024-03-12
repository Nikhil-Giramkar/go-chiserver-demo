package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	newsFeed "go-chiserver-demo/domain"
	chiserver "go-chiserver-demo/server"
)

func main() {
	exitChannel := make(chan os.Signal)
	signal.Notify(exitChannel, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverAddress := "localhost:3000"

	//Dummy Data
	feed := *newsFeed.New()
	feed.Add(newsFeed.Item{
		Title: "First Title",
		Post:  "First Post",
	})

	//Setup server and handlers
	server := chiserver.NewServer(serverAddress, &feed)

	//Start listening
	server.Start(ctx)

	fmt.Println("Press Ctrl + C to shutdown")
	<-exitChannel
}

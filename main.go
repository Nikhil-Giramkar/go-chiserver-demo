package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	newsFeed "go-chiserver-demo/domain"
	myLogger "go-chiserver-demo/logger"
	chiserver "go-chiserver-demo/server"
)

func main() {
	exitChannel := make(chan os.Signal)
	signal.Notify(exitChannel, os.Interrupt)
	l := myLogger.Get()

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
	l.Info("Server with handlers established")

	//Start listening
	server.Start(ctx)
	l.Info("Server running...")

	fmt.Println("Press Ctrl + C to shutdown")
	<-exitChannel
}

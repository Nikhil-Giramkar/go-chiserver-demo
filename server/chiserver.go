package server

import (
	"context"
	"fmt"
	newsFeed "go-chiserver-demo/domain"
	handlers "go-chiserver-demo/handlers"
	myLogger "go-chiserver-demo/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type serverChi struct {
	router *chi.Mux
	server *http.Server
}

// Exposes a router and server
func NewServer(address string, feed *newsFeed.Repo) *serverChi {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	server := setUpHandlers(address, router, feed)
	return &serverChi{router, server}
}

// Get Handlers added here
func setUpHandlers(address string, router *chi.Mux, feed *newsFeed.Repo) *http.Server {

	//http://localhost:3000/newsfeed
	l := myLogger.Get()

	router.Get("/newsfeed", handlers.GetNews(feed))

	//Pass JSON in request body at bove URL
	router.Post("/newsfeed", handlers.PostNews(feed))

	l.Info("Get and Post at endpoint: http://localhost:3000/newsfeed")

	return &http.Server{Addr: address, Handler: router}
}

// Start server by calling this function
func (s *serverChi) Start(ctx context.Context) {
	l := myLogger.Get()

	go func() {
		l.Info("HTTP Server running on: http://" + s.server.Addr)

		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Errorf("Error: %v", err)
		}
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting Down...")
		err := s.server.Shutdown(ctx)
		if err != nil {
			fmt.Errorf("HTTP Server ShutDown Error: %v", err)
		} else {
			fmt.Println("Shutdown successful")
		}

	}()
}

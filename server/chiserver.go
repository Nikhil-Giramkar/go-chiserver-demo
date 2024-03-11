package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nikhil-Giramkar/go-chiserver-demo/domain/newsFeed"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type serverChi struct {
	router *chi.Mux
	server *http.Server
}

func NewServer(address string, feed newsFeed.Getter) *serverChi {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	server := setUpGetHandlers(address, router, feed)
	return &serverChi{router, server}
}

func setUpGetHandlers(address string, router *chi.Mux, feed newsFeed.Getter) *http.Server {
	router.Get("/getnews", getNewsHandler(feed))

	return &http.Server{Addr: address, Handler: router}
}

func getNewsHandler(feed newsFeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.Getter()
		encoder := json.NewEncoder(w)

		encoder.SetIndent("", "	")
		encoder.Encode(items)
	}
}

func (s *serverChi) Start(ctx context.Context) {
	go func() {
		fmt.Println("HTTP Server running on: " + s.server.Addr)

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

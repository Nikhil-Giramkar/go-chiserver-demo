package server

import (
	"context"
	"encoding/json"
	"fmt"
	newsFeed "go-chiserver-demo/domain"
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
	server := setUpGetHandlers(address, router, feed)
	return &serverChi{router, server}
}

// Get Handlers added here
func setUpGetHandlers(address string, router *chi.Mux, feed *newsFeed.Repo) *http.Server {

	//http://localhost:3000/newsfeed

	router.Get("/newsfeed", getNewsHandler(feed))

	//Pass JSON in request body at bove URL
	router.Post("/newsfeed", postNewsHandler(feed))

	return &http.Server{Addr: address, Handler: router}
}

// Get All News
func getNewsHandler(feed *newsFeed.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		encoder := json.NewEncoder(w)

		encoder.SetIndent("", "	")
		encoder.Encode(items)
	}
}

// Post a news Feed
func postNewsHandler(feed *newsFeed.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		decoder := json.NewDecoder(r.Body)

		decoder.Decode(&request)

		feed.Add(newsFeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})

		w.Write([]byte("News Feed Item added successfully"))
	}
}

// Start server by calling this function
func (s *serverChi) Start(ctx context.Context) {
	go func() {
		fmt.Println("HTTP Server running on: http://" + s.server.Addr)

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

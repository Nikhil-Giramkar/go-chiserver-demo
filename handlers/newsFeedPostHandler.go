package handlers

import (
	"encoding/json"
	newsFeed "go-chiserver-demo/domain"
	myLogger "go-chiserver-demo/logger"
	"net/http"
)

// Post a news Feed
func PostNews(feed *newsFeed.Repo) http.HandlerFunc {
	l := myLogger.Get()

	return func(w http.ResponseWriter, r *http.Request) {
		l.Info("Post Request initialized")
		request := map[string]string{}
		decoder := json.NewDecoder(r.Body)

		decoder.Decode(&request)

		feed.Add(newsFeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})

		l.Info("Post success, new item added in NewsFeed")
		w.Write([]byte("News Feed Item added successfully"))
	}
}

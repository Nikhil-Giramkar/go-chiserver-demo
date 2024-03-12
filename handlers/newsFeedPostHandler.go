package handlers

import (
	"encoding/json"
	newsFeed "go-chiserver-demo/domain"
	"net/http"
)

// Post a news Feed
func PostNews(feed *newsFeed.Repo) http.HandlerFunc {
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

package handlers

import (
	"encoding/json"
	newsFeed "go-chiserver-demo/domain"
	"net/http"
)

// Get All News
func GetNews(feed *newsFeed.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		encoder := json.NewEncoder(w)

		encoder.SetIndent("", "	")
		encoder.Encode(items)
	}
}

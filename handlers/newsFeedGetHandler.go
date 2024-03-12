package handlers

import (
	"encoding/json"
	newsFeed "go-chiserver-demo/domain"
	myLogger "go-chiserver-demo/logger"
	"net/http"
)

// Get All News
func GetNews(feed *newsFeed.Repo) http.HandlerFunc {
	l := myLogger.Get()
	return func(w http.ResponseWriter, r *http.Request) {
		l.Info("Get Request initialized")
		items := feed.GetAll()
		encoder := json.NewEncoder(w)

		encoder.SetIndent("", "	")
		encoder.Encode(items)
		l.Info("News items encoded")
	}
}

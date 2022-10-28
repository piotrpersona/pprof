package driver

import (
	"encoding/json"
	"net/http"

	"github.com/google/pprof/profile"
)

type scrapeRequest struct {
	ScrapeFrom string `json:"scrapeFrom"`
}

func (ui *webInterface) scrape(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	var request scrapeRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	payload, err := ui.scraper.Scrape(ctx, request.ScrapeFrom)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	profile, err := profile.ParseData(payload)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	ui.prof = profile
	http.Redirect(writer, req, "/", http.StatusSeeOther)
	return
}

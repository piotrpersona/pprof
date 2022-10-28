package driver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/pprof/internal/scrape"
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
		http.Error(writer, fmt.Sprintf("cannot decode, err: %s", err), http.StatusInternalServerError)
		return
	}
	payload, err := ui.scraper.Scrape(ctx, request.ScrapeFrom)
	if err != nil {
		http.Error(writer, fmt.Sprintf("cannot scrape, err: %s", err), http.StatusInternalServerError)
		return
	}
	profile, err := profile.ParseData(payload)
	if err != nil {
		http.Error(writer, fmt.Sprintf("cannot parse, err: %s", err), http.StatusInternalServerError)
		return
	}

	profileDump := scrape.NewProfileDump(profile)
	_, err = ui.storage.SaveProfile(ctx, profileDump)
	if err != nil {
		http.Error(writer, fmt.Sprintf("cannot save, err: %s", err), http.StatusInternalServerError)
		return
	}

	ui.prof = profile

	http.Redirect(writer, req, "/", http.StatusSeeOther)
	return
}

type getProfileResponse struct {
	Profiles []*scrape.ProfileDump `json:"profiles,omitempty"`
}

func (ui *webInterface) getProfiles(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	dumpProfiles := ui.storage.GetProfiles(ctx)
	resp := getProfileResponse{Profiles: make([]*scrape.ProfileDump, 0)}
	for _, dprof := range dumpProfiles {
		resp.Profiles = append(resp.Profiles, dprof)
	}
	err := json.NewEncoder(writer).Encode(resp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

type getProfileRequest struct {
	ID string `json:"id"`
}

func (ui *webInterface) getProfile(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	var request getProfileRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	profileDump := ui.storage.GetProfile(ctx, request.ID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ui.prof = profileDump.Prof()

	http.Redirect(writer, req, "/", http.StatusSeeOther)
}

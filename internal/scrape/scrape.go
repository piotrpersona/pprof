package scrape

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Scraper interface {
	Scrape(ctx context.Context, url string) (payload []byte, err error)
}

type httpScraper struct{}

func NewScraper() (scraper Scraper) {
	scraper = &httpScraper{}
	return
}

func (s *httpScraper) Scrape(ctx context.Context, url string) (payload []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("cannot Get response from %s, err: %s", url, err)
		return
	}
	defer resp.Body.Close()
	payload, err = io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("cannot read response body, err: %s", err)
		return
	}
	return
}

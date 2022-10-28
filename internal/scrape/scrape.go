package scrape

import "context"

type Scraper interface {
	Scrape(ctx context.Context, url string) (err error)
}

type scraper struct{}

func NewScraper() (scraper Scraper, err error) {
	return
}

func (s *scraper) Scrape(ctx context.Context, url string) (err error) {
	
	return
}

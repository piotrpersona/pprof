package scrape

import (
	"context"
	"fmt"
	"testing"
)

func Test_scrape(t *testing.T) {
	ctx := context.Background()
	scraper := NewScraper()
	payload, err := scraper.Scrape(ctx, "http://localhost:9999/debug/pprof/heap")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(payload))
}

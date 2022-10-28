package driver

import (
	"fmt"
	"net/http"
)

func (ui *webInterface) scrape(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "scraping")
}

package rubygems

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

func Query(query string) (bool, error) {
	url := fmt.Sprintf("https://rubygems.org/api/v1/gems/%s.json", query)
	return httputil.GetCheckOK(url)
}

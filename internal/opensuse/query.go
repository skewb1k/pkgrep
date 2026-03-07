package opensuse

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

func Query(query string) (bool, error) {
	url := fmt.Sprintf("https://software.opensuse.org/package/%s", query)
	contains, err := httputil.GetBodyContains(url, "not found...")
	return !contains, err
}

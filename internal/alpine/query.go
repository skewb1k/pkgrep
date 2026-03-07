package alpine

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

func Query(query string) (bool, error) {
	url := fmt.Sprintf("https://pkgs.alpinelinux.org/packages?branch=edge&name=%s", query)
	contains, err := httputil.GetBodyContains(url, "No matching packages found...")
	return !contains, err
}

package guix

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

func Query(query string) (bool, error) {
	url := fmt.Sprintf("https://hpc.guix.info/package/%s", query)
	contains, err := httputil.GetBodyContains(url, "<title>Guix-HPC — Oops!</title>")
	return !contains, err
}

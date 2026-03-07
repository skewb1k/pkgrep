package kali

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

func Query(query string) (bool, error) {
	url := fmt.Sprintf("https://pkg.kali.org/pkg/%s", query)
	return httputil.GetCheckOK(url)
}

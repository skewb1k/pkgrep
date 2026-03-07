package pypi

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

func Query(query string) (bool, error) {
	url := fmt.Sprintf("https://pypi.org/pypi/%s/json", query)
	return httputil.GetCheckOK(url)
}

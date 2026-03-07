package chocolatey

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

func Query(query string) (bool, error) {
	url := fmt.Sprintf("https://community.chocolatey.org/packages/%s", query)
	return httputil.GetCheckOK(url)
}

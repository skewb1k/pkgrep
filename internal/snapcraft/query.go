package snapcraft

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

func Query(query string) (bool, error) {
	url := fmt.Sprintf("https://snapcraft.io/%s", query)
	return httputil.GetCheckOK(url)
}

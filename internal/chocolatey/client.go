package chocolatey

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

type Client struct{}

func (Client) Query(query string) (bool, error) {
	url := fmt.Sprintf("https://community.chocolatey.org/packages/%s", query)
	return httputil.GetCheckOK(url)
}

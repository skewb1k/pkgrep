package cratesio

import (
	"fmt"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

type Client struct{}

func (Client) Query(query string) (bool, error) {
	url := fmt.Sprintf("https://crates.io/api/v1/crates/%s", query)
	return httputil.GetCheckOK(url)
}

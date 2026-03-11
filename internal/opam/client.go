package opam

import (
	"fmt"
	"net/http"

	"github.com/skewb1k/pkgrep/internal/httputil"
)

type Client struct {
	HTTPClient *http.Client
}

func (c Client) Query(query string) (bool, error) {
	url := fmt.Sprintf("https://opam.ocaml.org/packages/%s/", query)
	ok, err := httputil.GetCheckOK(c.HTTPClient, url)
	return ok, err
}

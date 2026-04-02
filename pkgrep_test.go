package main

import (
	"encoding/hex"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/skewb1k/pkgrep/internal/alpine"
	"github.com/skewb1k/pkgrep/internal/aosc"
	"github.com/skewb1k/pkgrep/internal/archlinux"
	"github.com/skewb1k/pkgrep/internal/aur"
	"github.com/skewb1k/pkgrep/internal/chocolatey"
	"github.com/skewb1k/pkgrep/internal/clojars"
	"github.com/skewb1k/pkgrep/internal/cran"
	"github.com/skewb1k/pkgrep/internal/cratesio"
	"github.com/skewb1k/pkgrep/internal/debian"
	"github.com/skewb1k/pkgrep/internal/dub"
	"github.com/skewb1k/pkgrep/internal/fedora"
	"github.com/skewb1k/pkgrep/internal/guix"
	"github.com/skewb1k/pkgrep/internal/hackage"
	hexclient "github.com/skewb1k/pkgrep/internal/hex"
	"github.com/skewb1k/pkgrep/internal/homebrew"
	"github.com/skewb1k/pkgrep/internal/julia"
	"github.com/skewb1k/pkgrep/internal/kali"
	"github.com/skewb1k/pkgrep/internal/luarocks"
	"github.com/skewb1k/pkgrep/internal/macports"
	"github.com/skewb1k/pkgrep/internal/melpa"
	"github.com/skewb1k/pkgrep/internal/nixpkgs"
	"github.com/skewb1k/pkgrep/internal/npm"
	"github.com/skewb1k/pkgrep/internal/nuget"
	"github.com/skewb1k/pkgrep/internal/opam"
	"github.com/skewb1k/pkgrep/internal/opensuse"
	"github.com/skewb1k/pkgrep/internal/pkggodev"
	"github.com/skewb1k/pkgrep/internal/pubdev"
	"github.com/skewb1k/pkgrep/internal/pypi"
	"github.com/skewb1k/pkgrep/internal/rubygems"
	"github.com/skewb1k/pkgrep/internal/scoop"
	"github.com/skewb1k/pkgrep/internal/sisyphus"
	"github.com/skewb1k/pkgrep/internal/snapcraft"
	"github.com/skewb1k/pkgrep/internal/ubuntu"
	"github.com/skewb1k/pkgrep/internal/voidlinux"
)

type testcase struct {
	querier Querier
	pkg     string
}

var testUserAgent = "pkgrep-test/0 (https://github.com/skewb1k/pkgrep; mailto:skewb1kunix@gmail.com)"

var testHTTPClient = &http.Client{
	Transport: &UserAgentRoundTripper{
		RoundTripper: http.DefaultTransport,
		UserAgent:    testUserAgent,
	},
	Timeout: time.Minute,
}

var tests = []testcase{
	{alpine.Client{testHTTPClient}, "alpine-base"},
	{aosc.Client{testHTTPClient}, "kernel-base"},
	{archlinux.Client{testHTTPClient}, "linux"},
	{aur.Client{testHTTPClient}, "google-chrome"},
	{chocolatey.Client{testHTTPClient}, "go"},
	{clojars.Client{testHTTPClient}, "core.typed"},
	{cran.Client{testHTTPClient}, "ggplot2"},
	{cratesio.Client{testHTTPClient}, "chroma-ls"},
	{debian.Client{testHTTPClient}, "linux-base"},
	{dub.Client{testHTTPClient}, "vibe-d"},
	{fedora.Client{testHTTPClient}, "linux-firmware"},
	{guix.Client{testHTTPClient}, "go"},
	{hackage.Client{testHTTPClient}, "ghc"},
	{hexclient.Client{testHTTPClient}, "phoenix"},
	{homebrew.Client{testHTTPClient}, "go"},
	{julia.Client{testHTTPClient}, "plots"},
	{kali.Client{testHTTPClient}, "linux"},
	{luarocks.Client{testHTTPClient}, "lua-cjson"},
	{macports.Client{testHTTPClient}, "go"},
	{melpa.Client{testHTTPClient}, "magit"},
	{nixpkgs.Client{testHTTPClient}, "home-manager"},
	{npm.Client{testHTTPClient}, "npm"},
	{nuget.Client{testHTTPClient}, "Azure.Core"},
	{opam.Client{testHTTPClient}, "ocaml"},
	{opensuse.Client{testHTTPClient}, "linux-firmware"},
	{pkggodev.Client{testHTTPClient}, "http"},
	{pubdev.Client{testHTTPClient}, "http"},
	{pypi.Client{testHTTPClient}, "pip"},
	{rubygems.Client{testHTTPClient}, "rails"},
	{scoop.Client{testHTTPClient}, "go"},
	{sisyphus.Client{testHTTPClient}, "firmware-linux"},
	{snapcraft.Client{testHTTPClient}, "go"},
	{ubuntu.Client{testHTTPClient}, "linux-firmware"},
	{voidlinux.Client{testHTTPClient}, "linux-firmware"},
}

func Test(t *testing.T) {
	t.Parallel()

	b := make([]byte, 10)
	rand.Read(b)
	randomPackage := hex.EncodeToString(b)

	for _, tt := range tests {
		name := tt.querier.Name()
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			{
				found, err := tt.querier.Query(tt.pkg)
				if err != nil {
					t.Fatalf("%s: %v", name, err)
				}

				if !found {
					t.Errorf("%s: %s not found", name, tt.pkg)
				}
			}
			{
				found, err := tt.querier.Query(randomPackage)
				if err != nil {
					t.Fatalf("%s: %v", name, err)
				}

				if found {
					t.Errorf("%s: %s found", name, randomPackage)
				}
			}
		})
	}
}

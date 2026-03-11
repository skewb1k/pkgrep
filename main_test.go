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
	"github.com/skewb1k/pkgrep/internal/macports"
	"github.com/skewb1k/pkgrep/internal/melpa"
	"github.com/skewb1k/pkgrep/internal/nixpkgs"
	"github.com/skewb1k/pkgrep/internal/npm"
	"github.com/skewb1k/pkgrep/internal/nuget"
	"github.com/skewb1k/pkgrep/internal/opam"
	"github.com/skewb1k/pkgrep/internal/opensuse"
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
	name    string
	querier Querier
	pkg     string
}

var testHTTPClient = &http.Client{
	Timeout: time.Minute,
}

var tests = []testcase{
	{"Alpine", alpine.Client{testHTTPClient}, "alpine-base"},
	{"AOSC", aosc.Client{testHTTPClient}, "kernel-base"},
	{"Arch", archlinux.Client{testHTTPClient}, "linux"},
	{"AUR", aur.Client{testHTTPClient}, "google-chrome"},
	{"Chocolatey", chocolatey.Client{testHTTPClient}, "go"},
	{"CRAN", cran.Client{testHTTPClient}, "ggplot2"},
	{"crates.io", cratesio.Client{testHTTPClient}, "chroma-ls"},
	{"Debian", debian.Client{testHTTPClient}, "linux-base"},
	{"DUB", dub.Client{testHTTPClient}, "vibe-d"},
	{"Fedora", fedora.Client{testHTTPClient}, "linux-firmware"},
	{"Guix", guix.Client{testHTTPClient}, "go"},
	{"Hackage", hackage.Client{testHTTPClient}, "ghc"},
	{"Hex", hexclient.Client{testHTTPClient}, "phoenix"},
	{"Homebrew", homebrew.Client{testHTTPClient}, "go"},
	{"Julia", julia.Client{testHTTPClient}, "plots"},
	{"Kali", kali.Client{testHTTPClient}, "linux"},
	{"MacPorts", macports.Client{testHTTPClient}, "go"},
	{"MELPA", melpa.Client{testHTTPClient}, "magit"},
	{"Nixpkgs", nixpkgs.Client{testHTTPClient}, "home-manager"},
	{"NPM", npm.Client{testHTTPClient}, "npm"},
	{"NuGet", nuget.Client{testHTTPClient}, "Azure.Core"},
	{"opam", opam.Client{testHTTPClient}, "ocaml"},
	{"openSUSE", opensuse.Client{testHTTPClient}, "linux-firmware"},
	{"pub.dev", pubdev.Client{testHTTPClient}, "http"},
	{"PyPI", pypi.Client{testHTTPClient}, "pip"},
	{"RubyGems", rubygems.Client{testHTTPClient}, "rails"},
	{"Scoop", scoop.Client{testHTTPClient}, "go"},
	{"Sisyphus", sisyphus.Client{testHTTPClient}, "firmware-linux"},
	{"Snapcraft", snapcraft.Client{testHTTPClient}, "go"},
	{"Ubuntu", ubuntu.Client{testHTTPClient}, "linux-firmware"},
	{"Void", voidlinux.Client{testHTTPClient}, "linux-firmware"},
}

func Test(t *testing.T) {
	t.Parallel()

	b := make([]byte, 10)
	rand.Read(b)
	randomPackage := hex.EncodeToString(b)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			{
				found, err := tt.querier.Query(tt.pkg)
				if err != nil {
					t.Fatalf("%s: %v", tt.name, err)
				}

				if !found {
					t.Errorf("%s: %s not found", tt.name, tt.pkg)
				}
			}
			{
				found, err := tt.querier.Query(randomPackage)
				if err != nil {
					t.Fatalf("%s: %v", tt.name, err)
				}

				if found {
					t.Errorf("%s: %s found", tt.name, randomPackage)
				}
			}
		})
	}
}

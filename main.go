package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	"github.com/skewb1k/pkgrep/internal/alpine"
	"github.com/skewb1k/pkgrep/internal/archlinux"
	"github.com/skewb1k/pkgrep/internal/aur"
	"github.com/skewb1k/pkgrep/internal/chocolatey"
	"github.com/skewb1k/pkgrep/internal/cratesio"
	"github.com/skewb1k/pkgrep/internal/debian"
	"github.com/skewb1k/pkgrep/internal/fedora"
	"github.com/skewb1k/pkgrep/internal/guix"
	"github.com/skewb1k/pkgrep/internal/homebrew"
	"github.com/skewb1k/pkgrep/internal/kali"
	"github.com/skewb1k/pkgrep/internal/macports"
	"github.com/skewb1k/pkgrep/internal/nixpkgs"
	"github.com/skewb1k/pkgrep/internal/npm"
	"github.com/skewb1k/pkgrep/internal/opensuse"
	"github.com/skewb1k/pkgrep/internal/pypi"
	"github.com/skewb1k/pkgrep/internal/rubygems"
	"github.com/skewb1k/pkgrep/internal/sisyphus"
	"github.com/skewb1k/pkgrep/internal/snapcraft"
	"github.com/skewb1k/pkgrep/internal/ubuntu"
	"github.com/skewb1k/pkgrep/internal/voidlinux"
)

type Querier interface {
	// Query accepts a search query string and returns a
	// boolean indicating whether the package was found, or an error.
	// The query string should be URL-safe.
	Query(query string) (bool, error)
}

type Repository struct {
	Name    string
	Querier Querier
}

var repos = []Repository{
	{"Alpine", alpine.Client{}},
	{"Arch", archlinux.Client{}},
	{"AUR", aur.Client{}},
	{"Chocolatey", chocolatey.Client{}},
	{"crates.io", cratesio.Client{}},
	{"Debian", debian.Client{}},
	{"Fedora", fedora.Client{}},
	{"Guix", guix.Client{}},
	{"Homebrew", homebrew.Client{}},
	{"Kali", kali.Client{}},
	{"MacPorts", macports.Client{}},
	{"Nixpkgs", nixpkgs.Client{}},
	{"NPM", npm.Client{}},
	{"openSUSE", opensuse.Client{}},
	{"PyPI", pypi.Client{}},
	{"RubyGems", rubygems.Client{}},
	{"Sisyphus", sisyphus.Client{}},
	{"Snapcraft", snapcraft.Client{}},
	{"Ubuntu", ubuntu.Client{}},
	{"Void", voidlinux.Client{}},
}

type Result struct {
	Name  string
	Found bool
}

func main() {
	log.SetPrefix("pkgrep: ")
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s QUERY\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	query := flag.Arg(0)
	query = url.QueryEscape(query)

	var wg sync.WaitGroup
	results := make(chan Result)

	for _, repo := range repos {
		wg.Add(1)
		go func(r Repository) {
			defer wg.Done()

			found, err := r.Querier.Query(query)
			if err != nil {
				log.Println(err)
				return
			}

			results <- Result{
				Name:  r.Name,
				Found: found,
			}
		}(repo)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	foundAny := false
	for result := range results {
		if result.Found {
			fmt.Printf("*")
			foundAny = true
		} else {
			fmt.Print("-")
		}
		fmt.Printf(" %s\n", result.Name)
	}

	if !foundAny {
		os.Exit(1)
	}
}

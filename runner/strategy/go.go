package strategy

import (
	"strings"

	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/zachlatta/try/util"
)

type Go struct {
	dir string
	uri string
	env map[string]string
}

func (g *Go) Init(dir, uri string) {
	g.dir = dir
	g.uri = uri
	g.env = make(map[string]string)
}

func (g Go) ShouldUse() (bool, error) {
	files, err := util.RecursiveList(g.dir)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if strings.HasSuffix(file, ".go") {
			return true, nil
		}
	}

	return false, nil
}

func (g Go) uriParts() ([]string, error) {
	parsed, err := url.Parse(g.uri)
	if err != nil {
		return nil, err
	}

	pathParts := strings.Split(parsed.Path, "/")

	return append([]string{parsed.Host}, pathParts...), nil
}

func (g *Go) setupGopath() error {
	const root = ".gopath"

	uriParts, err := g.uriParts()
	if err != nil {
		return nil
	}

	allButLast := uriParts[:len(uriParts)-1]
	last := uriParts[len(uriParts)-1]

	toMakeSeparated := append([]string{g.dir, root, "src"}, allButLast...)
	toMake := filepath.Join(toMakeSeparated...)
	if err := os.MkdirAll(toMake, 0777); err != nil {
		return err
	}

	if err := os.Symlink(g.dir, filepath.Join(toMake, last)); err != nil {
		return err
	}

	g.env["GOPATH"] = filepath.Join(g.dir, root)

	return nil
}

func (g *Go) Setup() (map[string]string, error) {
	fmt.Println("Setting up GOPATH...")
	if err := g.setupGopath(); err != nil {
		return nil, err
	}

	// TODO: Install dependencies here

	return g.env, nil
}

func init() {
	strategies = append(strategies, &Go{})
}

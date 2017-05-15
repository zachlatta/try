package strategy

import (
	"fmt"
	"path/filepath"

	"github.com/zachlatta/try/util"
)

type Node struct {
	dir string
	uri string
	env map[string]string
}

func (n *Node) Init(dir, uri string) {
	n.dir = dir
	n.uri = uri
	n.env = make(map[string]string)
}

// findPackageJsons finds all of the package.json files in the current and child
// directories and returns a slice of paths of directories that have them.
func (n Node) findPackageJsons() ([]string, error) {
	files, err := util.RecursiveList(n.dir)
	if err != nil {
		return nil, err
	}

	dirs := []string{}

	for _, file := range files {
		if filepath.Base(file) == "package.json" {
			dirs = append(dirs, filepath.Dir(file))
		}
	}

	return dirs, nil
}

func (n Node) ShouldUse() (bool, error) {
	packageJsonDirs, err := n.findPackageJsons()
	if err != nil {
		return false, err
	}

	if len(packageJsonDirs) > 0 {
		return true, nil
	}

	return false, nil
}

func (n Node) Setup() (map[string]string, error) {
	fmt.Println("Installing Node dependencies...")

	packageJsonDirs, err := n.findPackageJsons()
	if err != nil {
		return nil, err
	}

	for _, packageJsonDir := range packageJsonDirs {
		fmt.Println(packageJsonDir)
		_, err := util.Run(packageJsonDir, n.env, "npm", "install")
		if err != nil {
			return nil, err
		}
	}

	return n.env, nil
}

func init() {
	strategies = append(strategies, &Node{})
}

package runner

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zachlatta/try/runner/shell"
	"github.com/zachlatta/try/runner/strategy"
	"github.com/zachlatta/try/runner/vcs"
)

func tempDir(prefix string) (name string, err error) {
	return ioutil.TempDir("", prefix)
}

func Run(repoUrl string) error {
	dir, err := tempDir("try")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	fmt.Println("Cloning repo...")
	if err := vcs.Clone(dir, repoUrl); err != nil {
		return err
	}

	strategy.InitAll(dir, repoUrl)

	env, err := strategy.SetupAll()
	if err != nil {
		return err
	}

	if err := shell.Start(dir, env); err != nil {
		return err
	}

	return nil
}

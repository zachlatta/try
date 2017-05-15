package shell

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/zachlatta/try/util"
)

const defaultShell = "bash"

func Start(dir string, env map[string]string) error {
	shellPath, err := exec.LookPath(defaultShell)
	if err != nil {
		return errors.New(defaultShell + " is not installed")
	}

	cmd := exec.Cmd{
		Path:   shellPath,
		Dir:    dir,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	if len(env) > 0 {
		fmt.Println("Starting", defaultShell+"...")
		fmt.Println()
		fmt.Println("> Please export the following environment variables:")
		fmt.Println()

		for _, env := range util.ConstructEnv(env, false) {
			fmt.Println("export", env)
		}

		fmt.Println()
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

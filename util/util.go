package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RecursiveList(dir string) (files []string, err error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		absolutePath := filepath.Join(dir, fileInfo.Name())

		if fileInfo.IsDir() {
			childFiles, err := RecursiveList(absolutePath)
			if err != nil {
				return nil, err
			}

			files = append(files, childFiles...)
		} else {
			files = append(files, absolutePath)
		}
	}

	return files, nil
}

func ConstructEnv(env map[string]string, insertCurrentEnv bool) (newEnv []string) {
	mutEnv := make(map[string]string)

	if env != nil {
		for key, val := range env {
			mutEnv[key] = val
		}
	}

	if insertCurrentEnv {
		for _, rawEnv := range os.Environ() {
			split := strings.SplitN(rawEnv, "=", 2)

			name := split[0]
			val := split[1]

			if _, exists := mutEnv[name]; !exists {
				mutEnv[name] = val
			}
		}
	}

	for key, val := range env {
		newEnv = append(newEnv, key+"="+val)
	}

	return newEnv
}

func Run(dir string, env map[string]string, cmd string, args ...string) (output string, err error) {
	cmdPath, err := exec.LookPath(cmd)
	if err != nil {
		return "", err
	}

	outputBuf := &bytes.Buffer{}

	c := exec.Cmd{
		Path:   cmdPath,
		Args:   append([]string{""}, args...),
		Dir:    dir,
		Stdout: outputBuf,
		Stderr: outputBuf,
		Env:    ConstructEnv(env, true),
	}

	if err := c.Run(); err != nil {
		return "", errors.New(outputBuf.String())
	}

	return outputBuf.String(), nil
}

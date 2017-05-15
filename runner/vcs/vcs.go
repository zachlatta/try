package vcs

import (
	"github.com/zachlatta/try/util"
)

func Clone(dir, uri string) error {
	// TODO: Support other VCS systems.
	return gitClone(dir, uri)
}

func gitClone(dir, uri string) error {
	_, err := util.Run(dir, nil, "git", "clone", uri, ".")
	if err != nil {
		return err
	}

	return nil
}

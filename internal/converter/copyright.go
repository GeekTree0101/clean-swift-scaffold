package converter

import (
	"os/exec"
	"os/user"
	"strings"
)

type Copyright interface {
	Get() (string, error)
}

type CopyrightImpl struct {
}

func (c *CopyrightImpl) Get() (string, error) {

	gitUsername, err := c.getGitUsername()

	if err == nil {
		return gitUsername, nil
	}

	user, err := user.Current()

	if err == nil {
		return user.Username, nil
	}

	return "", err
}

func (c *CopyrightImpl) getGitUsername() (string, error) {

	gitCmd := exec.Command("git", "config", "--global", "user.name")
	nameBytes, err := gitCmd.Output()

	if err == nil && len(nameBytes) > 0 {
		return strings.Trim(string(nameBytes), "\n"), nil
	}

	return "", err
}

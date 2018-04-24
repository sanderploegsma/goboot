package command

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

// Runner runs a command on the host OS and returns the command output.
type Runner interface {
	Run(command string, args []string) (string, error)
}

// NewRunner creates a new runner that uses `os/exec` to run commands
func NewRunner() Runner {
	return &runner{}
}

type runner struct{}

func (runner *runner) Run(command string, args []string) (string, error) {
	logrus.WithFields(logrus.Fields{
		"args":    args,
		"command": command,
	}).Debug("Invoking command")

	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output[:len(output)]), nil
}

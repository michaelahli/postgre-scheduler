package cmd

import (
	"os/exec"
)

type terminal struct {
	Cmd string
}

type Terminal interface {
	ExecuteBash(...string) (string, error)
}

func NewTerminal(command string) Terminal {
	return &terminal{
		Cmd: command,
	}
}

func (sh *terminal) ExecuteBash(args ...string) (output string, err error) {
	cmd := exec.Command(sh.Cmd, args...)

	stdout, err := cmd.Output()
	if err != nil {
		return output, err
	}

	output = string(stdout)

	return output, err
}

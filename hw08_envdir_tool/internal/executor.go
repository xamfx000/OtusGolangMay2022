package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	command.Env = prepareEnv(env)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	err := command.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
	}
	return 0
}

func prepareEnv(env Environment) []string {
	for name, envVar := range env {
		if envVar.NeedRemove {
			err := os.Unsetenv(name)
			if err != nil {
				fmt.Printf("%v", err)
			}
			continue
		}
		err := os.Setenv(name, envVar.Value)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	return os.Environ()
}

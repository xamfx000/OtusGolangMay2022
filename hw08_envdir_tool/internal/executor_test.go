package internal

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRunCmd(t *testing.T) {
	env := Environment{
		"BAR": {
			"bar",
			false,
		},
		"FOO": {
			"FOO",
			false,
		},
		"HOME": {
			"/usr/local/dump",
			false,
		},
		"PATH": {
			"/bin/",
			false,
		},
	}
	pwd, _ := os.LookupEnv("PWD")
	dirName := fmt.Sprintf("%s/bar", pwd)
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dirName)
	exitCode := RunCmd([]string{"mkdir", dirName}, env)
	_, err := os.ReadDir(dirName)
	require.NoError(t, err)
	require.Equal(t, 0, exitCode)
}

func TestRunCmd_NonZeroExitCodeWithoutErr(t *testing.T) {
	env := Environment{
		"BAR": {
			"bar",
			false,
		},
		"FOO": {
			"FOO",
			false,
		},
		"HOME": {
			"/usr/local/dump",
			false,
		},
		"PATH": {
			"/bin/",
			false,
		},
	}
	exitCode := RunCmd([]string{"exit 5"}, env)
	require.Equal(t, 0, exitCode)
}

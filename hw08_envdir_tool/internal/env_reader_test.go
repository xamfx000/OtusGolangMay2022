package internal

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	env, _ := ReadDir("../testdata/env")
	expected := Environment{
		"BAR": {
			"bar",
			false,
		},
		"EMPTY": {
			"",
			false,
		},
		"FOO": {
			"   foo\nwith new line",
			false,
		},
		"HELLO": {
			"\"hello\"",
			false,
		},
		"UNSET": {
			"",
			true,
		},
	}
	require.Equal(t, expected, env)
}

func TestReadDir_ErrorForNonExistentDirectory(t *testing.T) {
	env, err := ReadDir("/Users/test")
	require.Error(t, err)
	require.Nil(t, env)
}

func TestReadDir_NoEnvInEmptyDir(t *testing.T) {
	pwd, _ := os.Getwd()
	dirName := fmt.Sprintf("%s/temp", pwd)
	_ = os.Mkdir(dirName, 777)
	defer func() {
		_ = os.Remove(dirName)
	}()
	env, err := ReadDir(dirName)
	require.NoError(t, err)
	require.Equal(t, Environment{}, env)
}

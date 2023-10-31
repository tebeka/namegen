package main

import (
	"os/exec"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_humanize(t *testing.T) {
	testCases := []struct {
		name  string
		human string
	}{
		{"wednesday_addams", "Wednesday Addams"},
		{"", ""},
	}

	for _, tc := range testCases {
		name := tc.name
		if name == "" {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			out := humanize(tc.name)
			require.Equal(t, tc.human, out)
		})
	}
}

func build(t *testing.T) string {
	exe := path.Join(t.TempDir(), "namegen")
	if runtime.GOOS == "windows" {
		exe += ".exe"
	}
	err := exec.Command("go", "build", "-o", exe).Run()
	require.NoError(t, err)

	return exe
}

func TestExe(t *testing.T) {
	exe := build(t)

	cmd := exec.Command(exe)
	data, err := cmd.CombinedOutput()
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestExeHelp(t *testing.T) {
	exe := build(t)

	cmd := exec.Command(exe, "-h")
	data, err := cmd.CombinedOutput()
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

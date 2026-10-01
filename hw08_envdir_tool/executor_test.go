package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		name     string
		cmd      []string
		newEnv   Environment
		oldEnv   map[string]string
		exitCode int
	}{
		{
			name:     "empty",
			cmd:      []string{},
			newEnv:   Environment{},
			exitCode: 1,
		},
		{
			name:     "simple success",
			cmd:      []string{"/bin/sh", "-c", "0"},
			newEnv:   Environment{},
			exitCode: 0,
		},
		{
			name:     "simple error",
			cmd:      []string{"/bin/sh", "-c", "123"},
			newEnv:   Environment{},
			exitCode: 123,
		},
		{
			name: "simple add env",
			cmd:  []string{"/bin/sh", "-c", `test "$TEMP" = "bar"`},
			newEnv: Environment{
				"TEMP": {
					Value: "bar",
				},
			},
			exitCode: 0,
		},
		{
			name: "simple remove env",
			cmd:  []string{"/bin/sh", "-c", `test -z "$TEMP" = "bar"`},
			newEnv: Environment{
				"TEMP": {
					Value: "bar",
				},
			},
			oldEnv: map[string]string{
				"TEMP": "old",
			},
			exitCode: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.oldEnv != nil {
				for k, v := range test.oldEnv {
					t.Setenv(k, v)
				}
			}

			code := RunCmd(test.cmd, test.newEnv)

			require.Equal(t, test.exitCode, code)
		})
	}
}

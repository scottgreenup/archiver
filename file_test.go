package main

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/require"
)


func testAbsPath(t *testing.T) {
	SafeAbsPath := func(t *testing.T, path string) string {
		s, err := AbsPath(path)
		require.NoError(t, err)
		return s
	}

	require.True(t, strings.HasSuffix(SafeAbsPath(t, "myAwesomeFile"), "/myAwesomeFile"))
	require.True(t, strings.HasSuffix(SafeAbsPath(t, "~/myAwesomeFile"), "/myAwesomeFile"))
	require.False(t, strings.Contains(SafeAbsPath(t, "~/myAwesomeFile"), "~"))
}

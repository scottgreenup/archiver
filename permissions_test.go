package main

import (
	"testing"
	"os"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	testcases := []struct {
		FileName string
		Permission
		Class
	} {
		{"./testdata/permissions/root_root_r--------", ReadPermission, OwnerClass},
		{"./testdata/permissions/root_root_---r-----", ReadPermission, GroupClass},
		{"./testdata/permissions/root_root_------r--", ReadPermission, OtherClass},
		{"./testdata/permissions/user_user_r--------", ReadPermission, OwnerClass},
		{"./testdata/permissions/user_user_---r-----", ReadPermission, GroupClass},
		{"./testdata/permissions/user_user_------r--", ReadPermission, OtherClass},
	}

	for _, tc := range testcases {
		fi, err := os.Stat(tc.FileName)
		require.NoError(t, err)

		perm := FilePerm(fi.Mode())
		require.True(t, perm.Check(tc.Permission, tc.Class))
	}
}

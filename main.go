package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/pkg/errors"
)

type Archiver struct {
	FromPath string
	ToPath string
}

func ValidateTo(path string) error {
	return nil
}

func IsReadable(fi os.FileInfo) bool {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	filePerm := FilePerm(fi.Mode().Perm())

	// owner
	fileUser := strconv.Itoa(int(fi.Sys().(*syscall.Stat_t).Uid))
	if usr.Uid == fileUser && filePerm.Check(ReadPermission, OwnerClass) {
		return true
	}

	// group
	fileGroup := strconv.Itoa(int(fi.Sys().(*syscall.Stat_t).Gid))
	gids, err := usr.GroupIds()
	if err != nil {
		panic(err)
	}
	for _, gid := range gids {
		if gid == fileGroup {
			if filePerm.Check(ReadPermission, GroupClass) {
				return true
			} else {
				break
			}
		}
	}

	// others
	return filePerm.Check(ReadPermission, OtherClass)
}

func IsWriteable(path string) error {
	return nil
}

func ValidateFrom(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return errors.WithStack(err)
	}

	if !fi.Mode().IsDir() && !fi.Mode().IsRegular() {
		return errors.WithStack(errors.Errorf("%s is not a directory or regular file", path))
	}

	// TODO walk the file tree
	if ! IsReadable(fi) {
		return errors.Errorf("Unable to read %q", path)
	}

	return nil
}

func AbsPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			return "", errors.WithStack(err)
		}
		return filepath.Join(usr.HomeDir, path[2:]), nil
	}

	return filepath.Abs(path)
}

func NewArchiver(target ConfigTarget) (*Archiver, error) {
	from, err := AbsPath(target.From)
	if err != nil {
		return nil, err
	}
	if err := ValidateFrom(from); err != nil {
		return nil, err
	}

	to, err := AbsPath(target.To)
	if err != nil {
		return nil, err
	}
	if err := ValidateTo(to); err != nil {
		return nil, err
	}

	return &Archiver{
		FromPath: from,
		ToPath: to,
	}, nil

}

func main() {
	configFile := flag.String("config", "Archiverfile", "Path to config file")
	flag.Parse()

	config, err := NewConfig(*configFile)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", config)
	}

	a, err := NewArchiver(config.Targets["GPG"])
	fmt.Printf("%+v %+v\n", a, err)
}

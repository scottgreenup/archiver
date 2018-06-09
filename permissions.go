package main

import (
	"os"
)

type FilePerm os.FileMode
type Permission uint32
type Class uint32

const (
	ReadPermission    Permission = 4
	WritePermission   Permission = 2
	ExecutePermission Permission = 1

	OwnerClass Class = 6
	GroupClass Class = 3
	OtherClass Class = 0

)

func (p FilePerm) Check(permission Permission, class Class) bool {
	return (uint32(p) & (uint32(permission) << uint32(class))) != 0
}

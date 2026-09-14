// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package unix

import "syscall"

const (
	AT_EACCESS          = 0x200
	AT_FDCWD            = -0x64
	AT_REMOVEDIR        = 0x200
	AT_SYMLINK_NOFOLLOW = 0x100
	AT_SYMLINK_FOLLOW   = 0x400
	AT_EMPTY_PATH       = 0x1000
	UTIME_OMIT          = -0x2
)

var faccessat = syscall.Faccessat

func Unlinkat(dirfd int, path string, flags int) error {
	return syscall.Unlinkat(dirfd, path, flags)
}

func Openat(dirfd int, path string, flags int, perm uint32) (int, error) {
	return syscall.Openat(dirfd, path, flags, perm)
}

func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error {
	return syscall.Fstatat(dirfd, path, stat, flags)
}

func Readlinkat(dirfd int, path string, buf []byte) (int, error) {
	return syscall.Readlinkat(dirfd, path, buf)
}

func Mkdirat(dirfd int, path string, mode uint32) error {
	return syscall.Mkdirat(dirfd, path, mode)
}

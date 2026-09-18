// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package syscall

import "unsafe"

// FcntlFlock performs a fcntl call for the [F_GETLK], [F_SETLK] or
// [F_SETLKW] command.
//
// GNU/Hurd has no raw fcntl entry point usable from Go, so go through
// glibc's fcntl(2).  Flock_t has the layout of struct flock64 (64-bit
// offsets), so the command must select the 64-bit variant: on 32-bit Hurd
// the plain F_SETLK/F_GETLK commands describe the 32-bit struct flock.
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error {
	switch cmd {
	case F_GETLK:
		cmd = F_GETLK64
	case F_SETLK:
		cmd = F_SETLK64
	case F_SETLKW:
		cmd = F_SETLKW64
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fcntl)), 3, fd, uintptr(cmd), uintptr(unsafe.Pointer(lk)), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

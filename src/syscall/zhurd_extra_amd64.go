// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd && amd64

package syscall

// Hurd constants that are not present in the Linux/amd64 tables.
const (
	SO_REUSEPORT   = _SO_REUSEPORT
	SO_USELOOPBACK = _SO_USELOOPBACK
)

// On Hurd/amd64 rlim_t is 64-bit, so RLIM_INFINITY is RLIM64_INFINITY; the
// glibc headers do not define an un-suffixed name.
const _RLIM_INFINITY = _RLIM64_INFINITY

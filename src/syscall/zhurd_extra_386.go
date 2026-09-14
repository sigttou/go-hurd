// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd && 386

package syscall

// Hurd constants that are not present in the Linux/386 tables.
const (
	SO_REUSEPORT   = _SO_REUSEPORT
	SO_USELOOPBACK = _SO_USELOOPBACK
)

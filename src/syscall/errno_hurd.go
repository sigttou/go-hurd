// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package syscall

// errnoBase is the offset that GNU/Hurd adds to every errno code
// (glibc's <bits/errno.h> defines e.g. ENOENT as 0x40000002).  The
// generated errors table is indexed by the code without this offset,
// so Errno.Error must subtract it before looking up a message.
const errnoBase = 0x40000000

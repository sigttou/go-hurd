// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix && !hurd

package syscall

// errnoBase is zero everywhere except GNU/Hurd, whose errno codes are
// offset (see errno_hurd.go).  It keeps Errno.Error uniform.
const errnoBase = 0

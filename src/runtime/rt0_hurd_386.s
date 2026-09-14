// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT _rt0_386_hurd(SB),NOSPLIT,$0
	JMP	_rt0_386(SB)

TEXT _rt0_386_hurd_lib(SB),NOSPLIT,$0
	JMP	_rt0_386_lib(SB)

TEXT main(SB),NOSPLIT,$0
	// Remove the return address from the stack.
	// rt0_go doesn't expect it to be there.
	ADDL	$4, SP
	JMP	runtime·rt0_go(SB)

// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd && amd64

package runtime

// sigset is a glibc __sigset_t on Hurd/amd64: 8 bytes (one 64-bit word).
type sigset uint64

// siginfo_t on Hurd/amd64. sizeof(siginfo_t) == 56; the compiler inserts the
// padding needed to reach the 8-byte-aligned si_addr/si_band fields.
type siginfo struct {
	si_signo  int32
	si_errno  int32
	si_code   int32
	si_pid    int32
	si_uid    uint32
	si_addr   uintptr
	si_status int32
	si_band   int64
	__pad     [8]byte
}

type timespec struct {
	tv_sec  int64
	tv_nsec int64
}

//go:nosplit
func (ts *timespec) setNsec(ns int64) {
	ts.tv_sec = ns / 1e9
	ts.tv_nsec = ns % 1e9
}

//go:nosplit
func (ts *timespec) set(sec, nsec int64) {
	ts.tv_sec = sec
	ts.tv_nsec = nsec
}

type timeval struct {
	tv_sec  int64
	tv_usec int64
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = int64(x)
}

// User-level context layout (glibc 2.43, Hurd/amd64):
//
//	ucontext_t { uc_flags(0); uc_link(8); uc_stack(16, 24 bytes);
//	             uc_mcontext(40, 256 bytes); uc_sigmask(296, 8 bytes);
//	             __fpregs_mem(304, 512 bytes); __ssp(816, 32 bytes) }
//	mcontext_t { gregset_t gregs; fpregset_t fpregs; reserved1[8] }
//	gregset_t  = long long[23] { R8..R15, RDI, RSI, RBP, RSP, RBX, RDX,
//	                             RCX, RAX, RIP, CS, RFL, ERR, TRAPNO,
//	                             OLDMASK, CR2 }
type mcontext struct {
	gregs       [23]int64
	fpregs      uintptr
	__reserved1 [8]uint64
}

type ucontext struct {
	uc_flags     uint64
	uc_link      uintptr
	uc_stack     stackt
	uc_mcontext  mcontext
	uc_sigmask   sigset
	__fpregs_mem [512]byte
	__ssp        [4]uint64
}

const (
	_REG_R8      = 0
	_REG_R9      = 1
	_REG_R10     = 2
	_REG_R11     = 3
	_REG_R12     = 4
	_REG_R13     = 5
	_REG_R14     = 6
	_REG_R15     = 7
	_REG_RDI     = 8
	_REG_RSI     = 9
	_REG_RBP     = 10
	_REG_RSP     = 11
	_REG_RBX     = 12
	_REG_RDX     = 13
	_REG_RCX     = 14
	_REG_RAX     = 15
	_REG_RIP     = 16
	_REG_CS      = 17
	_REG_RFL     = 18
	_REG_ERR     = 19
	_REG_TRAPNO  = 20
	_REG_OLDMASK = 21
	_REG_CR2     = 22
)

type pthread uint64
type pthread_attr [48]byte
type semt [24]byte

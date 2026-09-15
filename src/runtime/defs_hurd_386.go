// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd && 386

package runtime

// sigset is a glibc __sigset_t on Hurd/i386: a single 32-bit word.
type sigset uint32

type siginfo struct {
	si_signo  int32
	si_errno  int32
	si_code   int32
	si_pid    int32
	si_uid    uint32
	si_addr   uintptr
	si_status int32
	si_band   int32
	__pad     int32 // sizeof(siginfo_t) == 36 on Hurd/i386
}

type timespec struct {
	tv_sec  int32
	tv_nsec int32
}

//go:nosplit
func (ts *timespec) setNsec(ns int64) {
	ts.tv_sec = int32(ns / 1e9)
	ts.tv_nsec = int32(ns % 1e9)
}

//go:nosplit
func (ts *timespec) set(sec, nsec int64) {
	ts.tv_sec = int32(sec)
	ts.tv_nsec = int32(nsec)
}

type timeval struct {
	tv_sec  int32
	tv_usec int32
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = x
}

// User-level context layout (glibc 2.43, Hurd/i386):
//
//	ucontext_t { uc_flags(0); uc_link(4); uc_sigmask(8); uc_stack(12);
//	             uc_mcontext(24, 456 bytes); reserved[5] }
//	mcontext_t { gregset_t gregs; fpregset_t fpregs }
//	gregset_t  = int[19] { GS,FS,ES,DS,EDI,ESI,EBP,ESP,EBX,EDX,ECX,EAX,
//	                        TRAPNO,ERR,EIP,CS,EFL,UESP,SS }
type mcontext struct {
	gregs  [19]int32
	fpregs [95]int32
}

type ucontext struct {
	uc_flags    uint32
	uc_link     uintptr
	uc_sigmask  sigset
	uc_stack    stackt
	uc_mcontext mcontext
	__reserved  [5]uint32
}

const (
	_REG_GS     = 0
	_REG_FS     = 1
	_REG_ES     = 2
	_REG_DS     = 3
	_REG_EDI    = 4
	_REG_ESI    = 5
	_REG_EBP    = 6
	_REG_ESP    = 7
	_REG_EBX    = 8
	_REG_EDX    = 9
	_REG_ECX    = 10
	_REG_EAX    = 11
	_REG_TRAPNO = 12
	_REG_ERR    = 13
	_REG_EIP    = 14
	_REG_CS     = 15
	_REG_EFL    = 16
	_REG_UESP   = 17
	_REG_SS     = 18
)

type pthread uint32
type pthread_attr [32]byte
type semt [20]byte

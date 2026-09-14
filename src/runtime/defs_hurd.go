// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package runtime

// Hurd/i386 definitions. Values were obtained by compiling a probe against
// the Debian GNU/Hurd glibc 2.43 headers on the target (see
// PHASE0-FINDINGS.md). Note that Hurd errno values are 0x40000000|n and the
// signal numbers differ from Linux/BSD.

const (
	_EPERM     = 0x40000001
	_ENOENT    = 0x40000002
	_EINTR     = 0x40000004
	_EIO       = 0x40000005
	_EBADF     = 0x40000009
	_ENOMEM    = 0x4000000c
	_EACCES    = 0x4000000d
	_EFAULT    = 0x4000000e
	_EBUSY     = 0x40000010
	_EINVAL    = 0x40000016
	_EAGAIN    = 0x40000023
	_ETIMEDOUT = 0x4000003c
	_ENOSYS    = 0x4000004e

	_PROT_NONE  = 0x0
	_PROT_READ  = 0x4
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x1

	_MAP_ANON    = 0x2
	_MAP_PRIVATE = 0x0
	_MAP_FIXED   = 0x100

	_MADV_DONTNEED = 0x4

	_SIGHUP    = 0x1
	_SIGINT    = 0x2
	_SIGQUIT   = 0x3
	_SIGILL    = 0x4
	_SIGTRAP   = 0x5
	_SIGABRT   = 0x6
	_SIGFPE    = 0x8
	_SIGKILL   = 0x9
	_SIGBUS    = 0xa
	_SIGSEGV   = 0xb
	_SIGSYS    = 0xc
	_SIGPIPE   = 0xd
	_SIGALRM   = 0xe
	_SIGTERM   = 0xf
	_SIGURG    = 0x10
	_SIGSTOP   = 0x11
	_SIGTSTP   = 0x12
	_SIGCONT   = 0x13
	_SIGCHLD   = 0x14
	_SIGTTIN   = 0x15
	_SIGTTOU   = 0x16
	_SIGIO     = 0x17
	_SIGXCPU   = 0x18
	_SIGXFSZ   = 0x19
	_SIGVTALRM = 0x1a
	_SIGPROF   = 0x1b
	_SIGWINCH  = 0x1c
	_SIGUSR1   = 0x1e
	_SIGUSR2   = 0x1f

	_FPE_INTDIV = 0x1
	_FPE_INTOVF = 0x2
	_FPE_FLTDIV = 0x3
	_FPE_FLTOVF = 0x4
	_FPE_FLTUND = 0x5
	_FPE_FLTRES = 0x6
	_FPE_FLTINV = 0x7
	_FPE_FLTSUB = 0x8

	_BUS_ADRALN = 0x1
	_BUS_ADRERR = 0x2
	_BUS_OBJERR = 0x3

	_SEGV_MAPERR = 0x1
	_SEGV_ACCERR = 0x2

	_ITIMER_REAL    = 0x0
	_ITIMER_VIRTUAL = 0x1
	_ITIMER_PROF    = 0x2

	_O_RDONLY   = 0x0
	_O_WRONLY   = 0x2
	_O_NONBLOCK = 0x8
	_O_CREAT    = 0x10
	_O_TRUNC    = 0x10000

	_SS_DISABLE = 0x4
	_SI_USER    = 0x0

	_SIG_BLOCK   = 0x1
	_SIG_UNBLOCK = 0x2
	_SIG_SETMASK = 0x3

	_SA_SIGINFO = 0x40
	_SA_RESTART = 0x2
	_SA_ONSTACK = 0x1

	_PTHREAD_CREATE_DETACHED = 0x1

	__SC_PAGE_SIZE        = 0x1e
	__SC_NPROCESSORS_ONLN = 0x54

	_F_SETFL = 0x4
	_F_GETFD = 0x1
	_F_GETFL = 0x3

	_CLOCK_REALTIME  = 0
	_CLOCK_MONOTONIC = 1

	_NSIG = 33
)

// sigset is a glibc __sigset_t on Hurd/i386: a single 32-bit word.
type sigset uint32

var sigset_all = ^sigset(0)

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

type timeval struct {
	tv_sec  int32
	tv_usec int32
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = x
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

// stack_t on Hurd/i386 is { void *ss_sp; size_t ss_size; int ss_flags; }.
type stackt struct {
	ss_sp    uintptr
	ss_size  uintptr
	ss_flags int32
}

// struct sigaction on Hurd/i386 has no sa_restorer.
type sigactiont struct {
	sa_handler uintptr
	sa_mask    sigset
	sa_flags   int32
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

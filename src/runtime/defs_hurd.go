// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package runtime

// Hurd definitions shared by all architectures. Values were obtained by
// compiling a probe against the Debian GNU/Hurd glibc headers on the target
// (see PHASE0-FINDINGS.md and amd64-probe/). Note that Hurd errno values are
// 0x40000000|n and the signal numbers differ from Linux/BSD. The signal
// numbers come from Mach and are the same on i386 and amd64.

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

// sigset_all is an all-ones mask of the architecture-specific sigset type.
var sigset_all = ^sigset(0)

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

// stack_t on Hurd is { void *ss_sp; size_t ss_size; int ss_flags; }.
type stackt struct {
	ss_sp    uintptr
	ss_size  uintptr
	ss_flags int32
}

// struct sigaction on Hurd has no sa_restorer.
type sigactiont struct {
	sa_handler uintptr
	sa_mask    sigset
	sa_flags   int32
}

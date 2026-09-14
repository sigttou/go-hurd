// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package runtime

// Hurd/i386 signal numbers (see PHASE0-FINDINGS.md). The array is indexed by
// the OS signal number and must have _NSIG entries.
var sigtable = [...]sigTabT{
	0:           {0, "SIGNONE: no trap"},
	_SIGHUP:     {_SigNotify + _SigKill, "SIGHUP: terminal line hangup"},
	_SIGINT:     {_SigNotify + _SigKill, "SIGINT: interrupt"},
	_SIGQUIT:    {_SigNotify + _SigThrow, "SIGQUIT: quit"},
	_SIGILL:     {_SigThrow + _SigUnblock, "SIGILL: illegal instruction"},
	_SIGTRAP:    {_SigThrow + _SigUnblock, "SIGTRAP: trace trap"},
	_SIGABRT:    {_SigNotify + _SigThrow, "SIGABRT: abort"},
	_SIGFPE:     {_SigPanic + _SigUnblock, "SIGFPE: floating-point exception"},
	_SIGKILL:    {0, "SIGKILL: kill"},
	_SIGBUS:     {_SigPanic + _SigUnblock, "SIGBUS: bus error"},
	_SIGSEGV:    {_SigPanic + _SigUnblock, "SIGSEGV: segmentation violation"},
	_SIGSYS:     {_SigThrow, "SIGSYS: bad system call"},
	_SIGPIPE:    {_SigNotify, "SIGPIPE: write to broken pipe"},
	_SIGALRM:    {_SigNotify, "SIGALRM: alarm clock"},
	_SIGTERM:    {_SigNotify + _SigKill, "SIGTERM: termination"},
	_SIGURG:     {_SigNotify, "SIGURG: urgent condition on socket"},
	_SIGSTOP:    {0, "SIGSTOP: stop"},
	_SIGTSTP:    {_SigNotify + _SigDefault, "SIGTSTP: keyboard stop"},
	_SIGCONT:    {_SigNotify + _SigDefault, "SIGCONT: continue"},
	_SIGCHLD:    {_SigNotify + _SigUnblock, "SIGCHLD: child status has changed"},
	_SIGTTIN:    {_SigNotify + _SigDefault, "SIGTTIN: background read from tty"},
	_SIGTTOU:    {_SigNotify + _SigDefault, "SIGTTOU: background write to tty"},
	_SIGIO:      {_SigNotify, "SIGIO: i/o now possible"},
	_SIGXCPU:    {_SigNotify, "SIGXCPU: cpu limit exceeded"},
	_SIGXFSZ:    {_SigNotify, "SIGXFSZ: file size limit exceeded"},
	_SIGVTALRM:  {_SigNotify, "SIGVTALRM: virtual alarm clock"},
	_SIGPROF:    {_SigNotify + _SigUnblock, "SIGPROF: profiling alarm clock"},
	_SIGWINCH:   {_SigNotify, "SIGWINCH: window size change"},
	_SIGUSR1:    {_SigNotify, "SIGUSR1: user-defined signal 1"},
	_SIGUSR2:    {_SigNotify, "SIGUSR2: user-defined signal 2"},
	32:          {_SigNotify, "signal 32"},
}

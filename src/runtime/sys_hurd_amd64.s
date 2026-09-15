// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "cgo/abi_amd64.h"

// System calls and other sys.stuff for amd64, GNU/Hurd.
//
// Hurd has no usable raw syscall ABI: everything goes through glibc. We call
// libc function pointers (cgo_import_dynamic slots) using the System V AMD64
// ABI. Like Solaris, there is no arch_prctl: glibc owns the TLS base.

// This is needed by asm_amd64.s.
TEXT runtime·settls(SB),NOSPLIT,$0
	RET

// func asmsyscall6(c *libcall)
// Called via asmcgocall, which passes the first argument in DI.
// NOT USING GO CALLING CONVENTION.
TEXT ·asmsyscall6(SB),NOSPLIT|NOFRAME,$0
	MOVQ	DI, BX			// BX = *libcall (callee-saved by C)
	MOVQ	libcall_fn(BX), AX	// AX = &data slot
	MOVQ	0(AX), AX		// AX = resolved libc function address
	MOVQ	libcall_args(BX), R11	// R11 = &args[0]
	MOVQ	0(R11), DI
	MOVQ	8(R11), SI
	MOVQ	16(R11), DX
	MOVQ	24(R11), CX
	MOVQ	32(R11), R8
	MOVQ	40(R11), R9
	MOVQ	AX, R10			// R10 = function address
	XORL	AX, AX			// AL = 0 vector regs (variadic libc)
	SUBQ	$8, SP			// 16-byte align before CALL
	CALL	R10
	ADDQ	$8, SP
	MOVQ	AX, libcall_r1(BX)
	CMPL	AX, $-1
	JNE	asmsyscall6_noerr
	get_tls(CX)
	MOVQ	g(CX), CX
	TESTQ	CX, CX
	JZ	asmsyscall6_noerr
	MOVQ	g_m(CX), CX
	MOVQ	(m_mOS + mOS_perrno)(CX), CX
	MOVQ	0(CX), CX
	MOVQ	CX, libcall_err(BX)
	RET
asmsyscall6_noerr:
	MOVQ	$0, libcall_err(BX)
	RET

// Runs on OS stack, called from runtime.exit when g == nil.
TEXT ·exit1(SB),NOSPLIT|NOFRAME,$0-4
	MOVL	code+0(FP), DI
	MOVQ	·libc_exit(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	INT	$3

// Runs on OS stack, called from runtime.write1 when g == nil.
TEXT ·write2(SB),NOSPLIT|NOFRAME,$0-24
	MOVQ	fd+0(FP), DI
	MOVQ	p+8(FP), SI
	MOVL	n+16(FP), DX
	MOVQ	·libc_write(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	MOVL	AX, ret+20(FP)
	RET

// Runs on OS stack, called from runtime.usleep when g == nil.
TEXT ·usleep1(SB),NOSPLIT|NOFRAME,$0-4
	MOVL	us+0(FP), DI
	MOVQ	·libc_usleep(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	RET

// Runs on OS stack.
TEXT ·osyield1(SB),NOSPLIT|NOFRAME,$0
	MOVQ	·libc_sched_yield(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	RET

// Runs on OS stack, called from runtime.sigprocmask when g == nil.
TEXT ·sigprocmask1(SB),NOSPLIT|NOFRAME,$0-24
	MOVQ	how+0(FP), DI
	MOVQ	new+8(FP), SI
	MOVQ	old+16(FP), DX
	MOVQ	·libc_pthread_sigmask(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	RET

// Runs on OS stack, called from runtime.sigaction when g == nil.
TEXT ·sigaction1(SB),NOSPLIT|NOFRAME,$0-24
	MOVQ	sig+0(FP), DI
	MOVQ	new+8(FP), SI
	MOVQ	old+16(FP), DX
	MOVQ	·libc_sigaction(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	RET

TEXT ·pthread_attr_init1(SB),NOSPLIT|NOFRAME,$0-12
	MOVQ	attr+0(FP), DI
	MOVQ	·libpthread_attr_init(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	MOVL	AX, ret+8(FP)
	RET

TEXT ·pthread_attr_setstacksize1(SB),NOSPLIT|NOFRAME,$0-20
	MOVQ	attr+0(FP), DI
	MOVQ	size+8(FP), SI
	MOVQ	·libpthread_attr_setstacksize(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	MOVL	AX, ret+16(FP)
	RET

TEXT ·pthread_attr_setdetachstate1(SB),NOSPLIT|NOFRAME,$0-16
	MOVQ	attr+0(FP), DI
	MOVL	state+8(FP), SI
	MOVQ	·libpthread_attr_setdetachstate(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	MOVL	AX, ret+12(FP)
	RET

TEXT ·pthread_create1(SB),NOSPLIT|NOFRAME,$0-36
	MOVQ	tid+0(FP), DI
	MOVQ	attr+8(FP), SI
	MOVQ	fn+16(FP), DX
	MOVQ	arg+24(FP), CX
	MOVQ	·libpthread_create(SB), AX
	SUBQ	$8, SP
	CALL	AX
	ADDQ	$8, SP
	MOVL	AX, ret+32(FP)
	RET

// uint32 tstart_hurd(M *newm);
// pthread start routine. Called by glibc with the C ABI (DI = newm).
TEXT ·tstart_hurd(SB),NOSPLIT,$0
	MOVQ	DI, BX			// newm (callee-saved)
	MOVQ	m_g0(BX), DX		// g

	// Make TLS entries point at g and m.
	get_tls(CX)
	MOVQ	DX, g(CX)
	MOVQ	BX, g_m(DX)

	// Lay out new m scheduler stack on the OS stack.
	MOVQ	SP, AX
	MOVQ	AX, (g_stack+stack_hi)(DX)
	SUBQ	$(0x100000), AX
	MOVQ	AX, (g_stack+stack_lo)(DX)
	ADDQ	$const_stackGuard, AX
	MOVQ	AX, g_stackguard0(DX)
	MOVQ	AX, g_stackguard1(DX)

	CLD
	CALL	runtime·stackcheck(SB)
	CALL	runtime·mstart(SB)

	XORL	AX, AX
	RET

// Called by glibc using the C ABI: handler(sig, info, ucontext).
TEXT runtime·sigtramp(SB),NOSPLIT|TOPFRAME|NOFRAME,$0
	// DI=sig, SI=info, DX=ucontext. Call the Go signal handler via the
	// ABI0 wrapper, which preserves the C callee-saved registers.
	SUBQ	$24, SP
	MOVL	DI, 0(SP)
	MOVQ	SI, 8(SP)
	MOVQ	DX, 16(SP)
	CALL	runtime·sigtrampgo(SB)
	ADDQ	$24, SP
	RET

// func sigfwd(fn uintptr, sig uint32, info *siginfo, ctx unsafe.Pointer)
TEXT runtime·sigfwd(SB),NOSPLIT,$0-32
	MOVQ	fn+0(FP), AX
	MOVL	sig+8(FP), DI
	MOVQ	info+16(FP), SI
	MOVQ	ctx+24(FP), DX
	MOVQ	SP, BX
	ANDQ	$~15, SP
	CALL	AX
	MOVQ	BX, SP
	RET

// Data slots holding resolved libc function addresses. The dynamic
// R_X86_64_64 relocations live here (in .data), not in .text. On amd64 the
// slot is 8 bytes wide.
DATA ·libc__errno_location+0(SB)/8, $imp_libc__errno_location(SB)
GLOBL ·libc__errno_location(SB), NOPTR, $8
DATA ·libc_clock_gettime+0(SB)/8, $imp_libc_clock_gettime(SB)
GLOBL ·libc_clock_gettime(SB), NOPTR, $8
DATA ·libc_close+0(SB)/8, $imp_libc_close(SB)
GLOBL ·libc_close(SB), NOPTR, $8
DATA ·libc_exit+0(SB)/8, $imp_libc_exit(SB)
GLOBL ·libc_exit(SB), NOPTR, $8
DATA ·libc_fcntl+0(SB)/8, $imp_libc_fcntl(SB)
GLOBL ·libc_fcntl(SB), NOPTR, $8
DATA ·libc_getpid+0(SB)/8, $imp_libc_getpid(SB)
GLOBL ·libc_getpid(SB), NOPTR, $8
DATA ·libc_getuid+0(SB)/8, $imp_libc_getuid(SB)
GLOBL ·libc_getuid(SB), NOPTR, $8
DATA ·libc_geteuid+0(SB)/8, $imp_libc_geteuid(SB)
GLOBL ·libc_geteuid(SB), NOPTR, $8
DATA ·libc_getgid+0(SB)/8, $imp_libc_getgid(SB)
GLOBL ·libc_getgid(SB), NOPTR, $8
DATA ·libc_getegid+0(SB)/8, $imp_libc_getegid(SB)
GLOBL ·libc_getegid(SB), NOPTR, $8
DATA ·libc_kill+0(SB)/8, $imp_libc_kill(SB)
GLOBL ·libc_kill(SB), NOPTR, $8
DATA ·libc_madvise+0(SB)/8, $imp_libc_madvise(SB)
GLOBL ·libc_madvise(SB), NOPTR, $8
DATA ·libc_malloc+0(SB)/8, $imp_libc_malloc(SB)
GLOBL ·libc_malloc(SB), NOPTR, $8
DATA ·libc_mmap+0(SB)/8, $imp_libc_mmap(SB)
GLOBL ·libc_mmap(SB), NOPTR, $8
DATA ·libc_mprotect+0(SB)/8, $imp_libc_mprotect(SB)
GLOBL ·libc_mprotect(SB), NOPTR, $8
DATA ·libc_munmap+0(SB)/8, $imp_libc_munmap(SB)
GLOBL ·libc_munmap(SB), NOPTR, $8
DATA ·libc_open+0(SB)/8, $imp_libc_open(SB)
GLOBL ·libc_open(SB), NOPTR, $8
DATA ·libc_pipe+0(SB)/8, $imp_libc_pipe(SB)
GLOBL ·libc_pipe(SB), NOPTR, $8
DATA ·libc_poll+0(SB)/8, $imp_libc_poll(SB)
GLOBL ·libc_poll(SB), NOPTR, $8
DATA ·libc_raise+0(SB)/8, $imp_libc_raise(SB)
GLOBL ·libc_raise(SB), NOPTR, $8
DATA ·libc_read+0(SB)/8, $imp_libc_read(SB)
GLOBL ·libc_read(SB), NOPTR, $8
DATA ·libc_sched_yield+0(SB)/8, $imp_libc_sched_yield(SB)
GLOBL ·libc_sched_yield(SB), NOPTR, $8
DATA ·libc_sem_init+0(SB)/8, $imp_libc_sem_init(SB)
GLOBL ·libc_sem_init(SB), NOPTR, $8
DATA ·libc_sem_post+0(SB)/8, $imp_libc_sem_post(SB)
GLOBL ·libc_sem_post(SB), NOPTR, $8
DATA ·libc_sem_timedwait+0(SB)/8, $imp_libc_sem_timedwait(SB)
GLOBL ·libc_sem_timedwait(SB), NOPTR, $8
DATA ·libc_sem_wait+0(SB)/8, $imp_libc_sem_wait(SB)
GLOBL ·libc_sem_wait(SB), NOPTR, $8
DATA ·libc_setitimer+0(SB)/8, $imp_libc_setitimer(SB)
GLOBL ·libc_setitimer(SB), NOPTR, $8
DATA ·libc_sigaction+0(SB)/8, $imp_libc_sigaction(SB)
GLOBL ·libc_sigaction(SB), NOPTR, $8
DATA ·libc_sigaltstack+0(SB)/8, $imp_libc_sigaltstack(SB)
GLOBL ·libc_sigaltstack(SB), NOPTR, $8
DATA ·libc_sysconf+0(SB)/8, $imp_libc_sysconf(SB)
GLOBL ·libc_sysconf(SB), NOPTR, $8
DATA ·libc_usleep+0(SB)/8, $imp_libc_usleep(SB)
GLOBL ·libc_usleep(SB), NOPTR, $8
DATA ·libc_write+0(SB)/8, $imp_libc_write(SB)
GLOBL ·libc_write(SB), NOPTR, $8
DATA ·libc_chdir+0(SB)/8, $imp_libc_chdir(SB)
GLOBL ·libc_chdir(SB), NOPTR, $8
DATA ·libc_chroot+0(SB)/8, $imp_libc_chroot(SB)
GLOBL ·libc_chroot(SB), NOPTR, $8
DATA ·libc_dup2+0(SB)/8, $imp_libc_dup2(SB)
GLOBL ·libc_dup2(SB), NOPTR, $8
DATA ·libc_execve+0(SB)/8, $imp_libc_execve(SB)
GLOBL ·libc_execve(SB), NOPTR, $8
DATA ·libc_fork+0(SB)/8, $imp_libc_fork(SB)
GLOBL ·libc_fork(SB), NOPTR, $8
DATA ·libc_ioctl+0(SB)/8, $imp_libc_ioctl(SB)
GLOBL ·libc_ioctl(SB), NOPTR, $8
DATA ·libc_setgid+0(SB)/8, $imp_libc_setgid(SB)
GLOBL ·libc_setgid(SB), NOPTR, $8
DATA ·libc_setgroups+0(SB)/8, $imp_libc_setgroups(SB)
GLOBL ·libc_setgroups(SB), NOPTR, $8
DATA ·libc_setrlimit+0(SB)/8, $imp_libc_setrlimit(SB)
GLOBL ·libc_setrlimit(SB), NOPTR, $8
DATA ·libc_setsid+0(SB)/8, $imp_libc_setsid(SB)
GLOBL ·libc_setsid(SB), NOPTR, $8
DATA ·libc_setuid+0(SB)/8, $imp_libc_setuid(SB)
GLOBL ·libc_setuid(SB), NOPTR, $8
DATA ·libc_setpgid+0(SB)/8, $imp_libc_setpgid(SB)
GLOBL ·libc_setpgid(SB), NOPTR, $8
DATA ·libc_pthread_sigmask+0(SB)/8, $imp_libc_pthread_sigmask(SB)
GLOBL ·libc_pthread_sigmask(SB), NOPTR, $8
DATA ·libpthread_attr_destroy+0(SB)/8, $imp_libpthread_attr_destroy(SB)
GLOBL ·libpthread_attr_destroy(SB), NOPTR, $8
DATA ·libpthread_attr_init+0(SB)/8, $imp_libpthread_attr_init(SB)
GLOBL ·libpthread_attr_init(SB), NOPTR, $8
DATA ·libpthread_attr_setstacksize+0(SB)/8, $imp_libpthread_attr_setstacksize(SB)
GLOBL ·libpthread_attr_setstacksize(SB), NOPTR, $8
DATA ·libpthread_attr_setdetachstate+0(SB)/8, $imp_libpthread_attr_setdetachstate(SB)
GLOBL ·libpthread_attr_setdetachstate(SB), NOPTR, $8
DATA ·libpthread_create+0(SB)/8, $imp_libpthread_create(SB)
GLOBL ·libpthread_create(SB), NOPTR, $8
DATA ·libpthread_self+0(SB)/8, $imp_libpthread_self(SB)
GLOBL ·libpthread_self(SB), NOPTR, $8
DATA ·libpthread_kill+0(SB)/8, $imp_libpthread_kill(SB)
GLOBL ·libpthread_kill(SB), NOPTR, $8

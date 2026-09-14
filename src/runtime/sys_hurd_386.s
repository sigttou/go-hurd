// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

// System calls and other sys.stuff for 386, GNU/Hurd.
//
// Hurd has no usable raw syscall ABI: everything goes through glibc. We call
// libc function pointers (cgo_import_dynamic slots) using the cdecl ABI.

// func asmsyscall6(c *libcall)
// Called via asmcgocall. c.args points at a [6]uintptr array; we push all six
// (the callee reads only as many as it needs). NOT USING GO CALLING
// CONVENTION: this function is called with the C ABI by asmcgocall.
TEXT ·asmsyscall6(SB),NOSPLIT,$0
	MOVL	c+0(FP), BX		// BX = *libcall (callee-saved by C)
	MOVL	libcall_fn(BX), AX
	MOVL	0(AX), AX		// AX = resolved libc function
	MOVL	libcall_args(BX), SI	// SI = &args[0]
	PUSHL	20(SI)
	PUSHL	16(SI)
	PUSHL	12(SI)
	PUSHL	8(SI)
	PUSHL	4(SI)
	PUSHL	0(SI)
	SUBL	$12, SP			// 16-byte align before CALL
	CALL	AX
	ADDL	$12, SP
	POPL	CX
	POPL	CX
	POPL	CX
	POPL	CX
	POPL	CX
	POPL	CX
	MOVL	AX, libcall_r1(BX)
	CMPL	AX, $-1
	JNE	asmsyscall6_noerr
	get_tls(CX)
	MOVL	g(CX), CX
	TESTL	CX, CX
	JZ	asmsyscall6_noerr
	MOVL	g_m(CX), CX
	MOVL	(m_mOS + mOS_perrno)(CX), CX
	MOVL	0(CX), CX
	MOVL	CX, libcall_err(BX)
	RET
asmsyscall6_noerr:
	MOVL	$0, libcall_err(BX)
	RET

// Runs on OS stack, called from runtime.exit when g == nil.
TEXT ·exit1(SB),NOSPLIT,$0-4
	MOVL	code+0(FP), AX
	MOVL	libc_exit(SB), CX
	MOVL	SP, DX
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	AX, 0(SP)
	CALL	CX
	MOVL	DX, SP
	INT	$3

// Runs on OS stack, called from runtime.write1 when g == nil.
TEXT ·write2(SB),NOSPLIT,$0-16
	MOVL	fd+0(FP), AX
	MOVL	p+4(FP), BX
	MOVL	n+8(FP), CX
	MOVL	libc_write(SB), DX
	MOVL	SP, SI
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	CX, 8(SP)
	MOVL	BX, 4(SP)
	MOVL	AX, 0(SP)
	CALL	DX
	MOVL	SI, SP
	MOVL	AX, ret+12(FP)
	RET

// Runs on OS stack, called from runtime.usleep when g == nil.
TEXT ·usleep1(SB),NOSPLIT,$0-4
	MOVL	us+0(FP), AX
	MOVL	libc_usleep(SB), CX
	MOVL	SP, DX
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	AX, 0(SP)
	CALL	CX
	MOVL	DX, SP
	RET

// Runs on OS stack.
TEXT ·osyield1(SB),NOSPLIT,$0
	MOVL	libc_sched_yield(SB), AX
	MOVL	SP, DX
	SUBL	$16, SP
	ANDL	$~15, SP
	CALL	AX
	MOVL	DX, SP
	RET

// Runs on OS stack, called from runtime.sigprocmask when g == nil.
TEXT ·sigprocmask1(SB),NOSPLIT,$0-12
	MOVL	how+0(FP), AX
	MOVL	new+4(FP), BX
	MOVL	old+8(FP), CX
	MOVL	libc_pthread_sigmask(SB), DX
	MOVL	SP, SI
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	CX, 8(SP)
	MOVL	BX, 4(SP)
	MOVL	AX, 0(SP)
	CALL	DX
	MOVL	SI, SP
	RET

// Runs on OS stack, called from runtime.sigaction when g == nil.
TEXT ·sigaction1(SB),NOSPLIT,$0-12
	MOVL	sig+0(FP), AX
	MOVL	new+4(FP), BX
	MOVL	old+8(FP), CX
	MOVL	libc_sigaction(SB), DX
	MOVL	SP, SI
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	CX, 8(SP)
	MOVL	BX, 4(SP)
	MOVL	AX, 0(SP)
	CALL	DX
	MOVL	SI, SP
	RET

TEXT ·pthread_attr_init1(SB),NOSPLIT,$0-8
	MOVL	attr+0(FP), AX
	MOVL	libpthread_attr_init(SB), CX
	MOVL	SP, DX
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	AX, 0(SP)
	CALL	CX
	MOVL	DX, SP
	MOVL	AX, ret+4(FP)
	RET

TEXT ·pthread_attr_setstacksize1(SB),NOSPLIT,$0-12
	MOVL	attr+0(FP), AX
	MOVL	size+4(FP), BX
	MOVL	libpthread_attr_setstacksize(SB), CX
	MOVL	SP, DX
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	BX, 4(SP)
	MOVL	AX, 0(SP)
	CALL	CX
	MOVL	DX, SP
	MOVL	AX, ret+8(FP)
	RET

TEXT ·pthread_attr_setdetachstate1(SB),NOSPLIT,$0-12
	MOVL	attr+0(FP), AX
	MOVL	state+4(FP), BX
	MOVL	libpthread_attr_setdetachstate(SB), CX
	MOVL	SP, DX
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	BX, 4(SP)
	MOVL	AX, 0(SP)
	CALL	CX
	MOVL	DX, SP
	MOVL	AX, ret+8(FP)
	RET

TEXT ·pthread_create1(SB),NOSPLIT,$0-20
	MOVL	tid+0(FP), AX
	MOVL	attr+4(FP), BX
	MOVL	fn+8(FP), CX
	MOVL	arg+12(FP), SI
	MOVL	libpthread_create(SB), DX
	MOVL	SP, DI
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	SI, 12(SP)
	MOVL	CX, 8(SP)
	MOVL	BX, 4(SP)
	MOVL	AX, 0(SP)
	CALL	DX
	MOVL	DI, SP
	MOVL	AX, ret+16(FP)
	RET

// uint32 tstart_hurd(M *newm);
// pthread start routine. Called by glibc with the C ABI.
TEXT ·tstart_hurd(SB),NOSPLIT,$0-8
	MOVL	newm+0(FP), DI		// newm
	MOVL	m_g0(DI), DX		// g

	// Make TLS entries point at g and m.
	get_tls(CX)
	MOVL	DX, g(CX)
	MOVL	DI, g_m(DX)

	// Lay out new m scheduler stack on the OS stack.
	MOVL	SP, AX
	MOVL	AX, (g_stack+stack_hi)(DX)
	SUBL	$0x100000, AX
	MOVL	AX, (g_stack+stack_lo)(DX)
	ADDL	$const_stackGuard, AX
	MOVL	AX, g_stackguard0(DX)
	MOVL	AX, g_stackguard1(DX)

	CLD
	CALL	runtime·stackcheck(SB)
	CALL	runtime·mstart(SB)

	XORL	AX, AX
	MOVL	AX, ret+4(FP)
	RET

// Called by glibc using the C ABI: handler(sig, info, ucontext).
TEXT runtime·sigtramp(SB),NOSPLIT|TOPFRAME,$12
	NOP	SP
	MOVL	16(SP), BX	// sig
	MOVL	BX, 0(SP)
	MOVL	20(SP), BX	// info
	MOVL	BX, 4(SP)
	MOVL	24(SP), BX	// ucontext
	MOVL	BX, 8(SP)
	CALL	runtime·sigtrampgo(SB)
	RET

// func sigfwd(fn uintptr, sig uint32, info *siginfo, ctx unsafe.Pointer)
TEXT runtime·sigfwd(SB),NOSPLIT,$0-16
	MOVL	fn+0(FP), AX
	MOVL	sig+4(FP), BX
	MOVL	info+8(FP), CX
	MOVL	ctx+12(FP), DX
	MOVL	SP, SI
	SUBL	$16, SP
	ANDL	$~15, SP
	MOVL	BX, 0(SP)
	MOVL	CX, 4(SP)
	MOVL	DX, 8(SP)
	MOVL	SI, 12(SP)
	CALL	AX
	MOVL	12(SP), AX
	MOVL	AX, SP
	RET

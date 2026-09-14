// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package runtime

// This file contains the main Hurd runtime libc bindings. Like the AIX port,
// Hurd has no raw syscall instruction that is usable from userland: all
// operations go through glibc. We call libc through //go:cgo_import_dynamic
// function-pointer slots and the 386 cdecl trampoline in sys_hurd_386.s.

import (
	"internal/abi"
	"internal/runtime/sys"
	"unsafe"
)

//go:cgo_import_dynamic imp_libc__errno_location __errno_location "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_clock_gettime clock_gettime "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_close close "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_exit _exit "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_fcntl fcntl "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_getpid getpid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_getuid getuid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_geteuid geteuid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_getgid getgid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_getegid getegid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_kill kill "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_madvise madvise "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_malloc malloc "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_mmap mmap "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_mprotect mprotect "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_munmap munmap "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_open open "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_pipe pipe "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_poll poll "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_raise raise "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_read read "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_sched_yield sched_yield "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_sem_init sem_init "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_sem_post sem_post "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_sem_timedwait sem_timedwait "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_sem_wait sem_wait "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_setitimer setitimer "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_sigaction sigaction "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_sigaltstack sigaltstack "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_sysconf sysconf "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_usleep usleep "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_write write "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_chdir chdir "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_chroot chroot "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_dup2 dup2 "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_execve execve "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_fork fork "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_ioctl ioctl "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_setgid setgid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_setgroups setgroups "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_setrlimit setrlimit "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_setsid setsid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_setuid setuid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_setpgid setpgid "libc.so.0.3"
//go:cgo_import_dynamic imp_libc_pthread_sigmask pthread_sigmask "libc.so.0.3"

//go:cgo_import_dynamic imp_libpthread_attr_destroy pthread_attr_destroy "libc.so.0.3"
//go:cgo_import_dynamic imp_libpthread_attr_init pthread_attr_init "libc.so.0.3"
//go:cgo_import_dynamic imp_libpthread_attr_setstacksize pthread_attr_setstacksize "libc.so.0.3"
//go:cgo_import_dynamic imp_libpthread_attr_setdetachstate pthread_attr_setdetachstate "libc.so.0.3"
//go:cgo_import_dynamic imp_libpthread_create pthread_create "libc.so.0.3"
//go:cgo_import_dynamic imp_libpthread_self pthread_self "libc.so.0.3"
//go:cgo_import_dynamic imp_libpthread_kill pthread_kill "libc.so.0.3"

var (
	libc__errno_location,
	libc_clock_gettime,
	libc_close,
	libc_exit,
	libc_fcntl,
	libc_getpid,
	libc_getuid,
	libc_geteuid,
	libc_getgid,
	libc_getegid,
	libc_kill,
	libc_madvise,
	libc_malloc,
	libc_mmap,
	libc_mprotect,
	libc_munmap,
	libc_open,
	libc_pipe,
	libc_poll,
	libc_raise,
	libc_read,
	libc_sched_yield,
	libc_sem_init,
	libc_sem_post,
	libc_sem_timedwait,
	libc_sem_wait,
	libc_setitimer,
	libc_sigaction,
	libc_sigaltstack,
	libc_sysconf,
	libc_usleep,
	libc_write,
	libc_chdir,
	libc_chroot,
	libc_dup2,
	libc_execve,
	libc_fork,
	libc_ioctl,
	libc_setgid,
	libc_setgroups,
	libc_setrlimit,
	libc_setsid,
	libc_setuid,
	libc_setpgid,
	libc_pthread_sigmask,
	libpthread_attr_destroy,
	libpthread_attr_init,
	libpthread_attr_setstacksize,
	libpthread_attr_setdetachstate,
	libpthread_create,
	libpthread_self,
	libpthread_kill libFunc
)

type libFunc uintptr

// asmsyscall6 is the cdecl trampoline defined in sys_hurd_386.s. Unlike AIX's
// function descriptor, on 386 we pass the code address via abi.FuncPCABI0.
// It always receives a pointer to a 6-element argument array.
//
//go:noescape
func asmsyscall6(c *libcall)

// asmcgocall trampolines used before the runtime (and g) is initialized.
func exit1(code int32)
func write2(fd, p uintptr, n int32) int32
func usleep1(us uint32)
func osyield1()
func sigaction1(sig, new, old uintptr)
func sigprocmask1(how, new, old uintptr)
func pthread_attr_init1(attr uintptr) int32
func pthread_attr_setstacksize1(attr uintptr, size uintptr) int32
func pthread_attr_setdetachstate1(attr uintptr, state int32) int32
func pthread_create1(tid, attr, fn, arg uintptr) int32

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscallN(fn *libFunc, args *[6]uintptr) (r, err uintptr) {
	gp := getg()
	mp := gp.m
	resetLibcall := true
	if mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = sys.GetCallerPC()
		// sp must be the last, because once async cpu profiler finds
		// all three values to be non-zero, it will use them
		mp.libcallsp = sys.GetCallerSP()
	} else {
		resetLibcall = false
	}

	c := libcall{
		fn:   uintptr(unsafe.Pointer(fn)),
		n:    6,
		args: uintptr(unsafe.Pointer(&args[0])),
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nosplit
func syscall0(fn *libFunc) (r, err uintptr) {
	var a [6]uintptr
	return syscallN(fn, &a)
}

//go:nosplit
func syscall1(fn *libFunc, a0 uintptr) (r, err uintptr) {
	a := [6]uintptr{a0}
	return syscallN(fn, &a)
}

//go:nosplit
func syscall2(fn *libFunc, a0, a1 uintptr) (r, err uintptr) {
	a := [6]uintptr{a0, a1}
	return syscallN(fn, &a)
}

//go:nosplit
func syscall3(fn *libFunc, a0, a1, a2 uintptr) (r, err uintptr) {
	a := [6]uintptr{a0, a1, a2}
	return syscallN(fn, &a)
}

//go:nosplit
func syscall4(fn *libFunc, a0, a1, a2, a3 uintptr) (r, err uintptr) {
	a := [6]uintptr{a0, a1, a2, a3}
	return syscallN(fn, &a)
}

//go:nosplit
func syscall5(fn *libFunc, a0, a1, a2, a3, a4 uintptr) (r, err uintptr) {
	a := [6]uintptr{a0, a1, a2, a3, a4}
	return syscallN(fn, &a)
}

//go:nosplit
func syscall6(fn *libFunc, a0, a1, a2, a3, a4, a5 uintptr) (r, err uintptr) {
	a := [6]uintptr{a0, a1, a2, a3, a4, a5}
	return syscallN(fn, &a)
}

//go:nosplit
func exit(code int32) {
	gp := getg()
	if gp != nil {
		syscall1(&libc_exit, uintptr(code))
		return
	}
	exit1(code)
}

//go:nosplit
func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
	gp := getg()
	if gp != nil {
		r, errno := syscall3(&libc_write, fd, uintptr(p), uintptr(n))
		if int32(r) < 0 {
			return -int32(errno)
		}
		return int32(r)
	}
	return write2(fd, uintptr(p), n)
}

//go:nosplit
func read(fd int32, p unsafe.Pointer, n int32) int32 {
	r, errno := syscall3(&libc_read, uintptr(fd), uintptr(p), uintptr(n))
	if int32(r) < 0 {
		return -int32(errno)
	}
	return int32(r)
}

//go:nosplit
func open(name *byte, mode, perm int32) int32 {
	r, _ := syscall3(&libc_open, uintptr(unsafe.Pointer(name)), uintptr(mode), uintptr(perm))
	return int32(r)
}

//go:nosplit
func closefd(fd int32) int32 {
	r, _ := syscall1(&libc_close, uintptr(fd))
	return int32(r)
}

//go:nosplit
func pipe() (r, w int32, errno int32) {
	var p [2]int32
	_, err := syscall1(&libc_pipe, uintptr(noescape(unsafe.Pointer(&p[0]))))
	return p[0], p[1], int32(err)
}

//go:nosplit
func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (unsafe.Pointer, int) {
	r, err0 := syscall6(&libc_mmap, uintptr(addr), uintptr(n), uintptr(prot), uintptr(flags), uintptr(fd), uintptr(off))
	if r == ^uintptr(0) {
		return nil, int(err0)
	}
	return unsafe.Pointer(r), int(err0)
}

//go:nosplit
func mprotect(addr unsafe.Pointer, n uintptr, prot int32) (unsafe.Pointer, int) {
	r, err0 := syscall3(&libc_mprotect, uintptr(addr), uintptr(n), uintptr(prot))
	if r == ^uintptr(0) {
		return nil, int(err0)
	}
	return unsafe.Pointer(r), int(err0)
}

//go:nosplit
func munmap(addr unsafe.Pointer, n uintptr) {
	r, err := syscall2(&libc_munmap, uintptr(addr), uintptr(n))
	if int32(r) == -1 {
		println("syscall munmap failed: ", hex(err))
		throw("syscall munmap")
	}
}

//go:nosplit
func madvise(addr unsafe.Pointer, n uintptr, flags int32) {
	// madvise is advisory. GNU/Hurd's glibc returns ENOSYS for
	// MADV_DONTNEED, so failures must not be fatal.
	syscall3(&libc_madvise, uintptr(addr), uintptr(n), uintptr(flags))
}

//go:nosplit
func sigaction(sig uintptr, new, old *sigactiont) {
	gp := getg()
	if gp != nil {
		r, err := syscall3(&libc_sigaction, sig, uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
		if int32(r) == -1 {
			println("Sigaction failed for sig: ", sig, " with error:", hex(err))
			throw("syscall sigaction")
		}
		return
	}
	sigaction1(sig, uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
}

//go:nosplit
func sigaltstack(new, old *stackt) {
	r, err := syscall2(&libc_sigaltstack, uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
	if int32(r) == -1 {
		println("syscall sigaltstack failed: ", hex(err))
		throw("syscall sigaltstack")
	}
}

//go:nosplit
func usleep_no_g(us uint32) {
	usleep1(us)
}

//go:nosplit
func usleep(us uint32) {
	r, err := syscall1(&libc_usleep, uintptr(us))
	if int32(r) == -1 {
		println("syscall usleep failed: ", hex(err))
		throw("syscall usleep")
	}
}

//go:nosplit
func clock_gettime(clockid int32, tp *timespec) int32 {
	r, _ := syscall2(&libc_clock_gettime, uintptr(clockid), uintptr(unsafe.Pointer(tp)))
	return int32(r)
}

//go:nosplit
func setitimer(mode int32, new, old *itimerval) {
	r, err := syscall3(&libc_setitimer, uintptr(mode), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
	if int32(r) == -1 {
		println("syscall setitimer failed: ", hex(err))
		throw("syscall setitimer")
	}
}

//go:nosplit
func malloc(size uintptr) unsafe.Pointer {
	r, _ := syscall1(&libc_malloc, size)
	return unsafe.Pointer(r)
}

//go:nosplit
func sem_init(sem *semt, pshared int32, value uint32) int32 {
	r, _ := syscall3(&libc_sem_init, uintptr(unsafe.Pointer(sem)), uintptr(pshared), uintptr(value))
	return int32(r)
}

//go:nosplit
func sem_wait(sem *semt) (int32, int32) {
	r, err := syscall1(&libc_sem_wait, uintptr(unsafe.Pointer(sem)))
	return int32(r), int32(err)
}

//go:nosplit
func sem_post(sem *semt) int32 {
	r, _ := syscall1(&libc_sem_post, uintptr(unsafe.Pointer(sem)))
	return int32(r)
}

//go:nosplit
func sem_timedwait(sem *semt, timeout *timespec) (int32, int32) {
	r, err := syscall2(&libc_sem_timedwait, uintptr(unsafe.Pointer(sem)), uintptr(unsafe.Pointer(timeout)))
	return int32(r), int32(err)
}

//go:nosplit
func raise(sig uint32) {
	r, err := syscall1(&libc_raise, uintptr(sig))
	if int32(r) == -1 {
		println("syscall raise failed: ", hex(err))
		throw("syscall raise")
	}
}

//go:nosplit
func raiseproc(sig uint32) {
	pid, err := syscall0(&libc_getpid)
	if int32(pid) == -1 {
		println("syscall getpid failed: ", hex(err))
		throw("syscall raiseproc")
	}
	syscall2(&libc_kill, pid, uintptr(sig))
}

//go:nosplit
func osyield_no_g() {
	osyield1()
}

//go:nosplit
func osyield() {
	r, err := syscall0(&libc_sched_yield)
	if int32(r) == -1 {
		println("syscall osyield failed: ", hex(err))
		throw("syscall osyield")
	}
}

//go:nosplit
func sysconf(name int32) uintptr {
	r, _ := syscall1(&libc_sysconf, uintptr(name))
	if int32(r) == -1 {
		throw("syscall sysconf")
	}
	return r
}

// pthread functions return their error code in the main return value, so the
// err return from syscall must not be used.

//go:nosplit
func pthread_attr_destroy(attr *pthread_attr) int32 {
	r, _ := syscall1(&libpthread_attr_destroy, uintptr(unsafe.Pointer(attr)))
	return int32(r)
}

//go:nosplit
func pthread_attr_init(attr *pthread_attr) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall1(&libpthread_attr_init, uintptr(unsafe.Pointer(attr)))
		return int32(r)
	}
	return pthread_attr_init1(uintptr(unsafe.Pointer(attr)))
}

//go:nosplit
func pthread_attr_setdetachstate(attr *pthread_attr, state int32) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall2(&libpthread_attr_setdetachstate, uintptr(unsafe.Pointer(attr)), uintptr(state))
		return int32(r)
	}
	return pthread_attr_setdetachstate1(uintptr(unsafe.Pointer(attr)), state)
}

//go:nosplit
func pthread_attr_setstacksize(attr *pthread_attr, size uint64) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall2(&libpthread_attr_setstacksize, uintptr(unsafe.Pointer(attr)), uintptr(size))
		return int32(r)
	}
	return pthread_attr_setstacksize1(uintptr(unsafe.Pointer(attr)), uintptr(size))
}

//go:nosplit
func pthread_create(tid *pthread, attr *pthread_attr, fn uintptr, arg unsafe.Pointer) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall4(&libpthread_create, uintptr(unsafe.Pointer(tid)), uintptr(unsafe.Pointer(attr)), fn, uintptr(arg))
		return int32(r)
	}
	return pthread_create1(uintptr(unsafe.Pointer(tid)), uintptr(unsafe.Pointer(attr)), fn, uintptr(arg))
}

//go:nosplit
func sigprocmask(how int32, new, old *sigset) {
	gp := getg()
	if gp != nil && gp.m != nil {
		r, err := syscall3(&libc_pthread_sigmask, uintptr(how), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
		if int32(r) != 0 {
			println("syscall pthread_sigmask failed: ", hex(err))
			throw("syscall pthread_sigmask")
		}
		return
	}
	sigprocmask1(uintptr(how), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
}

//go:nosplit
func pthread_self() pthread {
	r, _ := syscall0(&libpthread_self)
	return pthread(r)
}

//go:nosplit
func signalM(mp *m, sig int) {
	syscall2(&libpthread_kill, uintptr(pthread(mp.procid)), uintptr(sig))
}

// syscall_syscall6 and syscall_rawSyscall6 are used (via linkname) by the
// syscall package. fn is the address of a data slot holding the resolved libc
// function address.
//
//go:linkname syscall_syscall6
//go:cgo_unsafe_args
func syscall_syscall6(fn, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	a := [6]uintptr{a1, a2, a3, a4, a5, a6}
	r, e := syscallN((*libFunc)(unsafe.Pointer(fn)), &a)
	return r, 0, e
}

//go:linkname syscall_rawSyscall6
//go:cgo_unsafe_args
func syscall_rawSyscall6(fn, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	return syscall_syscall6(fn, nargs, a1, a2, a3, a4, a5, a6)
}

// The following functions are used by package syscall's exec_libc.go via
// linkname. They must not split the stack, because they run after fork in
// the child process. Each one calls a libc function through a data slot.

//go:linkname syscall_chdir syscall.chdir
//go:nosplit
func syscall_chdir(path uintptr) (err uintptr) {
	_, err = syscall1(&libc_chdir, path)
	return
}

//go:linkname syscall_chroot1 syscall.chroot1
//go:nosplit
func syscall_chroot1(path uintptr) (err uintptr) {
	_, err = syscall1(&libc_chroot, path)
	return
}

// like close, but must not split stack, for fork.
//
//go:linkname syscall_closeFD syscall.closeFD
//go:nosplit
func syscall_closeFD(fd uintptr) (err uintptr) {
	_, err = syscall1(&libc_close, fd)
	return
}

//go:linkname syscall_dup2child syscall.dup2child
//go:nosplit
func syscall_dup2child(old, new uintptr) (val, err uintptr) {
	val, err = syscall2(&libc_dup2, old, new)
	return
}

//go:linkname syscall_execve syscall.execve
//go:nosplit
func syscall_execve(path, argv, envp uintptr) (err uintptr) {
	_, err = syscall3(&libc_execve, path, argv, envp)
	return
}

// like exit, but must not split stack, for fork.
//
//go:linkname syscall_exit syscall.exit
//go:nosplit
func syscall_exit(code uintptr) {
	syscall1(&libc_exit, code)
}

//go:linkname syscall_fcntl1 syscall.fcntl1
//go:nosplit
func syscall_fcntl1(fd, cmd, arg uintptr) (val, err uintptr) {
	val, err = syscall3(&libc_fcntl, fd, cmd, arg)
	return
}

//go:linkname syscall_forkx syscall.forkx
//go:nosplit
func syscall_forkx(flags uintptr) (pid uintptr, err uintptr) {
	// Hurd's fork has no flags argument; the extra argument is ignored.
	pid, err = syscall0(&libc_fork)
	return
}

//go:linkname syscall_getpid syscall.getpid
//go:nosplit
func syscall_getpid() (pid, err uintptr) {
	pid, err = syscall0(&libc_getpid)
	return
}

//go:linkname syscall_ioctl syscall.ioctl
//go:nosplit
func syscall_ioctl(fd, req, arg uintptr) (err uintptr) {
	_, err = syscall3(&libc_ioctl, fd, req, arg)
	return
}

//go:linkname syscall_setgid syscall.setgid
//go:nosplit
func syscall_setgid(gid uintptr) (err uintptr) {
	_, err = syscall1(&libc_setgid, gid)
	return
}

//go:linkname syscall_setgroups1 syscall.setgroups1
//go:nosplit
func syscall_setgroups1(ngid, gid uintptr) (err uintptr) {
	_, err = syscall2(&libc_setgroups, ngid, gid)
	return
}

//go:linkname syscall_setrlimit1 syscall.setrlimit1
//go:nosplit
func syscall_setrlimit1(which uintptr, lim unsafe.Pointer) (err uintptr) {
	_, err = syscall2(&libc_setrlimit, which, uintptr(lim))
	return
}

//go:linkname syscall_setsid syscall.setsid
//go:nosplit
func syscall_setsid() (pid, err uintptr) {
	pid, err = syscall0(&libc_setsid)
	return
}

//go:linkname syscall_setuid syscall.setuid
//go:nosplit
func syscall_setuid(uid uintptr) (err uintptr) {
	_, err = syscall1(&libc_setuid, uid)
	return
}

//go:linkname syscall_setpgid syscall.setpgid
//go:nosplit
func syscall_setpgid(pid, pgid uintptr) (err uintptr) {
	_, err = syscall2(&libc_setpgid, pid, pgid)
	return
}

//go:linkname syscall_write1 syscall.write1
//go:nosplit
func syscall_write1(fd, buf, nbyte uintptr) (n, err uintptr) {
	n, err = syscall3(&libc_write, fd, buf, nbyte)
	return
}

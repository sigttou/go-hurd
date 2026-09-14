// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// GNU/Hurd system calls. Hurd has no usable raw syscall interface; all
// operations go through glibc. This file is compiled as ordinary Go code,
// but it is also input to mksyscall, which parses the //sys lines and
// generates system call stubs.

//go:build hurd

package syscall

import (
	"sync"
	"unsafe"
)

// Syscall is not directly usable on Hurd, which has no raw system call
// interface. It exists only to satisfy build dependencies and returns EINVAL.
func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno) {
	return 0, 0, EINVAL
}

// Syscall6 is not directly usable on Hurd, see Syscall.
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) {
	return 0, 0, EINVAL
}

// RawSyscall is not available on Hurd.
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno) {
	panic("RawSyscall not available on Hurd")
}

// RawSyscall6 is not available on Hurd.
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) {
	panic("RawSyscall6 not available on Hurd")
}

// Implemented in runtime/sys_hurd_386.s.
func rawSyscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

const (
	_F_DUP2FD_CLOEXEC = 0
)

/*
 * Wrapped
 */

func Access(path string, mode uint32) (err error) {
	return Faccessat(_AT_FDCWD, path, mode, 0)
}

//sys	fcntl(fd int, cmd int, arg int) (val int, err error)
//sys	Dup2(old int, new int) (err error)
//sysnb	pipe(p *[2]_C_int) (err error)

func Pipe(p []int) (err error) {
	if len(p) != 2 {
		return EINVAL
	}
	var pp [2]_C_int
	err = pipe(&pp)
	if err == nil {
		p[0] = int(pp[0])
		p[1] = int(pp[1])
	}
	return
}

//sys	readlink(path string, buf []byte, bufSize uint64) (n int, err error)

func Readlink(path string, buf []byte) (n int, err error) {
	s := uint64(len(buf))
	return readlink(path, buf, s)
}

//sys	utimes(path string, times *[2]Timeval) (err error)

func Utimes(path string, tv []Timeval) error {
	if len(tv) != 2 {
		return EINVAL
	}
	return utimes(path, (*[2]Timeval)(unsafe.Pointer(&tv[0])))
}

//sys	utimensat(dirfd int, path string, times *[2]Timespec, flag int) (err error)

func UtimesNano(path string, ts []Timespec) error {
	if len(ts) != 2 {
		return EINVAL
	}
	return utimensat(_AT_FDCWD, path, (*[2]Timespec)(unsafe.Pointer(&ts[0])), 0)
}

//sys	Unlinkat(dirfd int, path string, flags int) (err error) = unlinkat

//sys	getcwd(buf *byte, size uint64) (err error)

const ImplementsGetwd = true

// SYS_EXECVE is unused on Hurd (exec goes through libc), but exec_unix.go
// references it in a dead branch, so it must exist.
const SYS_EXECVE = 0

func Getwd() (ret string, err error) {
	for len := uint64(4096); ; len *= 2 {
		b := make([]byte, len)
		err := getcwd(&b[0], len)
		if err == nil {
			i := 0
			for b[i] != 0 {
				i++
			}
			return string(b[:i]), nil
		}
		if err != ERANGE {
			return "", err
		}
	}
}

func Getcwd(buf []byte) (n int, err error) {
	err = getcwd(&buf[0], uint64(len(buf)))
	if err == nil {
		i := 0
		for buf[i] != 0 {
			i++
		}
		n = i
	}
	return
}

//sysnb	getgroups(ngid int, gid *_Gid_t) (n int, err error)
//sysnb	setgroups(ngid int, gid *_Gid_t) (err error)

func Getgroups() (gids []int, err error) {
	n, err := getgroups(0, nil)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}

	// Sanity check group count. Max is 16 on BSD.
	if n < 0 || n > 1000 {
		return nil, EINVAL
	}

	a := make([]_Gid_t, n)
	n, err = getgroups(n, &a[0])
	if err != nil {
		return nil, err
	}
	gids = make([]int, n)
	for i, v := range a[0:n] {
		gids[i] = int(v)
	}
	return
}

func Setgroups(gids []int) (err error) {
	if len(gids) == 0 {
		return setgroups(0, nil)
	}

	a := make([]_Gid_t, len(gids))
	for i, v := range gids {
		a[i] = _Gid_t(v)
	}
	return setgroups(len(a), &a[0])
}

func direntIno(buf []byte) (uint64, bool) {
	return readInt(buf, unsafe.Offsetof(Dirent{}.Ino), unsafe.Sizeof(Dirent{}.Ino))
}

func direntReclen(buf []byte) (uint64, bool) {
	return readInt(buf, unsafe.Offsetof(Dirent{}.Reclen), unsafe.Sizeof(Dirent{}.Reclen))
}

func direntNamlen(buf []byte) (uint64, bool) {
	return readInt(buf, unsafe.Offsetof(Dirent{}.Namlen), unsafe.Sizeof(Dirent{}.Namlen))
}

func Gettimeofday(tv *Timeval) (err error) {
	err = gettimeofday(tv, nil)
	return
}

// TODO
func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	return -1, ENOSYS
}

//sys	fdopendir(fd int) (dir uintptr, err error) = fdopendir
//sys	readdir64(dir uintptr) (entry uintptr, err error) = readdir64
//sys	closedir(dir uintptr) (err error) = closedir

// Hurd's glibc does not implement getdirentries(3). Instead it exposes the
// readdir(3) interface, which is backed by the directory RPC. ReadDirent
// adapts readdir to the getdirentries-style buffer interface used by the os
// package, maintaining one DIR stream per file descriptor.
type dirState struct {
	dir  uintptr
	pend []byte
}

var (
	dirMu sync.Mutex
	dirs  = map[int]*dirState{}
)

func ReadDirent(fd int, buf []byte) (n int, err error) {
	dirMu.Lock()
	defer dirMu.Unlock()

	st := dirs[fd]
	if st == nil {
		// Duplicate the descriptor so that closedir does not close the
		// descriptor owned by the caller.
		nfd, e := Dup(fd)
		if e != nil {
			return 0, e
		}
		d, e := fdopendir(nfd)
		if e != nil {
			Close(nfd)
			return 0, e
		}
		st = &dirState{dir: d}
		dirs[fd] = st
	}

	if len(st.pend) > 0 {
		n = copy(buf, st.pend)
		st.pend = st.pend[n:]
		if n == len(buf) {
			return n, nil
		}
	}

	for {
		entry, _ := readdir64(st.dir)
		if entry == 0 {
			return n, nil
		}
		d := (*Dirent)(unsafe.Pointer(entry))
		namlen := int(d.Namlen)
		if namlen > 255 {
			namlen = 255
		}
		reclen := 12 + namlen + 1
		rec := make([]byte, reclen)
		*(*uint64)(unsafe.Pointer(&rec[0])) = d.Ino
		*(*uint16)(unsafe.Pointer(&rec[8])) = uint16(reclen)
		rec[10] = d.Type
		rec[11] = uint8(namlen)
		copy(rec[12:12+namlen], unsafe.Slice((*byte)(unsafe.Pointer(&d.Name[0])), namlen))
		rec[12+namlen] = 0

		if n+len(rec) > len(buf) {
			st.pend = rec
			if n == 0 {
				return 0, EINVAL
			}
			return n, nil
		}
		copy(buf[n:], rec)
		n += len(rec)
	}
}

//sys  wait4(pid _Pid_t, status *_C_int, options int, rusage *Rusage) (wpid _Pid_t, err error)

func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, err error) {
	var status _C_int
	var r _Pid_t
	err = ERESTART
	// Hurd wait4 may return with ERESTART errno, while the process is still
	// active.
	for err == ERESTART {
		r, err = wait4(_Pid_t(pid), &status, options, rusage)
	}
	wpid = int(r)
	if wstatus != nil {
		*wstatus = WaitStatus(status)
	}
	return
}

//sys	fsync(fd int) (err error)

func Fsync(fd int) error {
	return fsync(fd)
}

/*
 * Socket
 */
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	Listen(s int, backlog int) (err error) = listen
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sys	socket(domain int, typ int, proto int) (fd int, err error)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	Shutdown(s int, how int) (err error) = shutdown
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)

func (sa *SockaddrInet4) sockaddr() (unsafe.Pointer, _Socklen, error) {
	if sa.Port < 0 || sa.Port > 0xFFFF {
		return nil, 0, EINVAL
	}
	sa.raw.Len = SizeofSockaddrInet4
	sa.raw.Family = AF_INET
	p := (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
	p[0] = byte(sa.Port >> 8)
	p[1] = byte(sa.Port)
	sa.raw.Addr = sa.Addr
	return unsafe.Pointer(&sa.raw), SizeofSockaddrInet4, nil
}

func (sa *SockaddrInet6) sockaddr() (unsafe.Pointer, _Socklen, error) {
	if sa.Port < 0 || sa.Port > 0xFFFF {
		return nil, 0, EINVAL
	}
	sa.raw.Len = SizeofSockaddrInet6
	sa.raw.Family = AF_INET6
	p := (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
	p[0] = byte(sa.Port >> 8)
	p[1] = byte(sa.Port)
	sa.raw.Scope_id = sa.ZoneId
	sa.raw.Addr = sa.Addr
	return unsafe.Pointer(&sa.raw), SizeofSockaddrInet6, nil
}

func (sa *RawSockaddrUnix) setLen(n int) {
	sa.Len = uint8(3 + n) // 2 for Family, Len; 1 for NUL.
}

func (sa *SockaddrUnix) sockaddr() (unsafe.Pointer, _Socklen, error) {
	name := sa.Name
	n := len(name)
	if n > len(sa.raw.Path) {
		return nil, 0, EINVAL
	}
	sa.raw.Family = AF_UNIX
	sa.raw.setLen(n)
	for i := 0; i < n; i++ {
		sa.raw.Path[i] = int8(name[i])
	}
	// length is family (uint16), name, NUL.
	sl := _Socklen(2)
	if n > 0 {
		sl += _Socklen(n) + 1
	}

	return unsafe.Pointer(&sa.raw), sl, nil
}

func Getsockname(fd int) (sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	if err = getsockname(fd, &rsa, &len); err != nil {
		return
	}
	return anyToSockaddr(&rsa)
}

//sys	accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error)

func Accept(fd int) (nfd int, sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	nfd, err = accept(fd, &rsa, &len)
	if err != nil {
		return
	}
	sa, err = anyToSockaddr(&rsa)
	if err != nil {
		Close(nfd)
		nfd = 0
	}
	return
}

func recvmsgRaw(fd int, p, oob []byte, flags int, rsa *RawSockaddrAny) (n, oobn int, recvflags int, err error) {
	var msg Msghdr
	msg.Name = (*byte)(unsafe.Pointer(rsa))
	msg.Namelen = uint32(SizeofSockaddrAny)
	var iov Iovec
	if len(p) > 0 {
		iov.Base = &p[0]
		iov.SetLen(len(p))
	}
	var dummy byte
	if len(oob) > 0 {
		var sockType int
		sockType, err = GetsockoptInt(fd, SOL_SOCKET, SO_TYPE)
		if err != nil {
			return
		}
		// receive at least one normal byte
		if sockType != SOCK_DGRAM && len(p) == 0 {
			iov.Base = &dummy
			iov.SetLen(1)
		}
		msg.Control = &oob[0]
		msg.SetControllen(len(oob))
	}
	msg.Iov = &iov
	msg.Iovlen = 1
	if n, err = recvmsg(fd, &msg, flags); err != nil {
		return
	}
	oobn = int(msg.Controllen)
	recvflags = int(msg.Flags)
	return
}

func sendmsgN(fd int, p, oob []byte, ptr unsafe.Pointer, salen _Socklen, flags int) (n int, err error) {
	var msg Msghdr
	msg.Name = (*byte)(ptr)
	msg.Namelen = uint32(salen)
	var iov Iovec
	if len(p) > 0 {
		iov.Base = &p[0]
		iov.SetLen(len(p))
	}
	var dummy byte
	if len(oob) > 0 {
		var sockType int
		sockType, err = GetsockoptInt(fd, SOL_SOCKET, SO_TYPE)
		if err != nil {
			return 0, err
		}
		// send at least one normal byte
		if sockType != SOCK_DGRAM && len(p) == 0 {
			iov.Base = &dummy
			iov.SetLen(1)
		}
		msg.Control = &oob[0]
		msg.SetControllen(len(oob))
	}
	msg.Iov = &iov
	msg.Iovlen = 1
	if n, err = sendmsg(fd, &msg, flags); err != nil {
		return 0, err
	}
	if len(oob) > 0 && len(p) == 0 {
		n = 0
	}
	return n, nil
}

func (sa *RawSockaddrUnix) getLen() (int, error) {
	if sa.Len < 2 || sa.Len > SizeofSockaddrUnix {
		return 0, EINVAL
	}
	n := int(sa.Len) - 2 // subtract leading Family, Len
	for i := 0; i < n; i++ {
		if sa.Path[i] == 0 {
			n = i
			break
		}
	}
	return n, nil
}

func (sa *RawSockaddrUnix) adjustAbstract(sl _Socklen) _Socklen {
	return sl
}

func anyToSockaddr(rsa *RawSockaddrAny) (Sockaddr, error) {
	switch rsa.Addr.Family {
	case AF_UNIX:
		pp := (*RawSockaddrUnix)(unsafe.Pointer(rsa))
		sa := new(SockaddrUnix)
		n, err := pp.getLen()
		if err != nil {
			return nil, err
		}
		sa.Name = string(unsafe.Slice((*byte)(unsafe.Pointer(&pp.Path[0])), n))
		return sa, nil

	case AF_INET:
		pp := (*RawSockaddrInet4)(unsafe.Pointer(rsa))
		sa := new(SockaddrInet4)
		p := (*[2]byte)(unsafe.Pointer(&pp.Port))
		sa.Port = int(p[0])<<8 + int(p[1])
		sa.Addr = pp.Addr
		return sa, nil

	case AF_INET6:
		pp := (*RawSockaddrInet6)(unsafe.Pointer(rsa))
		sa := new(SockaddrInet6)
		p := (*[2]byte)(unsafe.Pointer(&pp.Port))
		sa.Port = int(p[0])<<8 + int(p[1])
		sa.Addr = pp.Addr
		sa.ZoneId = pp.Scope_id
		return sa, nil
	}
	return nil, EAFNOSUPPORT
}

/*
 * Wait
 */

type WaitStatus uint32

func (w WaitStatus) Stopped() bool { return w&0x40 != 0 }
func (w WaitStatus) StopSignal() Signal {
	if !w.Stopped() {
		return -1
	}
	return Signal(w>>8) & 0xFF
}

func (w WaitStatus) Exited() bool { return w&0xFF == 0 }
func (w WaitStatus) ExitStatus() int {
	if !w.Exited() {
		return -1
	}
	return int((w >> 8) & 0xFF)
}

func (w WaitStatus) Signaled() bool { return w&0x40 == 0 && w&0xFF != 0 }
func (w WaitStatus) Signal() Signal {
	if !w.Signaled() {
		return -1
	}
	return Signal(w>>16) & 0xFF
}

func (w WaitStatus) Continued() bool { return w&0x01000000 != 0 }

func (w WaitStatus) CoreDump() bool { return w&0x80 == 0x80 }

func (w WaitStatus) TrapCause() int { return -1 }

/*
 * ptrace
 */

func raw_ptrace(request int, pid int, addr *byte, data *byte) Errno {
	return ENOSYS
}

/*
 * Direct access
 */

//sys	Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error) = openat64
//sys	Chdir(path string) (err error)
//sys	Chmod(path string, mode uint32) (err error)
//sys	Chown(path string, uid int, gid int) (err error)
//sys	Chroot(path string) (err error)
//sys	Close(fd int) (err error)
//sys	Dup(fd int) (nfd int, err error)
//sys	Faccessat(dirfd int, path string, mode uint32, flags int) (err error)
//sys	Fchdir(fd int) (err error)
//sys	Fchmod(fd int, mode uint32) (err error)
//sys	Fchmodat(dirfd int, path string, mode uint32, flags int) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error)
//sys	Fpathconf(fd int, name int) (val int, err error)
//sys	Fstat(fd int, stat *Stat_t) (err error) = fstat64
//sys	Fstatfs(fd int, buf *Statfs_t) (err error) = fstatfs64
//sys	Ftruncate(fd int, length int64) (err error) = ftruncate64
//sysnb	Getgid() (gid int)
//sysnb	Getpid() (pid int)
//sys	Geteuid() (euid int)
//sys	Getegid() (egid int)
//sys	Getppid() (ppid int)
//sys	Getpriority(which int, who int) (n int, err error)
//sysnb	Getrlimit(which int, lim *Rlimit) (err error) = getrlimit64
//sysnb	Getrusage(who int, rusage *Rusage) (err error)
//sysnb	Getuid() (uid int)
//sys	Kill(pid int, signum Signal) (err error)
//sys	Lchown(path string, uid int, gid int) (err error)
//sys	Link(path string, link string) (err error)
//sys	Lstat(path string, stat *Stat_t) (err error) = lstat64
//sys	Mkdir(path string, mode uint32) (err error)
//sys	Mkdirat(dirfd int, path string, mode uint32) (err error)
//sys	Mknodat(dirfd int, path string, mode uint32, dev int) (err error)
//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = fstatat64
//sys	Readlinkat(dirfd int, path string, buf []byte) (n int, err error) = readlinkat
//sys	Open(path string, mode int, perm uint32) (fd int, err error) = open64
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = pread64
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = pwrite64
//sys	read(fd int, p []byte) (n int, err error)
//sys	Rename(from string, to string) (err error)
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Rmdir(path string) (err error)
//sys	Seek(fd int, offset int64, whence int) (newoffset int64, err error) = lseek64
//sysnb	Setegid(egid int) (err error)
//sysnb	Seteuid(euid int) (err error)
//sysnb	Setgid(gid int) (err error)
//sysnb	Setuid(uid int) (err error)
//sysnb	Setpgid(pid int, pgid int) (err error)
//sys	Setpriority(which int, who int, prio int) (err error)
//sysnb	Setregid(rgid int, egid int) (err error)
//sysnb	Setreuid(ruid int, euid int) (err error)
//sysnb	setrlimit(which int, lim *Rlimit) (err error) = setrlimit64
//sys	Stat(path string, stat *Stat_t) (err error) = stat64
//sys	Statfs(path string, buf *Statfs_t) (err error) = statfs64
//sys	Symlink(path string, link string) (err error)
//sys	Truncate(path string, length int64) (err error) = truncate64
//sys	Umask(newmask int) (oldmask int)
//sys	Unlink(path string) (err error)
//sysnb	Uname(buf *Utsname) (err error)
//sys	write(fd int, p []byte) (n int, err error)
//sys	writev(fd int, iovecs []Iovec) (n uintptr, err error)

//sys	gettimeofday(tv *Timeval, tzp *Timezone) (err error)

func setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}

func setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: int32(sec), Usec: int32(usec)}
}

func readlen(fd int, buf *byte, nbuf int) (n int, err error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_read)), 3, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

/*
 * Map
 */

var mapper = &mmapper{
	active: make(map[*byte][]byte),
	mmap:   mmap,
	munmap: munmap,
}

// mmap32 is the libc mmap with a 32-bit file offset. This is enough for
// practically all uses; larger offsets return EINVAL.
//sys	mmap32(addr uintptr, length uintptr, prot int, flag int, fd int, pos uintptr) (ret uintptr, err error) = mmap
//sys	munmap(addr uintptr, length uintptr) (err error)

func mmap(addr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error) {
	if offset != int64(int32(offset)) {
		return 0, EINVAL
	}
	r, err := mmap32(addr, length, prot, flags, fd, uintptr(offset))
	if r == ^uintptr(0) {
		return 0, err
	}
	return r, nil
}

func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error) {
	return mapper.Mmap(fd, offset, length, prot, flags)
}

func Munmap(b []byte) (err error) {
	return mapper.Munmap(b)
}

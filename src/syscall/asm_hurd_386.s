// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd && 386

#include "textflag.h"

// System calls for hurd/386 are implemented in ../runtime/os2_hurd.go.

TEXT ·syscall6(SB),NOSPLIT,$0
	JMP	runtime·syscall_syscall6(SB)

TEXT ·rawSyscall6(SB),NOSPLIT,$0
	JMP	runtime·syscall_rawSyscall6(SB)

// Data slots holding resolved libc function addresses. The dynamic
// R_386_32 relocations live here (in .data), not in .text.
DATA ·libc_Chdir+0(SB)/4, $imp_libc_Chdir(SB)
GLOBL ·libc_Chdir(SB), NOPTR, $4
DATA ·libc_Chmod+0(SB)/4, $imp_libc_Chmod(SB)
GLOBL ·libc_Chmod(SB), NOPTR, $4
DATA ·libc_Chown+0(SB)/4, $imp_libc_Chown(SB)
GLOBL ·libc_Chown(SB), NOPTR, $4
DATA ·libc_Chroot+0(SB)/4, $imp_libc_Chroot(SB)
GLOBL ·libc_Chroot(SB), NOPTR, $4
DATA ·libc_Close+0(SB)/4, $imp_libc_Close(SB)
GLOBL ·libc_Close(SB), NOPTR, $4
DATA ·libc_Dup+0(SB)/4, $imp_libc_Dup(SB)
GLOBL ·libc_Dup(SB), NOPTR, $4
DATA ·libc_Dup2+0(SB)/4, $imp_libc_Dup2(SB)
GLOBL ·libc_Dup2(SB), NOPTR, $4
DATA ·libc_Faccessat+0(SB)/4, $imp_libc_Faccessat(SB)
GLOBL ·libc_Faccessat(SB), NOPTR, $4
DATA ·libc_Fchdir+0(SB)/4, $imp_libc_Fchdir(SB)
GLOBL ·libc_Fchdir(SB), NOPTR, $4
DATA ·libc_Fchmod+0(SB)/4, $imp_libc_Fchmod(SB)
GLOBL ·libc_Fchmod(SB), NOPTR, $4
DATA ·libc_Fchmodat+0(SB)/4, $imp_libc_Fchmodat(SB)
GLOBL ·libc_Fchmodat(SB), NOPTR, $4
DATA ·libc_Fchown+0(SB)/4, $imp_libc_Fchown(SB)
GLOBL ·libc_Fchown(SB), NOPTR, $4
DATA ·libc_Fchownat+0(SB)/4, $imp_libc_Fchownat(SB)
GLOBL ·libc_Fchownat(SB), NOPTR, $4
DATA ·libc_Fpathconf+0(SB)/4, $imp_libc_Fpathconf(SB)
GLOBL ·libc_Fpathconf(SB), NOPTR, $4
DATA ·libc_Getegid+0(SB)/4, $imp_libc_Getegid(SB)
GLOBL ·libc_Getegid(SB), NOPTR, $4
DATA ·libc_Geteuid+0(SB)/4, $imp_libc_Geteuid(SB)
GLOBL ·libc_Geteuid(SB), NOPTR, $4
DATA ·libc_Getgid+0(SB)/4, $imp_libc_Getgid(SB)
GLOBL ·libc_Getgid(SB), NOPTR, $4
DATA ·libc_Getpid+0(SB)/4, $imp_libc_Getpid(SB)
GLOBL ·libc_Getpid(SB), NOPTR, $4
DATA ·libc_Getppid+0(SB)/4, $imp_libc_Getppid(SB)
GLOBL ·libc_Getppid(SB), NOPTR, $4
DATA ·libc_Getpriority+0(SB)/4, $imp_libc_Getpriority(SB)
GLOBL ·libc_Getpriority(SB), NOPTR, $4
DATA ·libc_Getrusage+0(SB)/4, $imp_libc_Getrusage(SB)
GLOBL ·libc_Getrusage(SB), NOPTR, $4
DATA ·libc_Getuid+0(SB)/4, $imp_libc_Getuid(SB)
GLOBL ·libc_Getuid(SB), NOPTR, $4
DATA ·libc_Kill+0(SB)/4, $imp_libc_Kill(SB)
GLOBL ·libc_Kill(SB), NOPTR, $4
DATA ·libc_Lchown+0(SB)/4, $imp_libc_Lchown(SB)
GLOBL ·libc_Lchown(SB), NOPTR, $4
DATA ·libc_Link+0(SB)/4, $imp_libc_Link(SB)
GLOBL ·libc_Link(SB), NOPTR, $4
DATA ·libc_Mkdir+0(SB)/4, $imp_libc_Mkdir(SB)
GLOBL ·libc_Mkdir(SB), NOPTR, $4
DATA ·libc_Mkdirat+0(SB)/4, $imp_libc_Mkdirat(SB)
GLOBL ·libc_Mkdirat(SB), NOPTR, $4
DATA ·libc_Mknodat+0(SB)/4, $imp_libc_Mknodat(SB)
GLOBL ·libc_Mknodat(SB), NOPTR, $4
DATA ·libc_Rename+0(SB)/4, $imp_libc_Rename(SB)
GLOBL ·libc_Rename(SB), NOPTR, $4
DATA ·libc_Renameat+0(SB)/4, $imp_libc_Renameat(SB)
GLOBL ·libc_Renameat(SB), NOPTR, $4
DATA ·libc_Rmdir+0(SB)/4, $imp_libc_Rmdir(SB)
GLOBL ·libc_Rmdir(SB), NOPTR, $4
DATA ·libc_Setegid+0(SB)/4, $imp_libc_Setegid(SB)
GLOBL ·libc_Setegid(SB), NOPTR, $4
DATA ·libc_Seteuid+0(SB)/4, $imp_libc_Seteuid(SB)
GLOBL ·libc_Seteuid(SB), NOPTR, $4
DATA ·libc_Setgid+0(SB)/4, $imp_libc_Setgid(SB)
GLOBL ·libc_Setgid(SB), NOPTR, $4
DATA ·libc_Setpgid+0(SB)/4, $imp_libc_Setpgid(SB)
GLOBL ·libc_Setpgid(SB), NOPTR, $4
DATA ·libc_Setpriority+0(SB)/4, $imp_libc_Setpriority(SB)
GLOBL ·libc_Setpriority(SB), NOPTR, $4
DATA ·libc_Setregid+0(SB)/4, $imp_libc_Setregid(SB)
GLOBL ·libc_Setregid(SB), NOPTR, $4
DATA ·libc_Setreuid+0(SB)/4, $imp_libc_Setreuid(SB)
GLOBL ·libc_Setreuid(SB), NOPTR, $4
DATA ·libc_Setuid+0(SB)/4, $imp_libc_Setuid(SB)
GLOBL ·libc_Setuid(SB), NOPTR, $4
DATA ·libc_Symlink+0(SB)/4, $imp_libc_Symlink(SB)
GLOBL ·libc_Symlink(SB), NOPTR, $4
DATA ·libc_Umask+0(SB)/4, $imp_libc_Umask(SB)
GLOBL ·libc_Umask(SB), NOPTR, $4
DATA ·libc_Uname+0(SB)/4, $imp_libc_Uname(SB)
GLOBL ·libc_Uname(SB), NOPTR, $4
DATA ·libc_Unlink+0(SB)/4, $imp_libc_Unlink(SB)
GLOBL ·libc_Unlink(SB), NOPTR, $4
DATA ·libc_accept+0(SB)/4, $imp_libc_accept(SB)
GLOBL ·libc_accept(SB), NOPTR, $4
DATA ·libc_bind+0(SB)/4, $imp_libc_bind(SB)
GLOBL ·libc_bind(SB), NOPTR, $4
DATA ·libc_closedir+0(SB)/4, $imp_libc_closedir(SB)
GLOBL ·libc_closedir(SB), NOPTR, $4
DATA ·libc_connect+0(SB)/4, $imp_libc_connect(SB)
GLOBL ·libc_connect(SB), NOPTR, $4
DATA ·libc_fcntl+0(SB)/4, $imp_libc_fcntl(SB)
GLOBL ·libc_fcntl(SB), NOPTR, $4
DATA ·libc_fdopendir+0(SB)/4, $imp_libc_fdopendir(SB)
GLOBL ·libc_fdopendir(SB), NOPTR, $4
DATA ·libc_fstat64+0(SB)/4, $imp_libc_fstat64(SB)
GLOBL ·libc_fstat64(SB), NOPTR, $4
DATA ·libc_fstatat64+0(SB)/4, $imp_libc_fstatat64(SB)
GLOBL ·libc_fstatat64(SB), NOPTR, $4
DATA ·libc_fstatfs64+0(SB)/4, $imp_libc_fstatfs64(SB)
GLOBL ·libc_fstatfs64(SB), NOPTR, $4
DATA ·libc_fsync+0(SB)/4, $imp_libc_fsync(SB)
GLOBL ·libc_fsync(SB), NOPTR, $4
DATA ·libc_ftruncate64+0(SB)/4, $imp_libc_ftruncate64(SB)
GLOBL ·libc_ftruncate64(SB), NOPTR, $4
DATA ·libc_getcwd+0(SB)/4, $imp_libc_getcwd(SB)
GLOBL ·libc_getcwd(SB), NOPTR, $4
DATA ·libc_getgroups+0(SB)/4, $imp_libc_getgroups(SB)
GLOBL ·libc_getgroups(SB), NOPTR, $4
DATA ·libc_getpeername+0(SB)/4, $imp_libc_getpeername(SB)
GLOBL ·libc_getpeername(SB), NOPTR, $4
DATA ·libc_getrlimit64+0(SB)/4, $imp_libc_getrlimit64(SB)
GLOBL ·libc_getrlimit64(SB), NOPTR, $4
DATA ·libc_getsid+0(SB)/4, $imp_libc_getsid(SB)
GLOBL ·libc_getsid(SB), NOPTR, $4
DATA ·libc_getsockname+0(SB)/4, $imp_libc_getsockname(SB)
GLOBL ·libc_getsockname(SB), NOPTR, $4
DATA ·libc_getsockopt+0(SB)/4, $imp_libc_getsockopt(SB)
GLOBL ·libc_getsockopt(SB), NOPTR, $4
DATA ·libc_gettimeofday+0(SB)/4, $imp_libc_gettimeofday(SB)
GLOBL ·libc_gettimeofday(SB), NOPTR, $4
DATA ·libc_listen+0(SB)/4, $imp_libc_listen(SB)
GLOBL ·libc_listen(SB), NOPTR, $4
DATA ·libc_lseek64+0(SB)/4, $imp_libc_lseek64(SB)
GLOBL ·libc_lseek64(SB), NOPTR, $4
DATA ·libc_lstat64+0(SB)/4, $imp_libc_lstat64(SB)
GLOBL ·libc_lstat64(SB), NOPTR, $4
DATA ·libc_mmap+0(SB)/4, $imp_libc_mmap(SB)
GLOBL ·libc_mmap(SB), NOPTR, $4
DATA ·libc_munmap+0(SB)/4, $imp_libc_munmap(SB)
GLOBL ·libc_munmap(SB), NOPTR, $4
DATA ·libc_open64+0(SB)/4, $imp_libc_open64(SB)
GLOBL ·libc_open64(SB), NOPTR, $4
DATA ·libc_openat64+0(SB)/4, $imp_libc_openat64(SB)
GLOBL ·libc_openat64(SB), NOPTR, $4
DATA ·libc_pipe+0(SB)/4, $imp_libc_pipe(SB)
GLOBL ·libc_pipe(SB), NOPTR, $4
DATA ·libc_pread64+0(SB)/4, $imp_libc_pread64(SB)
GLOBL ·libc_pread64(SB), NOPTR, $4
DATA ·libc_pwrite64+0(SB)/4, $imp_libc_pwrite64(SB)
GLOBL ·libc_pwrite64(SB), NOPTR, $4
DATA ·libc_read+0(SB)/4, $imp_libc_read(SB)
GLOBL ·libc_read(SB), NOPTR, $4
DATA ·libc_readdir64+0(SB)/4, $imp_libc_readdir64(SB)
GLOBL ·libc_readdir64(SB), NOPTR, $4
DATA ·libc_readlink+0(SB)/4, $imp_libc_readlink(SB)
GLOBL ·libc_readlink(SB), NOPTR, $4
DATA ·libc_readlinkat+0(SB)/4, $imp_libc_readlinkat(SB)
GLOBL ·libc_readlinkat(SB), NOPTR, $4
DATA ·libc_recvfrom+0(SB)/4, $imp_libc_recvfrom(SB)
GLOBL ·libc_recvfrom(SB), NOPTR, $4
DATA ·libc_recvmsg+0(SB)/4, $imp_libc_recvmsg(SB)
GLOBL ·libc_recvmsg(SB), NOPTR, $4
DATA ·libc_sendmsg+0(SB)/4, $imp_libc_sendmsg(SB)
GLOBL ·libc_sendmsg(SB), NOPTR, $4
DATA ·libc_sendto+0(SB)/4, $imp_libc_sendto(SB)
GLOBL ·libc_sendto(SB), NOPTR, $4
DATA ·libc_setgroups+0(SB)/4, $imp_libc_setgroups(SB)
GLOBL ·libc_setgroups(SB), NOPTR, $4
DATA ·libc_setrlimit64+0(SB)/4, $imp_libc_setrlimit64(SB)
GLOBL ·libc_setrlimit64(SB), NOPTR, $4
DATA ·libc_setsockopt+0(SB)/4, $imp_libc_setsockopt(SB)
GLOBL ·libc_setsockopt(SB), NOPTR, $4
DATA ·libc_shutdown+0(SB)/4, $imp_libc_shutdown(SB)
GLOBL ·libc_shutdown(SB), NOPTR, $4
DATA ·libc_socket+0(SB)/4, $imp_libc_socket(SB)
GLOBL ·libc_socket(SB), NOPTR, $4
DATA ·libc_socketpair+0(SB)/4, $imp_libc_socketpair(SB)
GLOBL ·libc_socketpair(SB), NOPTR, $4
DATA ·libc_stat64+0(SB)/4, $imp_libc_stat64(SB)
GLOBL ·libc_stat64(SB), NOPTR, $4
DATA ·libc_statfs64+0(SB)/4, $imp_libc_statfs64(SB)
GLOBL ·libc_statfs64(SB), NOPTR, $4
DATA ·libc_truncate64+0(SB)/4, $imp_libc_truncate64(SB)
GLOBL ·libc_truncate64(SB), NOPTR, $4
DATA ·libc_unlinkat+0(SB)/4, $imp_libc_unlinkat(SB)
GLOBL ·libc_unlinkat(SB), NOPTR, $4
DATA ·libc_utimensat+0(SB)/4, $imp_libc_utimensat(SB)
GLOBL ·libc_utimensat(SB), NOPTR, $4
DATA ·libc_utimes+0(SB)/4, $imp_libc_utimes(SB)
GLOBL ·libc_utimes(SB), NOPTR, $4
DATA ·libc_wait4+0(SB)/4, $imp_libc_wait4(SB)
GLOBL ·libc_wait4(SB), NOPTR, $4
DATA ·libc_write+0(SB)/4, $imp_libc_write(SB)
GLOBL ·libc_write(SB), NOPTR, $4
DATA ·libc_writev+0(SB)/4, $imp_libc_writev(SB)
GLOBL ·libc_writev(SB), NOPTR, $4

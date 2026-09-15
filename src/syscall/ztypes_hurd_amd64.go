// Code generated for GNU/Hurd (amd64) from glibc 2.43 headers. DO NOT EDIT.

//go:build hurd && amd64

package syscall

const (
	sizeofPtr      = 0x8
	sizeofShort    = 0x2
	sizeofInt      = 0x4
	sizeofLong     = 0x8
	sizeofLongLong = 0x8
	PathMax        = 0x1000
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int64
	_C_long_long int64
)

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec  int64
	Usec int64
}

type Timezone struct {
	Minuteswest int32
	Dsttime     int32
}

type Tms struct {
	Utime  int64
	Stime  int64
	Cutime int64
	Cstime int64
}

type Utimbuf struct {
	Actime  int64
	Modtime int64
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type _Gid_t uint32

type _Pid_t int32

type _Uid_t uint32

type Stat_t struct {
	X__fstype int32
	Dev       uint64
	Ino       uint64
	X__gen    uint32
	Rdev      uint64
	Mode      uint32
	Nlink     uint64
	Uid       uint32
	Gid       uint32
	Size      int64
	Atim      Timespec
	Mtim      Timespec
	Ctim      Timespec
	Blksize   int64
	Blocks    int64
	X__author uint32
	X__flags  uint32
	X__spare  [8]int32
}

type Fsid struct {
	X__val [2]int32
}

type Statfs_t struct {
	Type    uint32
	Bsize   uint64
	Blocks  uint64
	Bfree   uint64
	Bavail  uint64
	Files   uint64
	Ffree   uint64
	Fsid    Fsid
	Namelen uint64
	Favail  uint64
	Frsize  uint64
	Flags   uint64
	Spare   [3]uint32
	X__pad  [4]byte
}

type Dirent struct {
	Ino    uint64
	Reclen uint16
	Type   uint8
	Namlen uint8
	Name   [256]int8
}

type Flock_t struct {
	Type   int16
	Whence int16
	Start  int64
	Len    int64
	Pid    int32
}

type RawSockaddrInet4 struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type RawSockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Len    uint8
	Family uint8
	Path   [108]int8
}

type RawSockaddr struct {
	Len    uint8
	Family uint8
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [112]int8
}

type _Socklen uint32

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint64
}

type IPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type IPMreqn struct {
	Multiaddr [4]byte /* in_addr */
	Address   [4]byte /* in_addr */
	Ifindex   int32
}

type IPv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Interface uint32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     int32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type Inet4Pktinfo struct {
	Ifindex  int32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
}

type Inet6Pktinfo struct {
	Addr    [16]byte /* in6_addr */
	Ifindex uint32
}

type IPv6MTUInfo struct {
	Addr RawSockaddrInet6
	Mtu  uint32
}

type ICMPv6Filter struct {
	Data [8]uint32
}

type Ucred struct {
	Pid int32
	Uid uint32
	Gid uint32
}

type FdSet struct {
	Bits [16]int64
}

type Utsname struct {
	Sysname  [1024]int8
	Nodename [1024]int8
	Release  [1024]int8
	Version  [1024]int8
	Machine  [1024]int8
}

type Ifaddrs struct {
	Next    *Ifaddrs
	Name    *byte
	Flags   uint32
	Addr    *RawSockaddr
	Netmask *RawSockaddr
	Ifu     *RawSockaddr
	Data    *byte
}

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed uint32
	Ospeed uint32
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

const (
	SizeofSockaddrInet4 = 0x10
	SizeofSockaddrInet6 = 0x1c
	SizeofSockaddrAny   = 0x80
	SizeofSockaddrUnix  = 0x6e
	SizeofLinger        = 0x8
	SizeofIPMreq        = 0x8
	SizeofIPMreqn       = 0xc
	SizeofIPv6Mreq      = 0x14
	SizeofMsghdr        = 0x30
	SizeofCmsghdr       = 0xc
	SizeofInet4Pktinfo  = 0xc
	SizeofInet6Pktinfo  = 0x14
	SizeofIPv6MTUInfo   = 0x20
	SizeofICMPv6Filter  = 0x20
	SizeofUcred         = 0xc
)

// pollFd is used by internal/poll on Unix.
type pollFd struct {
	Fd      int32
	Events  int16
	Revents int16
}

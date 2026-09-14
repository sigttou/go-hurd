// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package runtime

import (
	"internal/runtime/atomic"
	"unsafe"
)

// Based on the AIX poll(2) netpoller.

//go:nosplit
func poll(pfds *pollfd, npfds uintptr, timeout uintptr) (int32, int32) {
	r, err := syscall3(&libc_poll, uintptr(unsafe.Pointer(pfds)), npfds, timeout)
	return int32(r), int32(err)
}

type pollfd struct {
	fd      int32
	events  int16
	revents int16
}

const _POLLIN = 0x0001
const _POLLOUT = 0x0004
const _POLLHUP = 0x0010
const _POLLERR = 0x0008

var (
	pfds           []pollfd
	pds            []*pollDesc
	mtxpoll        mutex
	mtxset         mutex
	rdwake         int32
	wrwake         int32
	pendingUpdates int32

	netpollWakeSig atomic.Uint32 // used to avoid duplicate calls of netpollBreak
)

func netpollinit() {
	r, w, errno := nonblockingPipe()
	if errno != 0 {
		throw("netpollinit: failed to create pipe")
	}
	rdwake = r
	wrwake = w

	pfds = make([]pollfd, 1, 128)
	pfds[0].fd = rdwake
	pfds[0].events = _POLLIN

	pds = make([]*pollDesc, 1, 128)
	pds[0] = nil
}

func netpollIsPollDescriptor(fd uintptr) bool {
	return fd == uintptr(rdwake) || fd == uintptr(wrwake)
}

func netpollwakeup() {
	if pendingUpdates == 0 {
		pendingUpdates = 1
		b := [1]byte{0}
		write(uintptr(wrwake), unsafe.Pointer(&b[0]), 1)
	}
}

func netpollopen(fd uintptr, pd *pollDesc) int32 {
	lock(&mtxpoll)
	netpollwakeup()

	lock(&mtxset)
	unlock(&mtxpoll)

	pd.user = uint32(len(pfds))
	pfds = append(pfds, pollfd{fd: int32(fd)})
	pds = append(pds, pd)
	unlock(&mtxset)
	return 0
}

func netpollclose(fd uintptr) int32 {
	lock(&mtxpoll)
	netpollwakeup()

	lock(&mtxset)
	unlock(&mtxpoll)

	for i := 0; i < len(pfds); i++ {
		if pfds[i].fd == int32(fd) {
			pfds[i] = pfds[len(pfds)-1]
			pfds = pfds[:len(pfds)-1]

			pds[i] = pds[len(pds)-1]
			pds[i].user = uint32(i)
			pds = pds[:len(pds)-1]
			break
		}
	}
	unlock(&mtxset)
	return 0
}

func netpollarm(pd *pollDesc, mode int) {
	lock(&mtxpoll)
	netpollwakeup()

	lock(&mtxset)
	unlock(&mtxpoll)

	switch mode {
	case 'r':
		pfds[pd.user].events |= _POLLIN
	case 'w':
		pfds[pd.user].events |= _POLLOUT
	}
	unlock(&mtxset)
}

// netpollBreak interrupts a poll.
func netpollBreak() {
	if !netpollWakeSig.CompareAndSwap(0, 1) {
		return
	}

	b := [1]byte{0}
	write(uintptr(wrwake), unsafe.Pointer(&b[0]), 1)
}

//go:nowritebarrierrec
func netpoll(delay int64) (gList, int32) {
	var timeout uintptr
	if delay < 0 {
		timeout = ^uintptr(0)
	} else if delay == 0 {
		return gList{}, 0
	} else if delay < 1e6 {
		timeout = 1
	} else if delay < 1e15 {
		timeout = uintptr(delay / 1e6)
	} else {
		timeout = 1e9
	}
retry:
	lock(&mtxpoll)
	lock(&mtxset)
	pendingUpdates = 0
	unlock(&mtxpoll)

	n, e := poll(&pfds[0], uintptr(len(pfds)), timeout)
	if n < 0 {
		if e != _EINTR {
			println("errno=", e, " len(pfds)=", len(pfds))
			throw("poll failed")
		}
		unlock(&mtxset)
		if timeout > 0 {
			return gList{}, 0
		}
		goto retry
	}
	if n != 0 && pfds[0].revents&(_POLLIN|_POLLHUP|_POLLERR) != 0 {
		if delay != 0 {
			var b [1]byte
			for read(rdwake, unsafe.Pointer(&b[0]), 1) == 1 {
			}
			netpollWakeSig.Store(0)
		}
		n--
	}
	var toRun gList
	delta := int32(0)
	for i := 1; i < len(pfds) && n > 0; i++ {
		pfd := &pfds[i]

		var mode int32
		if pfd.revents&(_POLLIN|_POLLHUP|_POLLERR) != 0 {
			mode += 'r'
			pfd.events &= ^_POLLIN
		}
		if pfd.revents&(_POLLOUT|_POLLHUP|_POLLERR) != 0 {
			mode += 'w'
			pfd.events &= ^_POLLOUT
		}
		if mode != 0 {
			pds[i].setEventErr(pfd.revents == _POLLERR, 0)
			delta += netpollready(&toRun, pds[i], mode)
			n--
		}
	}
	unlock(&mtxset)
	return toRun, delta
}

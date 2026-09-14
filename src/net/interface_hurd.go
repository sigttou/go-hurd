// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd

package net

import (
	"syscall"
	"unsafe"
)

// goString converts a NUL-terminated C string to a Go string.
func goString(p *byte) string {
	if p == nil {
		return ""
	}
	var n int
	for *(*byte)(unsafe.Add(unsafe.Pointer(p), n)) != 0 {
		n++
	}
	return string(unsafe.Slice(p, n))
}

func linkFlags(rawFlags int32) Flags {
	var f Flags
	if rawFlags&syscall.IFF_UP != 0 {
		f |= FlagUp
	}
	if rawFlags&syscall.IFF_RUNNING != 0 {
		f |= FlagRunning
	}
	if rawFlags&syscall.IFF_BROADCAST != 0 {
		f |= FlagBroadcast
	}
	if rawFlags&syscall.IFF_LOOPBACK != 0 {
		f |= FlagLoopback
	}
	if rawFlags&syscall.IFF_POINTOPOINT != 0 {
		f |= FlagPointToPoint
	}
	if rawFlags&syscall.IFF_MULTICAST != 0 {
		f |= FlagMulticast
	}
	return f
}

// If the ifindex is zero, interfaceTable returns mappings of all
// network interfaces. Otherwise it returns a mapping of a specific
// interface.
func interfaceTable(ifindex int) ([]Interface, error) {
	var ifap *syscall.Ifaddrs
	if err := syscall.Getifaddrs(&ifap); err != nil {
		return nil, err
	}
	defer syscall.Freeifaddrs(ifap)

	var ift []Interface
	seen := make(map[string]bool)
	for ifa := ifap; ifa != nil; ifa = ifa.Next {
		name := goString(ifa.Name)
		if name == "" || seen[name] {
			continue
		}
		seen[name] = true
		idx := int(syscall.Ifnametoindex(ifa.Name))
		if ifindex != 0 && ifindex != idx {
			continue
		}
		mtu := 1500
		if ifa.Flags&syscall.IFF_LOOPBACK != 0 {
			mtu = 65536
		}
		ift = append(ift, Interface{
			Index: idx,
			MTU:   mtu,
			Name:  name,
			Flags: linkFlags(int32(ifa.Flags)),
		})
		if ifindex != 0 {
			break
		}
	}
	return ift, nil
}

// If the ifi is nil, interfaceAddrTable returns addresses for all
// network interfaces. Otherwise it returns addresses for a specific
// interface.
func interfaceAddrTable(ifi *Interface) ([]Addr, error) {
	var ifap *syscall.Ifaddrs
	if err := syscall.Getifaddrs(&ifap); err != nil {
		return nil, err
	}
	defer syscall.Freeifaddrs(ifap)

	var ift []Addr
	for ifa := ifap; ifa != nil; ifa = ifa.Next {
		name := goString(ifa.Name)
		if ifi != nil && ifi.Name != name {
			continue
		}
		if ifa.Addr == nil || ifa.Netmask == nil {
			continue
		}
		switch ifa.Addr.Family {
		case syscall.AF_INET:
			sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(ifa.Addr))
			nm := (*syscall.RawSockaddrInet4)(unsafe.Pointer(ifa.Netmask))
			ip := IPv4(sa.Addr[0], sa.Addr[1], sa.Addr[2], sa.Addr[3])
			mask := IPMask{nm.Addr[0], nm.Addr[1], nm.Addr[2], nm.Addr[3]}
			ift = append(ift, &IPNet{IP: ip, Mask: mask})
		case syscall.AF_INET6:
			sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(ifa.Addr))
			nm := (*syscall.RawSockaddrInet6)(unsafe.Pointer(ifa.Netmask))
			ip := make(IP, IPv6len)
			copy(ip, sa.Addr[:])
			mask := make(IPMask, IPv6len)
			copy(mask, nm.Addr[:])
			ift = append(ift, &IPNet{IP: ip, Mask: mask})
		}
	}
	return ift, nil
}

// interfaceMulticastAddrTable returns addresses for a specific
// interface.
func interfaceMulticastAddrTable(ifi *Interface) ([]Addr, error) {
	return nil, nil
}

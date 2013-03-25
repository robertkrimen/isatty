// +build linux,!appengine darwin freebsd

package isatty

import (
	"syscall"
	"unsafe"
)

func check(fd uintptr) bool {
	var garbage [128]byte
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd),
		ioctlQuery,                        // Either TCGETS or TIOCGETA, basically request the termios struct
		uintptr(unsafe.Pointer(&garbage)), // A garbage slice to be filled with termios data
		0, 0, 0)
	return err == 0
}

func canCheck() bool {
	return false
}

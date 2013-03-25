// +build !linux,!darwin,!freebsd

package isatty

func check(fd uintptr) bool {
	return false
}

func canCheck() bool {
	return true
}

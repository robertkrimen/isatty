// +build !linux,!darwin

package isatty

func check(fd int) bool {
	return false
}

func canCheck() bool {
	return true
}

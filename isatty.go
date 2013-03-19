// Package isatty uses an ioctl request to determine if a file descriptor is a terminal.
//
// Originally adapted from: https://code.google.com/p/go/source/browse/ssh/terminal/util.go?repo=crypto
package isatty

// Check will return true if the file descriptor is a terminal and false otherwise.
//
// If Check is unable to determine the terminal status of a file descriptor, then it returns false. This could
// be the case if this package is used on Windows, which lacks the necessary ioctl access.
func Check(fd uintptr) bool {
	return check(fd)
}

// CanCheck will return true if Check is capable of returning a genuine answer.
//
//      Darwin (Mac OS X)   CanCheck() => true   # syscall.Syscall6(..., TIOCGETA, ...)
//      Linux               CanCheck() => true   # syscall.Syscall6(..., TCGETS, ...)
//      Windows             CanCheck() => false  # N/A
//
func CanCheck() bool {
	return canCheck()
}

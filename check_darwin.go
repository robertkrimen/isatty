// +build darwin

package isatty

import "syscall"

const ioctlQuery = syscall.TIOCGETA

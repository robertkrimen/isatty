// +build darwin freebsd

package isatty

import "syscall"

const ioctlQuery = syscall.TIOCGETA

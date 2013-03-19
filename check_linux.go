// +build linux

package isatty

import "syscall"

const ioctlQuery = syscall.TCGETS

# isatty
--
    import "github.com/robertkrimen/isatty"

Package isatty uses an ioctl request to determine if a file descriptor is a terminal.

Originally adapted from: https://code.google.com/p/go/source/browse/ssh/terminal/util.go?repo=crypto

## Usage

#### func  CanCheck

```go
func CanCheck() bool
```
CanCheck will return true if Check is capable of returning a genuine answer.

    Darwin (Mac OS X)   CanCheck() => true   # syscall.Syscall6(..., TIOCGETA, ...)
    Linux               CanCheck() => true   # syscall.Syscall6(..., TCGETS, ...)
    Windows             CanCheck() => false  # N/A

#### func  Check

```go
func Check(fd uintptr) bool
```
Check will return true if the file descriptor is a terminal and false otherwise.

If Check is unable to determine the terminal status of a file descriptor, then
it returns false. This could be the case if this package is used on Windows,
which lacks the necessary ioctl access.

--
**godocdown** http://github.com/robertkrimen/godocdown

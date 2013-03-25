# isatty
--
    import "github.com/robertkrimen/isatty"

Package isatty tries to determine if a file descriptor is connected to a terminal.

     import (
         "github.com/robertkrimen/isatty"
         "os"
     )

     func isTerminalStdin() bool {
         return isatty.Check(os.Stdin.Fd())
     }

Originally adapted from: https://code.google.com/p/go/source/browse/ssh/terminal/util.go?repo=crypto

## Usage

#### func  CanCheck

```go
func CanCheck() bool
```
CanCheck will return true if Check is capable of returning a genuine answer.

    Darwin (Mac OS X)   CanCheck() => true   # syscall.Syscall6(..., TIOCGETA, ...)
    Linux               CanCheck() => true   # syscall.Syscall6(..., TCGETS, ...)
    FreeBSD             CanCheck() => true   # syscall.Syscall6(..., TIOCGETA, ...)
    Windows             CanCheck() => false  # N/A
    (Otherwise)         CanCheck() => false  # N/A

#### func  Check

```go
func Check(fd uintptr) bool
```
Check will return true if the file descriptor is a terminal and false otherwise.

If Check is unable to determine the terminal status of a file descriptor, then
it returns false. This could be the case if this package is used on Windows or
an unknown/unconsidered platform.

--
**godocdown** http://github.com/robertkrimen/godocdown

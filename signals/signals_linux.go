//go:build linux || darwin || freebsd || netbsd || openbsd || solaris || macos

package signals

import "syscall"

const (
	SIGUSR1 = syscall.SIGUSR1
	SIGUSR2 = syscall.SIGUSR2
	SIGTSTP = syscall.SIGTSTP
)

var Kill func(pid int, sig syscall.Signal) error = syscall.Kill

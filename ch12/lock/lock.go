package main

// Windows用Locking files

import (
	"fmt"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

var (
	modkernel32      = syscall.NewLazyDLL("kernel32.dll")
	procLockFileEx   = modkernel32.NewProc("LockFileEx")
	procUnlockFileEx = modkernel32.NewProc("UnlockFileEx")
)

/*
FileLock FileLock structure
*/
type FileLock struct {
	m  sync.Mutex
	fd syscall.Handle
}

/*
NewFileLock Create new FileLock instance
*/
func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic("filename needed")
	}
	fd, err := syscall.CreateFile(
		&(syscall.StringToUTF16(filename)[0]),
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_ALWAYS,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0)
	if err != nil {
		panic(err)
	}

	return &FileLock{fd: fd}

}

/*
Lock Create new FileLock instance
*/
func (m *FileLock) Lock() {
	m.m.Lock()
	var ol *syscall.Overlapped
	r1, _, el := syscall.Syscall6(
		procLockFileEx.Addr(),
		6,
		uintptr(m.fd),
		// LOCKFILE_EXCLUSIVE_LOCK = 0x00000002
		// 排他ロックのためのオプション
		// see https://docs.microsoft.com/en-us/windows/desktop/api/fileapi/nf-fileapi-lockfileex
		uintptr(0x00000002),
		uintptr(0),
		uintptr(1),
		uintptr(0),
		uintptr(unsafe.Pointer(ol)),
	)
	if r1 == 0 {
		panic(error(el))
	} else {
		panic(syscall.EINVAL)
	}
}

/*
Unlock Create new FileLock instance
*/
func (m *FileLock) Unlock() {
	var ol *syscall.Overlapped
	rl, _, el := syscall.Syscall6(
		procUnlockFileEx.Addr(),
		5,
		uintptr(m.fd),
		uintptr(0),
		uintptr(1),
		uintptr(0),
		uintptr(unsafe.Pointer(ol)),
		0,
	)
	if rl == 0 {
		if el != 0 {
			panic(error(el))
		} else {
			panic(syscall.EINVAL)
		}
	}
	m.m.Unlock()
}

// As of 2018/8/1, this program fails
func main() {
	l := NewFileLock("lock.go")
	fmt.Println("try locking...")
	l.Lock()
	fmt.Println("locked!")
	time.Sleep(10 * time.Second)
	l.Unlock()
	fmt.Println("unlock")
}

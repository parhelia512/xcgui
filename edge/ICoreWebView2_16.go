package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2_16 是 ICoreWebView2_15 接口的延续，支持打印相关功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_16
type ICoreWebView2_16 struct {
	Vtbl *ICoreWebView2_16Vtbl
}

type ICoreWebView2_16Vtbl struct {
	ICoreWebView2_15Vtbl
	Print            ComProc
	ShowPrintUI      ComProc
	PrintToPdfStream ComProc
}

func (i *ICoreWebView2_16) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_16) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_16) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
Print
ShowPrintUI
PrintToPdfStream
*/

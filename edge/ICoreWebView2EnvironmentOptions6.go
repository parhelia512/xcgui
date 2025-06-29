package edge

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"
)

// ICoreWebView2EnvironmentOptions6 提供用于创建 WebView2 环境以管理浏览器扩展的其他选项。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions6
type ICoreWebView2EnvironmentOptions6 struct {
	Vtbl *ICoreWebView2EnvironmentOptions6Vtbl
}

type ICoreWebView2EnvironmentOptions6Vtbl struct {
	IUnknownVtbl
	GetAreBrowserExtensionsEnabled ComProc
	PutAreBrowserExtensionsEnabled ComProc
}

func (i *ICoreWebView2EnvironmentOptions6) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions6) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions6) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAreBrowserExtensionsEnabled 获取是否启用浏览器扩展功能.
func (i *ICoreWebView2EnvironmentOptions6) GetAreBrowserExtensionsEnabled() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetAreBrowserExtensionsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetAreBrowserExtensionsEnabled 设置是否启用浏览器扩展功能. 默认为 false.
func (i *ICoreWebView2EnvironmentOptions6) SetAreBrowserExtensionsEnabled(value bool) error {
	r, _, _ := i.Vtbl.PutAreBrowserExtensionsEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetAreBrowserExtensionsEnabled 获取是否启用浏览器扩展功能.
func (i *ICoreWebView2EnvironmentOptions6) MustGetAreBrowserExtensionsEnabled() bool {
	value, err := i.GetAreBrowserExtensionsEnabled()
	ReportErrorAtuo(err)
	return value
}

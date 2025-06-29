package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2Settings2 是 ICoreWebView2Settings 接口的延续, 支持管理 UserAgent.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings2
type ICoreWebView2Settings2 struct {
	Vtbl *ICoreWebView2Settings2Vtbl
}

type ICoreWebView2Settings2Vtbl struct {
	ICoreWebView2SettingsVtbl
	GetUserAgent ComProc
	PutUserAgent ComProc
}

func (i *ICoreWebView2Settings2) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings2) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings2) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetUserAgent 获取 UserAgent。
func (i *ICoreWebView2Settings2) GetUserAgent() (string, error) {
	var _userAgent *uint16
	r, _, _ := i.Vtbl.GetUserAgent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_userAgent)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	userAgent := common.UTF16PtrToString(_userAgent)
	wapi.CoTaskMemFree(unsafe.Pointer(_userAgent))
	return userAgent, nil
}

// SetUserAgent 设置 UserAgent。
func (i *ICoreWebView2Settings2) SetUserAgent(userAgent string) error {
	_userAgent, err := syscall.UTF16PtrFromString(userAgent)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutUserAgent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_userAgent)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetUserAgent 获取 UserAgent。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings2) MustGetUserAgent() string {
	userAgent, err := i.GetUserAgent()
	ReportErrorAtuo(err)
	return userAgent
}

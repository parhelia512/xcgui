package wapi

import "syscall"

var (
	// Library.
	Ntdll = syscall.NewLazyDLL("ntdll.dll")

	// Functions.
	rtlMoveMemory = Ntdll.NewProc("RtlMoveMemory")
)

// RtlMoveMemory 将源内存块的内容复制到目标内存块，并支持重叠的源内存块和目标内存块.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/devnotes/rtlmovememory.
//
// Destination: 指向要复制字节的目标内存块的指针.
//
// Source: 指向要复制字节的源内存块的指针.
//
// Length: 从源复制到目标中的字节数.
func RtlMoveMemory(Destination uintptr, Source uintptr, Length uint) {
	rtlMoveMemory.Call(Destination, Source, uintptr(Length))
}

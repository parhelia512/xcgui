package xc

import (
	"syscall"
	"unsafe"
)

// 动画项_置回调.
//
// hAnimationItem: 动画项句柄.
//
// callback: 回调函数.
func XAnimaItem_SetCallback(hAnimationItem int, callback FunAnimationItem) {
	ptr := syscall.NewCallback(callback)
	finalCallbackPtr := createFloatCallback(ptr)
	xAnimaItem_SetCallback.Call(uintptr(hAnimationItem), finalCallbackPtr)
}

// Windows API: VirtualAlloc
var (
	virtualAlloc = kernel32.NewProc("VirtualAlloc")
)

// 汇编代码（机器码）.
// 功能：将 XMM1 (float) 移动到 RDX，然后跳转到 Go 函数.
// 逻辑：
//
//	MOVD RDX, XMM1  (把 XMM1 里的浮点数位模式移到 RDX)
//	MOV R10, [目标地址]
//	JMP R10
func createFloatCallback(targetFn uintptr) uintptr {
	// 机器码模板
	// 48 0F 7E CA  -> MOVD RDX, XMM1
	// 49 BA [8字节地址] -> MOV R10, imm64
	// 41 FF E2     -> JMP R10
	
	code := []byte{
		// MOVD RDX, XMM1 (前缀)
		0x66, 0x48, 0x0F, 0x7E, 0xCA,
		// MOV R10, imm64 (将目标函数地址放入 R10)
		0x49, 0xBA,
		// 8字节地址
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// JMP R10
		0x41, 0xFF, 0xE2,
	}

	// 填入 Go 函数地址（小端序）
	for i := 0; i < 8; i++ {
		code[7+i] = byte(targetFn >> (i * 8))
	}

	// 分配可执行内存
	// MEM_COMMIT (0x1000) | MEM_RESERVE (0x2000), PAGE_EXECUTE_READWRITE (0x40)
	mem, _, _ := virtualAlloc.Call(
		0,
		uintptr(len(code)),
		0x1000|0x2000,
		0x40)

	if mem == 0 {
		panic("VirtualAlloc failed")
	}

	// 将机器码写入内存
	for i, b := range code {
		*(*byte)(unsafe.Pointer(mem + uintptr(i))) = b
	}
	return mem
}

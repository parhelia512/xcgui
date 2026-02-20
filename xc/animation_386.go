package xc

import "syscall"

// 动画项_置回调.
//
// hAnimationItem: 动画项句柄.
//
// callback: 回调函数.
func XAnimaItem_SetCallback(hAnimationItem int, callback FunAnimationItem) {
	xAnimaItem_SetCallback.Call(uintptr(hAnimationItem), syscall.NewCallback(callback))
}

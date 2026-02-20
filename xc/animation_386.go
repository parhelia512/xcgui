package xc

import (
	"syscall"
)

// getAnimaItemCallbackPtr 获取动画项回调函数指针
func getAnimaItemCallbackPtr() uintptr {
	animaItemCallbackOnce.Do(func() {
		animaItemCallbackPtr = syscall.NewCallback(animaItemCallbackShell)
	})
	return animaItemCallbackPtr
}

package widget_test

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"testing"
	"time"
)

func TestElement_SetFocus(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		layout := widget.NewLayoutEle(50, 50, 150, 50, w.Handle)
		edit := widget.NewEdit(0, 0, 100, 30, layout.Handle)
		widget.NewButton(50, 100, 100, 30, "setfocus", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			edit.SetFocus()
			return 0
		})
	})
}

func TestElement_SetLeft(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		btn := widget.NewButton(10, 40, 100, 30, "setleft+5", w.Handle)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			btn.SetLeft(btn.GetLeft()+5, true)
			return 0
		})
	})
}

func TestElement_SetTop(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		btn := widget.NewButton(10, 80, 100, 30, "settop+5", w.Handle)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			btn.SetTop(btn.GetTop()+5, true)
			return 0
		})

		edit := widget.NewEdit(10, 40, 100, 30, w.Handle)
		edit.AddEvent_KeyDown(func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
			fmt.Println("键按下", wParam, lParam)
			return 0
		})
		edit.AddEvent_SysKeyDown(func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
			fmt.Println("系统键按下", wParam, lParam)
			return 0
		})
	})
}

func TestElement_AddEvent_MOUSESTAY(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		btn := widget.NewButton(20, 40, 300, 100, "按钮", w.Handle)
		btn.AddEvent_MouseStay(func(hEle int, pbHandled *bool) int {
			t.Log("鼠标进入")
			return 0
		})
		btn.AddEvent_MouseHover(func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
			t.Log("鼠标悬停", nFlags, pPt)
			return 0
		})
		btn.AddEvent_MouseLeave(func(hEle int, hEleStay int, pbHandled *bool) int {
			t.Log("鼠标离开", hEleStay)
			return 0
		})
	})
}

func Test_onXE_DESTROY_END(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		btn1 := widget.NewButton(20, 40, 300, 100, "2秒后自动销毁", w.Handle)
		btn1.AddEvent_Destroy(func(hEle int, pbHandled *bool) int {
			t.Log("btn1 销毁事件")
			return 0
		})

		time.AfterFunc(time.Second*2, func() {
			xc.XC_CallUT(func() {
				btn1.Destroy()
				w.Redraw(false)
			})
		})
	})
}

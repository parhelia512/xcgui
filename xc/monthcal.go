package xc

import (
	"github.com/twgh/xcgui/xcc"
	"unsafe"
)

// 月历_创建, 创建日期时间元素, 返回元素句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func XMonthCal_Create(x, y, cx, cy int32, hParent int) int {
	r, _, _ := xMonthCal_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 月历_取内部按钮, 获取内部按钮元素.
//
// hEle: 元素句柄.
//
// nType: 按钮类型, MonthCal_Button_Type_.
func XMonthCal_GetButton(hEle int, nType xcc.MonthCal_Button_Type_) int {
	r, _, _ := xMonthCal_GetButton.Call(uintptr(hEle), uintptr(nType))
	return int(r)
}

// 月历_置当前日期, 设置月历选中的年月日.
//
// hEle: 元素句柄.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func XMonthCal_SetToday(hEle int, nYear int32, nMonth int32, nDay int32) {
	xMonthCal_SetToday.Call(uintptr(hEle), uintptr(nYear), uintptr(nMonth), uintptr(nDay))
}

// 月历_取当前日期, 获取月历当前年月日.
//
// hEle: 元素句柄.
//
// pnYear: 年.
//
// pnMonth: 月.
//
// pnDay: 日.
func XMonthCal_GetToday(hEle int, pnYear *int32, pnMonth *int32, pnDay *int32) {
	xMonthCal_GetToday.Call(uintptr(hEle), uintptr(unsafe.Pointer(pnYear)), uintptr(unsafe.Pointer(pnMonth)), uintptr(unsafe.Pointer(pnDay)))
}

// 月历_取选择日期, 获取月历选中的年月日.
//
// hEle: 元素句柄.
//
// pnYear: 年.
//
// pnMonth: 月.
//
// pnDay: 日.
func XMonthCal_GetSelDate(hEle int, pnYear *int32, pnMonth *int32, pnDay *int32) {
	xMonthCal_GetSelDate.Call(uintptr(hEle), uintptr(unsafe.Pointer(pnYear)), uintptr(unsafe.Pointer(pnMonth)), uintptr(unsafe.Pointer(pnDay)))
}

// 月历_置文本颜色.
//
// hEle: 元素句柄.
//
// nFlag: 1:周六, 周日文字颜色, 2:日期文字的颜色; 其它周文字颜色, 使用元素自身颜色.
//
// color: xc.RGBA 颜色值.
func XMonthCal_SetTextColor(hEle int, nFlag int32, color uint32) {
	xMonthCal_SetTextColor.Call(uintptr(hEle), uintptr(nFlag), uintptr(color))
}

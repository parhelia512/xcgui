package webviewloader

import _ "embed"

//go:embed sdk/x86/WebView2Loader.dll
var Dll []byte

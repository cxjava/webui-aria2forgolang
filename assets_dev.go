// +build !vfs
//go:generate go run assets_generate.go

package main

import "net/http"

// Assets contains project assets.
var yaaw http.FileSystem = http.Dir("vfsdata/webui-aria2-master")
var aria http.FileSystem = http.Dir("vfsdata/yaaw-master")

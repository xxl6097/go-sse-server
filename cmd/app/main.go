package main

import (
	"fmt"
	"github.com/xxl6097/go-sse-server/internal/server"
	"github.com/xxl6097/go-sse-server/pkg"
	"github.com/xxl6097/go-sse-server/pkg/u"
)

func init() {
	if u.IsMacOs() {
		pkg.AppVersion = "v0.0.3"
		pkg.BinName = "sse_v0.0.20_darwin_arm64"
	}
}
func main() {
	fmt.Println("Hello World")
	server.Serve("admin", "admin", 9999)
}

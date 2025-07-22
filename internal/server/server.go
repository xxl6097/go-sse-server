package server

import (
	"github.com/xxl6097/go-http/pkg/httpserver"
	"github.com/xxl6097/go-sse-server/assets"
	"github.com/xxl6097/go-sse/pkg/sse"
	"github.com/xxl6097/go-sse/pkg/sse/iface"
	"net/http"
	"time"
)

func initSSE() iface.ISseServer {
	return sse.New().
		InvalidateFun(func(request *http.Request) (string, error) {
			return time.Now().Format("20060102150405.999999999"), nil
		}).
		Register(func(server iface.ISseServer, client *iface.Client) {
			server.Stream("内置丰富的开发模板，包括前后端开发所需的所有工具，如pycharm、idea、navicat、vscode以及XTerminal远程桌面管理工具等模板，用户可以轻松部署和管理各种应用程序和工具", time.Millisecond*500)
		}).
		UnRegister(nil).
		Done()
}
func Serve(username, password string, port int) {
	httpserver.New().
		CORSMethodMiddleware().
		BasicAuth(username, password).
		Handle("/api/sse", initSSE().Handler()).
		AddRoute(assets.NewRoute()).
		Done(port)
}

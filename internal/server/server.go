package server

import (
	"github.com/xxl6097/go-http/pkg/httpserver"
	"github.com/xxl6097/go-sse-server/assets"
	"github.com/xxl6097/go-sse-server/internal/api"
	"github.com/xxl6097/go-sse/pkg/sse"
	"github.com/xxl6097/go-sse/pkg/sse/isse"
	"net/http"
	"time"
)

func createSSE() isse.ISseServer {
	return sse.New().
		InvalidateFun(func(request *http.Request) (string, error) {
			return time.Now().Format("20060102150405.999999999"), nil
		}).
		Register(func(server isse.ISseServer, client *isse.Client) {
			server.Stream("内置丰富的开发模板，包括前后端开发所需的所有工具，如pycharm、idea、navicat、vscode以及XTerminal远程桌面管理工具等模板，用户可以轻松部署和管理各种应用程序和工具", time.Millisecond*500)
		}).
		UnRegister(nil).
		Done()
}

func Serve(username, password string, port int) {
	sseApi := createSSE()
	httpserver.New().
		CORSMethodMiddleware().
		BasicAuth(username, password).
		AddRoute(api.NewRoute(sseApi)).
		Handle("/api/sse", sseApi.Handler()).
		AddRoute(assets.NewRoute()).
		Done(port)
}

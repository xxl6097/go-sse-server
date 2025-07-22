package api

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/pkg/ihttpserver"
	"github.com/xxl6097/go-sse/pkg/sse/isse"
	"net/http"
)

type Route struct {
	sseApi isse.ISseServer
}

func NewRoute(sse isse.ISseServer) ihttpserver.IRoute {
	opt := &Route{
		sseApi: sse,
	}
	return opt
}

func (this *Route) Setup(router *mux.Router) {
	router.HandleFunc("/api/clients/get", this.GetClients).Methods(http.MethodGet)
}

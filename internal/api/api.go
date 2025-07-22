package api

import (
	"errors"
	"github.com/xxl6097/go-sse-server/pkg/u"
	"github.com/xxl6097/go-sse/pkg/sse/isse"
	"net/http"
)

func (this *Route) GetClients(w http.ResponseWriter, r *http.Request) {
	u.ProcessRequest(w, r, func() (any, error) {
		if this.sseApi == nil {
			return nil, errors.New("sseApi is nil")
		}
		clients := this.sseApi.GetClients()
		var data []*isse.Client
		for _, cls := range clients {
			data = append(data, cls)
		}
		return data, nil
	})
}

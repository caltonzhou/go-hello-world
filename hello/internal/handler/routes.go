// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5

package handler

import (
	"net/http"

	"go-hello-world/hello/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: HelloHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_api_version",
				Handler: GetApiVersionHandler(serverCtx),
			},
		},
	)
}

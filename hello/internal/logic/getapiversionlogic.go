package logic

import (
	"context"
	"fmt"
	"os"

	"go-hello-world/hello/internal/svc"
	"go-hello-world/hello/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiVersionLogic {
	return &GetApiVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiVersionLogic) GetApiVersion(req *types.ApiVersionRequest) (resp *types.ApiVersionResponse, err error) {
	// todo: add your logic here and delete this line

	value := os.Getenv("API_VERSION")

	if value == "" {
		fmt.Println("API_VERSION is not set")
		value = "unknown"
	} else {
		fmt.Println("API_VERSION:", value)
	}
	resp = &types.ApiVersionResponse{
		ApiVersion: value,
	}
	return
}

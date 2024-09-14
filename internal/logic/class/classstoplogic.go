package class

import (
	"context"

	"edumaster/internal/svc"
	"edumaster/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassStopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassStopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassStopLogic {
	return &ClassStopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassStopLogic) ClassStop(req *types.ClassStopReq) (resp *types.ClassStopResp, err error) {
	// todo: add your logic here and delete this line

	return
}

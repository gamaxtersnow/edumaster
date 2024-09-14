package classschedule

import (
	"context"

	"edumaster/internal/svc"
	"edumaster/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ClassScheduleListReq) (resp *types.ClassScheduleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

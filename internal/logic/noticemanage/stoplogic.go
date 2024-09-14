package noticemanage

import (
	"context"

	"edumaster/internal/svc"
	"edumaster/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopLogic {
	return &StopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopLogic) Stop(req *types.NoticeManageStopReq) (resp *types.NoticeManageStopResp, err error) {
	// todo: add your logic here and delete this line

	return
}

package notification

import (
	"context"

	"edumaster/internal/svc"
	"edumaster/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyLogic {
	return &ModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyLogic) Modify(req *types.NotificationModifyReq) (resp *types.NotificationModifyResp, err error) {
	// todo: add your logic here and delete this line

	return
}

package classattendancerecord

import (
	"context"

	"edumaster/internal/svc"
	"edumaster/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.ClassAttendanceRecordInfoReq) (resp *types.ClassAttendanceRecordInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

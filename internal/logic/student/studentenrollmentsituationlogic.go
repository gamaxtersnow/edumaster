package student

import (
	"context"

	"edumaster/internal/svc"
	"edumaster/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentenrollmentsituationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentenrollmentsituationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentenrollmentsituationLogic {
	return &StudentenrollmentsituationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentenrollmentsituationLogic) Studentenrollmentsituation(req *types.StudentenrollmentsituationReq) (resp *types.StudentEnrollmentSituationResp, err error) {
	// todo: add your logic here and delete this line

	return
}

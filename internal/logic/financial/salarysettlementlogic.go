package financial

import (
	"context"

	"edumaster/internal/svc"
	"edumaster/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SalarysettlementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSalarysettlementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SalarysettlementLogic {
	return &SalarysettlementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SalarysettlementLogic) Salarysettlement(req *types.SalarySettlementReq) (resp *types.SalarySettlementResp, err error) {
	// todo: add your logic here and delete this line

	return
}

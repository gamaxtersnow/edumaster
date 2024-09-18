package student

import (
	"context"
	"edumaster/internal/svc"
	"edumaster/internal/types"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.StudentAddReq) (resp *types.StudentAddResp, err error) {
	student, err := l.svcCtx.StudentModel.FindOneByPhoneNumber(l.ctx, req.PhoneNumber)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, err
	}
	if student != nil {
		return nil, errors.New("该手机号已注册")
	}
	err = l.svcCtx.StudentModel.AddStudent(l.ctx, l.svcCtx.ContactsModel, l.svcCtx.ParentAccountsModel, l.svcCtx.LearningGoalsModel, req)
	if err != nil {
		return nil, err
	}
	//将消息推送到es消息队列 todo
	return nil, nil

}

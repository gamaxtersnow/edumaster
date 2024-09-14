package student

import (
	"context"

	"edumaster/internal/svc"
	"edumaster/internal/types"

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
	if err := l.svcCtx.Validate.Struct(req); err != nil {
		return nil, err
	}
	// 调用 model 层添加学生
	err = l.svcCtx.StudentModel.AddStudent(l.ctx, l.svcCtx.ContactsModel, l.svcCtx.ParentAccountsModel, l.svcCtx.LearningGoalsModel, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
	//return &types.StudentAddResp{Message: "Student added successfully"}, nil

}

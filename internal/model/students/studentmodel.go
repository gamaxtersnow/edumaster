package students

import (
	"context"
	"edumaster/internal/types"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ StudentModel = (*customStudentModel)(nil)

type (
	// StudentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentModel.
	StudentModel interface {
		studentModel
		AddStudent(ctx context.Context, c ContactModel, p ParentAccountModel, l LearningGoalModel, req *types.StudentAddReq) error
	}

	customStudentModel struct {
		*defaultStudentModel
	}
)

// NewStudentModel returns a model for the database table.
func NewStudentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StudentModel {
	return &customStudentModel{
		defaultStudentModel: newStudentModel(conn, c, opts...),
	}
}
func (m *customStudentModel) AddStudent(ctx context.Context, c ContactModel, p ParentAccountModel, l LearningGoalModel, req *types.StudentAddReq) error {
	return m.TransCtx(ctx, func(session sqlx.Session) error {
		now := time.Now().Unix()
		data := &Student{
			Name:           req.Name,
			PhoneNumber:    req.PhoneNumber,
			AccountType:    req.AccountType,
			DateOfBirth:    req.DateOfBirth,
			City:           req.City,
			Gender:         req.Gender,
			Grade:          req.Grade,
			Major:          req.Major,
			WechatId:       req.WechatId,
			WechatNickname: req.WechatNickname,
			School:         req.School,
			StudentType:    req.StudentType,
			Subscription:   req.Subscription,
			LoginPassword:  req.LoginPassword,
			Notes:          req.Notes,
			UpdatedAt:      now,
			CreatedAt:      now,
		}
		res, err := m.TransInsertCtx(ctx, session, data)
		if err != nil {
			return err
		}
		stuId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		for _, contact := range req.Contacts {
			data := &Contact{
				StudentId:    stuId,
				Name:         contact.Name,
				PhoneNumber:  contact.PhoneNumber,
				WechatId:     contact.WechatId,
				Type:         contact.Type,
				Relationship: contact.Relationship,
				UpdatedAt:    now,
				CreatedAt:    now,
			}
			_, err = c.TransInsertCtx(ctx, session, data)
			if err != nil {
				return err
			}
		}
		for _, parent := range req.ParentAccounts {
			data := &ParentAccount{
				StudentId:     stuId,
				Name:          parent.Name,
				PhoneNumber:   parent.PhoneNumber,
				LoginPassword: parent.LoginPassword,
				Relationship:  parent.Relationship,
				UpdatedAt:     now,
				CreatedAt:     now,
			}
			_, err = p.TransInsertCtx(ctx, session, data)
			if err != nil {
				return err
			}
		}
		for _, learningGoal := range req.LearningGoals {
			data := &LearningGoal{
				StudentId:     stuId,
				TargetScore:   learningGoal.TargetScore,
				EntryScore:    learningGoal.EntryScore,
				TargetCountry: learningGoal.TargetCountry,
				TargetCourse:  learningGoal.TargetCourse,
				UpdatedAt:     now,
				CreatedAt:     now,
			}
			_, err = l.TransInsertCtx(ctx, session, data)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

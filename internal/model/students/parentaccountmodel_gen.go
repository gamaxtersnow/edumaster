// Code generated by goctl. DO NOT EDIT.

package students

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	parentAccountFieldNames          = builder.RawFieldNames(&ParentAccount{})
	parentAccountRows                = strings.Join(parentAccountFieldNames, ",")
	parentAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(parentAccountFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	parentAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(parentAccountFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheParentAccountIdPrefix = "cache:parentAccount:id:"
)

type (
	parentAccountModel interface {
		Insert(ctx context.Context, data *ParentAccount) (sql.Result, error)
		TransInsertCtx(ctx context.Context, session sqlx.Session, data *ParentAccount) (sql.Result, error)

		FindOne(ctx context.Context, id int64) (*ParentAccount, error)
		Update(ctx context.Context, data *ParentAccount) error
		Delete(ctx context.Context, id int64) error
		TransCtx(ctx context.Context, fn func(session sqlx.Session) error) error
	}

	defaultParentAccountModel struct {
		sqlc.CachedConn
		table string
	}

	ParentAccount struct {
		Id            int64  `db:"id"`
		StudentId     int64  `db:"student_id"`     // 关联的学生ID
		Name          string `db:"name"`           // 家长姓名
		Relationship  string `db:"relationship"`   // 与学生的关系
		PhoneNumber   string `db:"phone_number"`   // 家长手机号
		LoginPassword string `db:"login_password"` // 家长登录密码
		CreatedAt     int64  `db:"created_at"`     // 创建时间（Unix时间戳）
		UpdatedAt     int64  `db:"updated_at"`     // 更新时间（Unix时间戳）
	}
)

func newParentAccountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultParentAccountModel {
	return &defaultParentAccountModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`parent_account`",
	}
}

func (m *defaultParentAccountModel) Delete(ctx context.Context, id int64) error {
	parentAccountIdKey := fmt.Sprintf("%s%v", cacheParentAccountIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, parentAccountIdKey)
	return err
}

func (m *defaultParentAccountModel) TransCtx(ctx context.Context, fn func(session sqlx.Session) error) error {
	err := m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		err := fn(session)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *defaultParentAccountModel) FindOne(ctx context.Context, id int64) (*ParentAccount, error) {
	parentAccountIdKey := fmt.Sprintf("%s%v", cacheParentAccountIdPrefix, id)
	var resp ParentAccount
	err := m.QueryRowCtx(ctx, &resp, parentAccountIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", parentAccountRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultParentAccountModel) Insert(ctx context.Context, data *ParentAccount) (sql.Result, error) {
	parentAccountIdKey := fmt.Sprintf("%s%v", cacheParentAccountIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, parentAccountRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.StudentId, data.Name, data.Relationship, data.PhoneNumber, data.LoginPassword)
	}, parentAccountIdKey)
	return ret, err
}

func (m *defaultParentAccountModel) TransInsertCtx(ctx context.Context, session sqlx.Session, data *ParentAccount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, parentAccountRowsExpectAutoSet)
	return session.ExecCtx(ctx, query, data.StudentId, data.Name, data.Relationship, data.PhoneNumber, data.LoginPassword)
}

func (m *defaultParentAccountModel) Update(ctx context.Context, data *ParentAccount) error {
	parentAccountIdKey := fmt.Sprintf("%s%v", cacheParentAccountIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, parentAccountRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.StudentId, data.Name, data.Relationship, data.PhoneNumber, data.LoginPassword, data.Id)
	}, parentAccountIdKey)
	return err
}

func (m *defaultParentAccountModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheParentAccountIdPrefix, primary)
}

func (m *defaultParentAccountModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", parentAccountRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultParentAccountModel) tableName() string {
	return m.table
}

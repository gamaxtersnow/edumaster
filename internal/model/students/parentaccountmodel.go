package students

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ParentAccountModel = (*customParentAccountModel)(nil)

type (
	// ParentAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customParentAccountModel.
	ParentAccountModel interface {
		parentAccountModel
	}

	customParentAccountModel struct {
		*defaultParentAccountModel
	}
)

// NewParentAccountModel returns a model for the database table.
func NewParentAccountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ParentAccountModel {
	return &customParentAccountModel{
		defaultParentAccountModel: newParentAccountModel(conn, c, opts...),
	}
}

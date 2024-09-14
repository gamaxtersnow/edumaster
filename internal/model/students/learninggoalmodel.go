package students

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LearningGoalModel = (*customLearningGoalModel)(nil)

type (
	// LearningGoalModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLearningGoalModel.
	LearningGoalModel interface {
		learningGoalModel
	}

	customLearningGoalModel struct {
		*defaultLearningGoalModel
	}
)

// NewLearningGoalModel returns a model for the database table.
func NewLearningGoalModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LearningGoalModel {
	return &customLearningGoalModel{
		defaultLearningGoalModel: newLearningGoalModel(conn, c, opts...),
	}
}

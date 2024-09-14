package svc

import (
	"edumaster/internal/config"
	"edumaster/internal/model/qiwei"
	"edumaster/internal/model/students"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type ServiceContext struct {
	Config              config.Config
	Localizer           *i18n.Localizer
	Cache               cache.Cache
	Validate            *validator.Validate
	QwHttpClient        *qiwei.HttpClient
	StudentModel        students.StudentModel
	ContactsModel       students.ContactsModel
	ParentAccountsModel students.ParentAccountsModel
	LearningGoalsModel  students.LearningGoalsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.MustNewRedis(redis.RedisConf{
		Host: c.Cache[0].Host,
		Type: c.Cache[0].Type,
		Pass: c.Cache[0].Pass,
		Tls:  c.Cache[0].Tls,
	})
	stat := &cache.Stat{}
	errNotFound := errors.New("item not found in cache")
	cacheNode := cache.NewNode(redisClient, syncx.NewSingleFlight(), stat, errNotFound)
	httpClient := &http.Client{}
	cookieJar, _ := cookiejar.New(nil)
	httpClient.Jar = cookieJar
	httpClient.Timeout = 15 * time.Second
	QwHttpClient := qiwei.NewHttpClient(300, c.QWApiConf, cacheNode)
	return &ServiceContext{
		Config:              c,
		Cache:               cacheNode,
		Validate:            validator.New(),
		QwHttpClient:        QwHttpClient,
		StudentModel:        students.NewStudentModel(conn, c.Cache),
		ContactsModel:       students.NewContactsModel(conn, c.Cache),
		ParentAccountsModel: students.NewParentAccountsModel(conn, c.Cache),
		LearningGoalsModel:  students.NewLearningGoalsModel(conn, c.Cache),
	}
}

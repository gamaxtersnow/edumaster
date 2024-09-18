package svc

import (
	"edumaster/internal/config"
	"edumaster/internal/model/qiwei"
	"edumaster/internal/model/students"
	"errors"
	zh_locales "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"strings"
	"time"
)

type ServiceContext struct {
	Config              config.Config
	Localizer           *i18n.Localizer
	Cache               cache.Cache
	Validate            *validator.Validate
	Trans               ut.Translator
	QwHttpClient        *qiwei.HttpClient
	StudentModel        students.StudentModel
	ContactsModel       students.ContactModel
	ParentAccountsModel students.ParentAccountModel
	LearningGoalsModel  students.LearningGoalModel
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
	validate, trans := initValidator()
	return &ServiceContext{
		Config:              c,
		Cache:               cacheNode,
		Validate:            validate,
		Trans:               trans,
		QwHttpClient:        QwHttpClient,
		StudentModel:        students.NewStudentModel(conn, c.Cache),
		ContactsModel:       students.NewContactModel(conn, c.Cache),
		ParentAccountsModel: students.NewParentAccountModel(conn, c.Cache),
		LearningGoalsModel:  students.NewLearningGoalModel(conn, c.Cache),
	}
}
func initValidator() (*validator.Validate, ut.Translator) {
	validate := validator.New()
	// 注册自定义错误消息
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	zhT := zh_locales.New()
	uni := ut.New(zhT, zhT)
	trans, _ := uni.GetTranslator("zh")
	_ = zh.RegisterDefaultTranslations(validate, trans)
	return validate, trans
}

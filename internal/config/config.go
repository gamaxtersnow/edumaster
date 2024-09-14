package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	QWApiConf QWApiConf
	Mysql     struct {
		DataSource string
	}
	Cache cache.CacheConf
}
type QWApiConf struct {
	CorpId  string
	Secret  string
	BaseUrl string
	Api     map[string]Api
}
type Api struct {
	Name string
	Uri  string
}

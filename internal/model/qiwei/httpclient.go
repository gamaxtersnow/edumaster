package qiwei

import (
	"edumaster/internal/config"
	"edumaster/internal/types"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const qwAuthorizationTokenKey = "qw-authorization-token"

type HttpClient struct {
	client  *http.Client
	ticker  *time.Ticker
	limiter chan struct{}
	conf    config.QWApiConf
	cache   cache.Cache
}

func NewHttpClient(rate int, conf config.QWApiConf, cache cache.Cache) *HttpClient {
	ticker := time.NewTicker(time.Minute / time.Duration(rate))
	limiter := make(chan struct{}, rate)
	go func() {
		defer close(limiter)
		for range ticker.C {
			limiter <- struct{}{}
		}
	}()
	return &HttpClient{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		ticker:  ticker,
		limiter: limiter,
		conf:    conf,
		cache:   cache,
	}
}

func (h *HttpClient) Token() string {
	token := ""
	err := h.cache.Get(qwAuthorizationTokenKey, &token)
	if err == nil {
		return token
	}
	if h.cache.IsNotFound(err) {
		params := url.Values{}
		params.Set("corpid", h.conf.CorpId)
		params.Set("corpsecret", h.conf.Secret)
		resp, err := h.client.Get(h.getRequestUrl("get_access_token") + "?" + params.Encode())
		if err != nil {
			logx.Errorf("请求微信token出错，err：%v", err)
			time.Sleep(time.Second)
			return h.Token()
		}
		defer func() {
			_ = resp.Body.Close()
		}()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logx.Errorf("读取返回结果出错，err为：%v", err)
			time.Sleep(time.Second)
			return h.Token()
		}
		qwToken := &types.QwTokenResp{}
		err = json.Unmarshal(body, qwToken)
		if err != nil {
			logx.Errorf("json解析错误，返回结果为：%s", body)
			time.Sleep(time.Second)
			return h.Token()

		}
		if qwToken.AccessToken == "" {
			logx.Errorf("未获取token，返回结果为：%s", body)
			time.Sleep(time.Second)
			return h.Token()

		}
		_ = h.cache.SetWithExpire(qwAuthorizationTokenKey, qwToken.AccessToken, time.Second*time.Duration(qwToken.ExpiresIn-600))
		return qwToken.AccessToken
	}
	logx.Errorf("redis出错，未获取token，err：%v", err)
	time.Sleep(time.Second)
	return h.Token()
}

func (h *HttpClient) Get(name string, params url.Values) (*http.Response, error) {
	<-h.limiter
	req, err := h.NewRequest(http.MethodGet, h.getRequestUrl(name)+"&"+params.Encode(), nil, nil)
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HttpClient) Post(name string, body io.Reader) (*http.Response, error) {
	<-h.limiter
	req, err := h.NewRequest(http.MethodPost, h.getRequestUrl(name), nil, body)
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HttpClient) getRequestUrl(name string) string {
	return strings.TrimRight(h.conf.BaseUrl, "/") + "/" + strings.TrimPrefix(h.conf.Api[name].Uri, "/")
}

func (h *HttpClient) NewRequest(method, requestUrl string, header map[string]string, body io.Reader) (*http.Request, error) {
	params := url.Values{}
	params.Set("access_token", h.Token())
	req, err := http.NewRequest(method, requestUrl+"?"+params.Encode(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	return req, nil
}

func (h *HttpClient) Stop() {
	h.ticker.Stop()
}

package qwapi

import (
	"beisen/internal/model/qiwei"
	"beisen/internal/types"
	"bytes"
	"context"
	"encoding/json"
	"io"
)

var _ QwUserModel = (*customQwUserModel)(nil)

const BusinessEmailType = 1
const PersonalBusinessType = 2

type (
	QwUserModel interface {
		GetUserIdByEmail(ctx context.Context, params *types.QwUserIdByEmailReq) (*types.QwUserIdResp, error)
		GetUserIdByMobile(ctx context.Context, params *types.QwUserIdByMobileReq) (*types.QwUserIdResp, error)
	}
	customQwUserModel struct {
		client *qiwei.HttpClient
	}
)

func NewCustomQwUserModel(client *qiwei.HttpClient) QwUserModel {
	return &customQwUserModel{
		client: client,
	}
}
func (m *customQwUserModel) GetUserIdByEmail(ctx context.Context, params *types.QwUserIdByEmailReq) (*types.QwUserIdResp, error) {
	jsonBytes, err := json.Marshal(params)
	resp, _ := m.client.Post("get_userid_by_email", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	qwUserIdResp := &types.QwUserIdResp{}
	err = json.Unmarshal(body, qwUserIdResp)
	if err != nil {
		return nil, err
	}
	return qwUserIdResp, nil
}
func (m *customQwUserModel) GetUserIdByMobile(ctx context.Context, params *types.QwUserIdByMobileReq) (*types.QwUserIdResp, error) {
	jsonBytes, err := json.Marshal(params)
	resp, _ := m.client.Post("get_userid_by_mobile", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	qwUserIdResp := &types.QwUserIdResp{}
	err = json.Unmarshal(body, qwUserIdResp)
	if err != nil {
		return nil, err
	}
	return qwUserIdResp, nil
}

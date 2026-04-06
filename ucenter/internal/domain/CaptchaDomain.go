package domain

import (
	"encoding/json"
	"mscoin-common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaDomain struct {
}
type vaptchaReq struct {
	Id        string `json:"id"`
	Secretkey string `json:"secretkey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}
type vaptchaRsp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

func NewCaptchaDomain() *CaptchaDomain {
	return &CaptchaDomain{}
}

func (d *CaptchaDomain) Verify(server string, vid string, key string, token string, scene int, ip string) bool {
	//发送POST请求
	req, err := tools.Post(server, &vaptchaReq{
		Id:        vid,
		Secretkey: key,
		Scene:     scene,
		Token:     token,
		Ip:        ip,
	})
	if err != nil {
		logx.Error(err)
		return false
	}
	result := &vaptchaRsp{}
	err = json.Unmarshal(req, result)
	if err != nil {
		logx.Error(err)
		return false
	}
	return result.Success == 1
}

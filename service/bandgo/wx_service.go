package bandgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	bandgoModel "github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
	bandgoReq "github.com/flipped-aurora/gin-vue-admin/server/model/bandgo/request"
)

type WeiXinService struct{}

var wxCache = &bandgoModel.WxCache{}

// GetAccessToken 获取微信access_token
func (s *WeiXinService) GetAccessToken(req bandgoReq.WXLogin) (string, error) {
	// 检查缓存是否有效
	if wxCache.AccessToken != "" && time.Now().Unix() < wxCache.AccessTokenExpires {
		return wxCache.AccessToken, nil
	}

	// 检查配置
	fmt.Println("global.GVA_CONFIG.WeChat", global.GVA_CONFIG.WeChat)
	if global.GVA_CONFIG.WeChat.AppID == "" || global.GVA_CONFIG.WeChat.AppSecret == "" {
		return "", errors.New("微信配置缺失，请检查配置")
	}

	// 请求access_token

	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s&code=%s",
		global.GVA_CONFIG.WeChat.WxBaseTokenURL,
		global.GVA_CONFIG.WeChat.AppID,
		global.GVA_CONFIG.WeChat.AppSecret,
		req.Code)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if result.ErrCode != 0 {
		return "", fmt.Errorf("获取access_token失败: %s", result.ErrMsg)
	}

	// 更新缓存
	wxCache.AccessToken = result.AccessToken
	wxCache.AccessTokenExpires = time.Now().Unix() + result.ExpiresIn - 300 // 提前5分钟过期

	return wxCache.AccessToken, nil
}

// GetJsapiTicket 获取微信jsapi_ticket
func (s *WeiXinService) GetJsapiTicket() (string, error) {
	// 检查缓存是否有效
	if wxCache.JsapiTicket != "" && time.Now().Unix() < wxCache.JsapiTicketExpires {
		return wxCache.JsapiTicket, nil
	}

	// 获取access_token
	accessToken, err := s.GetAccessToken(bandgoReq.WXLogin{})
	if err != nil {
		return "", err
	}

	// 请求jsapi_ticket
	url := fmt.Sprintf("%s?access_token=%s&type=jsapi", global.GVA_CONFIG.WeChat.WxJsapiTicketURL, accessToken)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		Ticket    string `json:"ticket"`
		ExpiresIn int64  `json:"expires_in"`
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
	}
	fmt.Println("body", body)
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		return "", fmt.Errorf("获取jsapi_ticket失败: %s", result.ErrMsg)
	}

	// 更新缓存
	wxCache.JsapiTicket = result.Ticket
	wxCache.JsapiTicketExpires = time.Now().Unix() + result.ExpiresIn - 300 // 提前5分钟过期

	return wxCache.JsapiTicket, nil
}

// GetUserInfo 获取用户信息
func (s *WeiXinService) GetWeiXinUserInfo(accessToken string) (interface{}, error) {
	url := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN",
		global.GVA_CONFIG.WeChat.WxUserInfoURL,
		accessToken,
		"") // TODO: 需要从access_token响应中获取openid

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if errcode, ok := result["errcode"]; ok {
		return nil, fmt.Errorf("获取用户信息失败: %v", errcode)
	}

	return result, nil
}

package common

import (
	"fmt"
	"net/http"

	"github.com/shawnwang5/go-wechat-invoice/config"
	commonModel "github.com/shawnwang5/go-wechat-invoice/model/common"
	requestutils "github.com/shawnwang5/go-wechat-invoice/utils/request"
)

const (
	BASE_URL = config.BASE_URL
)

// GetAccessToken 获取微信的 Access Token
//
// 参数：
//   - appId: 应用 id
//   - appSecret: 应用 secret
//
// 返回值：
//   - res: 响应数据
//   - err: error
func GetAccessToken(appId, appSecret string) (res *commonModel.GetAccessTokenRes, err error) {
	// url := fmt.Sprintf(`%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`, BASE_URL, appId, appSecret)
	url := fmt.Sprintf(`%s/cgi-bin/stable_token`, BASE_URL)
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := make(map[string]interface{})
	data["grant_type"] = "client_credential" // 授权类型，固定值
	data["appid"] = appId                    // 应用 id
	data["secret"] = appSecret               // 应用 secret
	data["force_refresh"] = false            // 是否强制刷新
	res, err = requestutils.HttpRequest[commonModel.GetAccessTokenRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

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
	url := fmt.Sprintf(`%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`, BASE_URL, appId, appSecret)
	method := http.MethodGet
	headers := make(map[string]string)
	params := make(map[string]string)
	data := make(map[string]interface{})
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

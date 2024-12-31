package common

// 获取微信的 Access Token
type GetAccessTokenRes struct {
	AccessToken string `json:"access_token" mapstructure:"access_token"` // 获取到的凭证
	ExpiresIn   uint   `json:"expires_in" mapstructure:"expires_in"`     // 凭证有效时间，单位：秒
	Errcode     int    `json:"errcode" mapstructure:"errcode"`           // 错误码
	Errmsg      string `json:"errmsg" mapstructure:"errmsg"`             // 错误信息
}

package common

import "testing"

// 测试 GetAccessTokenTest
func TestGetAccessToken(t *testing.T) {
	appId := "xxxx"
	appSecret := "xxxx"
	res, err := GetAccessToken(appId, appSecret)
	if err != nil {
		t.Errorf("GetAccessToken error: %s", err.Error())
		return
	}

	if res.AccessToken == "" {
		t.Errorf("GetAccessToken error: %s", "AccessToken is empty")
		return
	}
}

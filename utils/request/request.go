package requestutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func HttpRequest[T any](
	urlStr string,
	method string,
	headers map[string]string,
	params map[string]string,
	data any) (res *T, err error) {
	// 创建URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	query := u.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	u.RawQuery = query.Encode()

	// 将数据编码为JSON
	buf := new(bytes.Buffer)
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	// 创建请求
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 发送 form 请求
func PostForm[T any](
	urlStr string,
	data map[string]string) (res *T, err error) {
	formData := url.Values{}
	for k, v := range data {
		formData.Set(k, v)
	}

	// 创建请求
	resp, err := http.PostForm(urlStr, formData)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// ReadDataFromResponse 从 http.Response 中读取数据
//
// 参数：
//   - resp: http.Response 对象
//
// 返回值：
//   - data: 响应数据
//   - err: error
func ReadDataFromResponse[T any](resp *http.Response) (res *T, err error) {
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

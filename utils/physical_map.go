package utils

import (
	"anmo/utils/request"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bitly/go-simplejson"
)

func GetQQMapCity(qqMapApiUrl string, qqMapKey string, latitude string, longitude string) (city string, err error) {
	urlStr := qqMapApiUrl //腾讯地图api接口地址
	method := "GET"
	headers := map[string]string{}
	params := map[string]string{
		"location": latitude + "," + longitude, //纬度latitude,经度longitude
		"key":      qqMapKey,                   //腾讯地图key
	}
	var data map[string]string
	resp, err1 := request.HttpRequest(urlStr,
		method,
		headers,
		params,
		data,
	)
	if err1 != nil {
		return "", err1
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("err2:", errors.New(resp.Status).Error())
		return "", errors.New("网络异常")
	}

	body, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("err3:", err1)
		return "", err1
	}
	fmt.Println("body:")
	fmt.Println(string(body))

	res, err1 := simplejson.NewJson(body)

	if err1 != nil {
		fmt.Printf("%v\n", err)
		return "", err1
	}

	if status, e := res.Get("status").Int(); e != nil {
		fmt.Println("get:", e)
		return "", e
	} else {
		fmt.Println("status:", status)
		if status == 0 {
			if result := res.Get("result"); result != nil {
				if ad := result.Get("ad_info"); ad != nil {
					if ci := ad.Get("city"); ci != nil {
						city, err = ci.String()
						fmt.Println("get city:", city)
						return
					} else {
						return "", errors.New("不存在键:city")
					}
				} else {
					return "", errors.New("不存在键:ad_info")
				}
			} else {
				return "", errors.New("不存在键:result")
			}
		}
	}
	message := "传参有误"
	if msg := res.Get("message"); msg != nil {
		if message1, e1 := msg.String(); e1 == nil {
			message = message1
		}
	}
	return "", errors.New(message)
}

package open

import (
	"fmt"
	"order/api/pkg/helper"
	"sort"
	"strings"
	"time"
)

type PinDuoDuo struct {
	AppKey      string `json:"app_key"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
}

func NewPinDuoDuo(appKey, AppSecret, accessToken string) *PinDuoDuo {
	return &PinDuoDuo{
		AppKey:      appKey,
		AppSecret:   AppSecret,
		AccessToken: accessToken,
	}
}

func (p *PinDuoDuo) GenSignData(url string, params map[string]interface{}) map[string]interface{} {
	commonParams := make(map[string]interface{})
	commonParams["app_key"] = p.AppKey
	commonParams["access_token"] = p.AccessToken
	commonParams["type"] = url
	commonParams["timestamp"] = time.Now().Unix()

	newParams := helper.MergeMap(commonParams, params)

	var keys []string
	for k := range newParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var dataStr string
	for _, k := range keys {
		dataStr = fmt.Sprintf("%s%s%v", dataStr, k, newParams[k])
	}
	dataStr = fmt.Sprintf("%s%v%s", p.AppSecret, dataStr, p.AppSecret)

	newParams["sign"] = strings.ToUpper(helper.Md5V(dataStr))

	return newParams
}

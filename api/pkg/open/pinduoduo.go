package open

import (
	"sort"
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

func (p *PinDuoDuo) genSign(url string, params map[string]interface{}) {
	commonParams := make(map[string]interface{})
	commonParams["app_key"] = p.AppKey
	commonParams["access_token"] = p.AccessToken
	commonParams["type"] = url
	commonParams["timestamp"] = time.Now().Unix()

	newParams := mergeMap(commonParams, params)

	var keys []string
	for k := range newParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

}

func mergeMap[T int | string](params1, params2 map[T]interface{}) map[T]interface{} {
	newParams := make(map[T]interface{})
	for i, v := range params1 {
		for j, w := range params2 {
			if i == j {
				newParams[i] = w
			} else {
				if _, ok := newParams[i]; !ok {
					newParams[i] = v
				}
				if _, ok := newParams[j]; !ok {
					newParams[j] = w
				}
			}
		}
	}
	return newParams
}

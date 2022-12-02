package open

import (
	"fmt"
	"github.com/guonaihong/gout"
	"order/api/pkg/helper"
	"testing"
)

func TestPinDuoDuo_GenSignData(t *testing.T) {
	
	params := helper.H{"page": 1, "pageSize": 10}
	pinDuoDuoClient := NewPinDuoDuo("f6be2ad9ca52697fbc558b6dad102ffb", "3ebc27525da043dabe1554cb2e64197e60c6541c", "hzm3a55lmixoenkkixxuuolukzshkyvrxfmzi1qadq0xws8aptgtgami")
	reqParams := pinDuoDuoClient.GenSignData("bg.goods.label.get", params)
	
	var resp interface{}
	host := "https://kj-openapi.pinduoduo.com/openapi/router"
	err := gout.POST(host).Debug(true).SetJSON(reqParams).BindBody(&resp).Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

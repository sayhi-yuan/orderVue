package open

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order/api/pkg/helper"
	"strings"
	"testing"
)

func TestPinDuoDuo_GenSignData(t *testing.T) {
	params := helper.H{"page": 1, "pageSize": 10}
	pinDuoDuoClient := NewPinDuoDuo("f6be2ad9ca52697fbc558b6dad102ffb", "3ebc27525da043dabe1554cb2e64197e60c6541c", "hzm3a55lmixoenkkixxuuolukzshkyvrxfmzi1qadq0xws8aptgtgami")
	reqParams := pinDuoDuoClient.GenSignData("bg.goods.label.get", params)

	client := &http.Client{}

	marshal, _ := json.Marshal(&reqParams)

	resp, err := client.Post("https://kj-openapi.pinduoduo.com/openapi/router", "application/json", strings.NewReader(string(marshal)))
	if err != nil {
		panic(err)
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(all))
}

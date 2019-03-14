package main

import (
	"fmt"
	huawei "go_sample/huaweiPush"
)

func main() {
	ClientId := "100307583"
	ClientSecret := "01b8c58558f3e1f2aa8d8b633308a986"
	client := huawei.NewClient(ClientId, ClientSecret)

	// getToken
	// accessToken := client.GetToken()
	// fmt.Println("accessToken", accessToken)

	token := []string{"0861331032881895300000533500CN01", "111"}
	payload := huawei.NewMessage().SetContent("huawei-content").SetTitle("huawei-title").SetIntent("com://hzhu.m.logo/openwith?link=%s&push_id=%s").SetAppPkgName("com.hzhu.m").Json()
	fmt.Println(payload)
	result, _ := client.PushMsgList(token, payload)
	fmt.Println("result", result)
}


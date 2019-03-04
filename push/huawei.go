package main

import (
	"fmt"
	"golang.org/x/net/context"
	"github.com/yilee/huawei-push"
)

func main() {
	client := &huaweipush.HuaweiPushClient{
		clientID:     "100307583",
		clientSecret: "01b8c58558f3e1f2aa8d8b633308a986",
	}
	res, err := client.SingleSend(context.TODO(), huaweipush.NewSingleNotification("0861331032881895300000533500CN01", "message").SetRequestID("requestID").SetHighPriority())
	fmt.Println(res, err)
}

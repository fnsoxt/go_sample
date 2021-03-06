package main

import (
	"fmt"
	"golang.org/x/net/context"
	"github.com/yilee/huawei-push"
)

func main() {
	client := huaweipush.NewClient("100307583", "01b8c58558f3e1f2aa8d8b633308a986")
	singleNotification := huaweipush.NewSingleNotification("0861331032881895300000533500CN01", "message").SetRequestID("requestID").SetHighPriority()
	res, err := client.SingleSend(context.TODO(), singleNotification)
	fmt.Println(res, err)
}

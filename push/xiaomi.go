package main

import (
	"fmt"
	"golang.org/x/net/context"
	"github.com/yilee/xiaomi-push"
)

func main() {
	var client = xiaomipush.NewClient("Y2Gqtk49HSXzSc4FM1pkvA==", []string{"com.hzhu.m"})
	var msg1 *xiaomipush.Message = xiaomipush.NewAndroidMessage("title", "body").SetPayload("this is payload1")
	res, err := client.SendToAlias(context.Background(), msg1, "5434710")
	fmt.Println(res, err)
}

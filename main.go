package main

import (
	"os"

	"github.com/RyaoChengfeng/wzj-checkin/app"
	"github.com/RyaoChengfeng/wzj-checkin/config"
)

func main() {
	if len(os.Args) < 2 {
		go app.StartCron()
		wait := make(chan any)
		<-wait
	}
	openId := os.Args[1]
	// https://v18.teachermate.cn/wechat-pro-ssr/?openid=94ca7f35917470493be402e4ab130316&from=wzj
	app.CheckInTask(openId, config.C.CheckIn.Lon, config.C.CheckIn.Lat)
}

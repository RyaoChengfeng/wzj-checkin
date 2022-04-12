package util

import (
	"github.com/RyaoChengfeng/wzj-checkin/config"
	. "github.com/RyaoChengfeng/wzj-checkin/util/log"
	"github.com/robfig/cron"
	"strconv"
)

func StartCron() {
	c := cron.New()
	Logger.Info("start checkInTask, cron: " + config.TimerEveryTwentySecond)
	err := c.AddFunc(config.TimerEveryTwentySecond, CheckInTask)
	if err != nil {
		Logger.Error("cron add func error:", err)
	}
	c.Start()
}

func CheckInTask() {
	lon, _ := strconv.ParseFloat(config.C.CheckIn.Lon, 64)
	lat, _ := strconv.ParseFloat(config.C.CheckIn.Lat, 64)
	coordinate := Coordinate{
		Lon: lon,
		Lat: lat,
	}
	ok, err := UserCheckIn(config.C.CheckIn.OpenID, coordinate)
	if ok {
		if err != nil {
			Logger.Info("check in success, error in not nil: ", err)
		}
		Logger.Info("check status active")
	} else {
		Logger.Info("check in fail, error: ", err)
	}
}

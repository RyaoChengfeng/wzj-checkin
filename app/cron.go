package app

import (
	"github.com/RyaoChengfeng/wzj-checkin/config"
	. "github.com/RyaoChengfeng/wzj-checkin/log"
	"github.com/robfig/cron"
)

func StartCron() {
	c := cron.New()
	Logger.Info("start checkInTask, cron: " + config.Timer)
	err := c.AddFunc(config.Timer, cronCheckIn)
	if err != nil {
		Logger.Error("cron add func error:", err)
	}
	c.Start()
}

func cronCheckIn() {
	CheckInTask(config.C.CheckIn.OpenID, config.C.CheckIn.Lon, config.C.CheckIn.Lat)
}

func CheckInTask(openId string, lon float64, lat float64) {
	Logger.Infof("try check in...")
	coordinate := Coordinate{
		Lon: lon,
		Lat: lat,
	}
	err := UserCheckIn(openId, coordinate)
	if err != nil {
		Logger.Info("check in error: ", err)
	}
}

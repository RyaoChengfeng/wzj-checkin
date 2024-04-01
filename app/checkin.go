package app

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/RyaoChengfeng/wzj-checkin/config"
	. "github.com/RyaoChengfeng/wzj-checkin/log"
)

type SignActive struct {
	CourseId  int    `json:"courseId"`
	SignId    int    `json:"signId"`
	IsGPS     int    `json:"isGPS"`
	IsQR      int    `json:"isQR"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	StartYear int    `json:"startYear"`
	Term      string `json:"term"`
	Cover     string `json:"cover"`
}

type CheckIn struct {
	CourseId int `json:"courseId"`
	SignId   int `json:"signId"`
}

type SignIn struct {
	SignRank    int `json:"signRank"`
	StudentRank int `json:"studentRank"`
}

type Coordinate struct {
	Lon float64 `bson:"lon" json:"lon"`
	Lat float64 `bson:"lat" json:"lat"`
}

func UserCheckIn(openId string, coordinate Coordinate) error {
	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate, br, zstd",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,ja;q=0.6",
		"Connection":      "keep-alive",
		"Content-Type":    "application/json",
		"Host":            "v18.teachermate.cn",
		"openId":          openId,
	}
	body, err := HttpGet(config.WzjSignActive, headers)
	if err != nil {
		Logger.Errorf("WzjSignActive error: %v", err)
		return err
	}
	var signActive []SignActive
	err = json.Unmarshal(body, &signActive)
	if err != nil {
		Logger.Errorf("json Unmarshal error: %v", err)
		return err
	}
	for _, signCode := range signActive {
		payload := CheckIn{
			CourseId: signCode.CourseId,
			SignId:   signCode.SignId,
		}
		if signCode.IsGPS != 0 {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			coordinate.Lon += float64(rand.Intn(40)-20) * 0.000001
			coordinate.Lat += float64(rand.Intn(40)-20) * 0.000001
			headers["lon"] = strconv.FormatFloat(coordinate.Lon, 'f', 5, 64)
			headers["lat"] = strconv.FormatFloat(coordinate.Lat, 'f', 5, 64)
		}

		body, err := HttpPost(config.WzjSignIn, headers, payload)
		if err != nil {
			Logger.Errorf("WzjSignIn error: %v", err)
			return err
		}
		var signIn SignIn
		err = json.Unmarshal(body, &signIn)
		if err != nil {
			Logger.Errorf("json Unmarshal error: %v", err)
			return err
		}
		Logger.Infof("CheckIn success. SignRank:%v, StudentRank:%v", signIn.SignRank, signIn.StudentRank)
	}
	return nil
}

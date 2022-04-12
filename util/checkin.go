package util

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/RyaoChengfeng/wzj-checkin/config"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//type WZJCourse struct {
//	ID           int    `json:"id"`
//	Name         string `json:"name"`
//	Topic        string `json:"topic"`
//	Code         string `json:"code"`
//	College      string `json:"college"`
//	Department   string `json:"department"`
//	DiscussionID int    `json:"discussionId"`
//	Selected     bool   `json:"selected"`
//}

type Coordinate struct {
	Lon float64 `bson:"lon" json:"lon"`
	Lat float64 `bson:"lat" json:"lat"`
}

var nameMap = map[string]string{
	"token-hash": "wx_csrf_name",
	"openid":     "openid",
	"sign-id":    "signId",
	"course-id":  "courseId",
}

// UserCheckIn see https://github.com/yun-mu/wzj-sign-in-weixin
func UserCheckIn(textOpenid string, coordinate Coordinate) (bool, error) {
	client := &http.Client{}

	// Request the HTML page.
	res, err := http.Get(fmt.Sprintf(config.URLWZJSignIn, textOpenid))
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, errors.New("request wrong")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return false, err
	}

	data := url.Values{}
	// is-sign apply-gps is-qr-sign
	applyGps := false
	doc.Find("input[type=hidden]").Each(func(i int, s *goquery.Selection) {
		if id, idOk := s.Attr("id"); idOk {
			if name, nameOk := nameMap[id]; nameOk {
				value, _ := s.Attr("value")
				data.Set(name, value)
			}
			if id == "apply-gps" {
				if value, _ := s.Attr("value"); value == "1" {
					applyGps = true
				}
			}
		}
	})

	if data.Get("courseId") != "" && data.Get("openid") != "" && data.Get("signId") != "" {
		if applyGps {
			rand.Seed(time.Now().UnixNano())
			coordinate.Lon += float64(rand.Intn(40)-20) * 0.000001
			coordinate.Lat += float64(rand.Intn(40)-20) * 0.000001
			data.Set("lon", strconv.FormatFloat(coordinate.Lon, 'f', 5, 64))
			data.Set("lat", strconv.FormatFloat(coordinate.Lat, 'f', 5, 64))
		} else {
			data.Set("lon", "0")
			data.Set("lat", "0")
		}

		req, err := http.NewRequest("POST", config.URLWZJStuSignIn, ioutil.NopCloser(strings.NewReader(data.Encode())))
		if err != nil {
			return false, err
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_4 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12H143 MicroMessenger/6.2.3 NetType/WIFI Language/zh_CN")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		res, err = client.Do(req)
		if err != nil {
			return false, err
		}
		defer res.Body.Close()
		if res.StatusCode == 200 {
			return true, nil
		}
	}
	return false, err
}

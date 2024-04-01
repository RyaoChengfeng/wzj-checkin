package config

const (
	APIPrefix     = "/api"
	Timer         = "@every 10s"
	WzjSignActive = "https://v18.teachermate.cn/wechat-api/v1/class-attendance/student/active_signs"
	WzjSignIn     = "https://v18.teachermate.cn/wechat-api/v1/class-attendance/student-sign-in"
)

//GET /wechat-api/v1/class-attendance/student/active_signs HTTP/1.1
//Accept: */*
//Accept-Encoding: gzip, deflate, br, zstd
//Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,ja;q=0.6
//Connection: keep-alive
//Content-Type: application/json
//Cookie: session=eyJvcGVuSWQiOiI5NGNhN2YzNTkxNzQ3MDQ5M2JlNDAyZTRhYjEzMDMxNiJ9; session.sig=B-B3er_WLXmArgfcPzoG02TvE_w
//Host: v18.teachermate.cn
//Referer: https://v18.teachermate.cn/wechat-pro-ssr/student/sign?openid=94ca7f35917470493be402e4ab130316
//Sec-Fetch-Dest: empty
//Sec-Fetch-Mode: cors
//Sec-Fetch-Site: same-origin
//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36
//openId: 94ca7f35917470493be402e4ab130316
//sec-ch-ua: "Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"
//sec-ch-ua-mobile: ?0
//sec-ch-ua-platform: "macOS"

// [{"courseId":1352000,"signId":3108925,"isGPS":0,"isQR":0,"name":"测试课堂","code":"MU452","startYear":2024,"term":"春","cover":"https://app.teachermate.com.cn/covers/science4.png"}]

//POST /wechat-api/v1/class-attendance/student-sign-in HTTP/1.1
//Accept: */*
//Accept-Encoding: gzip, deflate, br, zstd
//Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,ja;q=0.6
//Connection: keep-alive
//Content-Length: 37
//Content-Type: application/json
//Cookie: session=eyJvcGVuSWQiOiI5NGNhN2YzNTkxNzQ3MDQ5M2JlNDAyZTRhYjEzMDMxNiJ9; session.sig=B-B3er_WLXmArgfcPzoG02TvE_w
//Host: v18.teachermate.cn
//Origin: https://v18.teachermate.cn
//Referer: https://v18.teachermate.cn/wechat-pro-ssr/student/sign/list/1352000
//Sec-Fetch-Dest: empty
//Sec-Fetch-Mode: cors
//Sec-Fetch-Site: same-origin
//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36
//openId: 94ca7f35917470493be402e4ab130316
//sec-ch-ua: "Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"
//sec-ch-ua-mobile: ?0
//sec-ch-ua-platform: "macOS"

// {"signRank":3,"studentRank":1}

//https://www.teachermate.com.cn/api/v1/qr/attendance/b4e739ed2685a6bb2afcc5e40a3546eb3b4d8ddf63a14c4edf5068c36a0ea53d865814be898b6bad84763654925462a2

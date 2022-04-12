package config

const (
	APIPrefix = "/api"

	TimerEveryTwentySecond = "@every 20s"

	URLWechatAccessToken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	URLWZJSignIn         = "https://v18.teachermate.cn/wechat/wechat/guide/signin?openid=%s"
	URLWZJStuSignIn      = "https://v18.teachermate.cn/wechat-api/v1/class-attendance/student-sign-in"
	URLWZHDisCourseList  = "https://v18.teachermate.cn/wechat-api/v1/students/courses?type=discussions"
	URLWZJCourseSelect   = "https://v18.teachermate.cn/wechat-api/v1/discussions/select"
)
## 微助教自动签到（堂堂复活）

1、configure `env/config.yml`

2、Start

- **Just use `go run main.go` to run it**

- **start in docker**

  You can create a docker image with `docker build -t wzj-checkin .`

  Use `docker-compose up` to start via docker

> 没想到大四还要用上这玩意，微助教好像只有华科在用
> 
> 需要注意的是，openId会玄学失效，建议在签到前获取openId并配置，然后手动run一下，用于对付使用二维码和地理位置的课程
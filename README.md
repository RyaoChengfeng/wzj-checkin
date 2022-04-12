## 微助教自动签到

1、configure `env/config.yml`

2、Start

- **Just use `go run main.go` to run it**

- **start in docker**

  You can create a docker image with `docker build -t wzj-checkin .`

  Use `docker-compose up` to start via docker

> API有一堆问题，但我目前懒得管，建议每次直接配置checkin_url让它一直跑着吧（**
FROM golang:1.22.0 AS builder
COPY . /src
WORKDIR /src
ENV GOPROXY="https://goproxy.io" GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app  .

FROM alpine:3.15
COPY ./env /env
COPY --from=builder /build/app /app
ENTRYPOINT ["sh","-c","/app"]
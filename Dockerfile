FROM golang:1.17-alpine AS build

LABEL maintainer="mail@xiaoliu.org"

WORKDIR /go/src/app

ADD . .

ENV GOPROXY="https://goproxy.cn,direct" \
    CGO_ENABLED=0

RUN go build -ldflags="-w -s" -o zhimiao-app

FROM zhimiao/alpine:latest AS prod

WORKDIR /zhimiao

COPY --from=build /go/src/app/zhimiao-app .

RUN chmod +x zhimiao-app

EXPOSE 1325

ENTRYPOINT ["/zhimiao/zhimiao-app"]

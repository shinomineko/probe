FROM golang:1.16-alpine as build

WORKDIR /go/src/probe

COPY ./*.go /go/src/probe

RUN go mod init probe \
	&& go get -u github.com/labstack/echo/v4 \
	&& go mod tidy \
	&& go build -o /go/bin/probe

FROM alpine:3.13

RUN adduser user -D -h /nonexistent -s /sbin/nologin --gecos ""

COPY --from=build /go/bin/probe /probe

USER user

EXPOSE 8080

ENTRYPOINT [ "/probe" ]

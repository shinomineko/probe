FROM golang:1.16-alpine as build

WORKDIR /go/src/probe

COPY ./*.go /go/src/probe

RUN go mod init probe \
	&& go get -u github.com/labstack/echo/v4 \
	&& go mod tidy \
	&& CGO_ENABLED=0 go build -o /go/bin/probe

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/probe /probe

EXPOSE 8080

ENTRYPOINT [ "/probe" ]

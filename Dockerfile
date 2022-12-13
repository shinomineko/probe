FROM cgr.dev/chainguard/go:latest as build

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /work/probe

FROM cgr.dev/chainguard/static:latest

COPY --from=build /work/probe /probe

EXPOSE 8080

ENTRYPOINT [ "/probe" ]

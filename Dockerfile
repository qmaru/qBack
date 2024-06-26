FROM golang:1.22-alpine as builder

RUN apk add upx ca-certificates tzdata

WORKDIR /usr/src

COPY . /usr/src

RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w -extldflags='static'" -trimpath -o app \
    && upx --best --lzma app

FROM scratch as prod

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/src/app /qBack

ENTRYPOINT ["/qBack"]

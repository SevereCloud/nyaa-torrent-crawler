FROM golang:1.12 AS builder
WORKDIR /go/src/github.com/SevereCloud/nyaa-torrent-crawler
COPY . .

RUN go get -d -v ./...
RUN  CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -v

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/src/github.com/SevereCloud/nyaa-torrent-crawler/nyaa-torrent-crawler /
CMD ["/nyaa-torrent-crawler"]

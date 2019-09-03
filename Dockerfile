FROM alpine AS builder
RUN apk update && apk add ca-certificates 

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY nyaa-torrent-crawler /
CMD ["/nyaa-torrent-crawler"]

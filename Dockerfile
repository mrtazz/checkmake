FROM golang:1.13 as builder

COPY . /checkmake

WORKDIR /checkmake

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make binaries

FROM alpine:3.9

USER nobody

COPY --from=builder /checkmake /
ENTRYPOINT ["./checkmake", "/Makefile"]

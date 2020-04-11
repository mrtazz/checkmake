FROM golang:1.13 as builder

COPY . /checkmake

RUN cd /go/src/github.com/mrtazz/checkmake && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make binaries
RUN cd /go/src/github.com/mrtazz/checkmake && make test

FROM alpine:3.9
RUN apk add make
USER nobody

COPY --from=builder /checkmake /
ENTRYPOINT ["./checkmake", "/Makefile"]

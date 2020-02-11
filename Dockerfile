FROM golang:1.13 as builder

COPY . /go/src/github.com/mrtazz/checkmake 

RUN cd /go/src/github.com/mrtazz/checkmake && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make binaries

FROM alpine:3.9

USER nobody

COPY --from=builder /go/src/github.com/mrtazz/checkmake/checkmake /
ENTRYPOINT ["./checkmake", "/Makefile"]

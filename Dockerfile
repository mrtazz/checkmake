FROM golang:1.13 as builder

ARG BUILDER_NAME
ARG BUILDER_EMAIL

COPY . /go/src/github.com/mrtazz/checkmake

RUN cd /go/src/github.com/mrtazz/checkmake && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make binaries
RUN cd /go/src/github.com/mrtazz/checkmake && make test

FROM alpine:3.9
RUN apk add make
USER nobody

COPY --from=builder /go/src/github.com/mrtazz/checkmake/checkmake /
ENTRYPOINT ["./checkmake", "/Makefile"]

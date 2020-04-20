FROM golang:1.13 as builder

ARG BUILDER_NAME BUILDER_EMAIL

COPY . /go/src/github.com/mrtazz/checkmake

WORKDIR /go/src/github.com/mrtazz/checkmake
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make binaries
RUN make test

FROM alpine:3.9
RUN apk add make
USER nobody

COPY --from=builder /go/src/github.com/mrtazz/checkmake/checkmake /
ENTRYPOINT ["./checkmake", "/Makefile"]

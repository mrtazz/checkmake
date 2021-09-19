FROM golang:1.13 as builder

ENV BUILDER_NAME=checkmake
ENV BUILDER_EMAIL=mrtazz@github.com
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY . /checkmake
WORKDIR /checkmake

RUN make binaries
RUN make test

FROM alpine:3.9
RUN apk add make
USER nobody

COPY --from=builder /checkmake /
ENTRYPOINT ["./checkmake", "/Makefile"]

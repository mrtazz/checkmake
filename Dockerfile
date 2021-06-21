FROM golang:1.13 as builder

ARG BUILDER_NAME
ARG BUILDER_EMAIL

COPY . /checkmake

RUN cd /checkmake && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make binaries
RUN cd /checkmake && make test

FROM alpine:3.9
RUN apk add make
USER nobody

COPY --from=builder /checkmake /
ENTRYPOINT ["./checkmake", "/Makefile"]

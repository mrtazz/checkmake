FROM golang:1.13 as builder

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
COPY . /checkmake

WORKDIR /checkmake

RUN make binaries

FROM alpine:3.9

USER nobody

COPY --from=builder /checkmake /
ENTRYPOINT ["./checkmake", "/Makefile"]

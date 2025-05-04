FROM golang:1.24 AS builder

ARG HAMMER_VERSION="latest"

RUN CGO_ENABLED=0 go install github.com/daichirata/hammer@${HAMMER_VERSION}   

FROM gcr.io/distroless/static

COPY --from=builder /go/bin/hammer /

ENTRYPOINT [ "/hammer" ]
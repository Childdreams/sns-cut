FROM golang:latest as builder

ARG GOPROXY
ENV GORPOXY ${GOPROXY}

ADD . /builder

WORKDIR /builder

RUN go build cmd/snsserver/main.go

FROM panwenbin/alpinetz:latest

COPY --from=builder /builder/main /app/rpc-sns-cut

WORKDIR /app

CMD ["./rpc-sns-cut"]

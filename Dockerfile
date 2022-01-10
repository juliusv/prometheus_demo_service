FROM golang:1.17-alpine as builder

WORKDIR /source
COPY . /source

RUN go mod download
RUN go build -v -o prometheus_demo_service .

FROM        alpine:3
MAINTAINER  Julius Volz <julius.volz@gmail.com>

COPY --from=builder /source/prometheus_demo_service  /bin/prometheus_demo_service

EXPOSE     8080
ENTRYPOINT [ "/bin/prometheus_demo_service" ]

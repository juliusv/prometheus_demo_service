FROM golang:1.17-alpine as builder
WORKDIR /source
COPY . /source

RUN go mod download
RUN go build -v -o prometheus_demo_service .

FROM        alpine:3
ARG PORT_ARG=8080
ENV PORT=${PORT_ARG}
MAINTAINER  Julius Volz <julius.volz@gmail.com>
COPY --from=builder /source/prometheus_demo_service  /bin/prometheus_demo_service
EXPOSE     ${PORT}
ENTRYPOINT [ "/bin/prometheus_demo_service" ]

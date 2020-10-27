FROM        quay.io/prometheus/busybox:latest
MAINTAINER  Julius Volz <julius.volz@gmail.com>

COPY prometheus_demo_service  /bin/prometheus_demo_service

EXPOSE     8080
ENTRYPOINT [ "/bin/prometheus_demo_service" ]

FROM alpine:3.20
LABEL maintainer="Juraj Bubniak <juraj.bubniak@gmail.com>"

RUN addgroup -S foxesscloud_exporter \
    && adduser -D -S -s /sbin/nologin -G foxesscloud_exporter foxesscloud_exporter

RUN apk --no-cache add tzdata ca-certificates

COPY foxesscloud_exporter /bin

USER foxesscloud_exporter

ENTRYPOINT ["foxesscloud_exporter"]
CMD ["server"]

FROM alpine

COPY run.sh /root
RUN set -x && \
    chmod +x /root/run.sh && \
    echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories && \
    echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories && \
    apk update && \
    apk upgrade && \
    apk add --no-cache mongodb && \
    rm /usr/bin/mongoperf && \
    rm -rf /var/cache/apk/*

VOLUME /data/db
EXPOSE 27017 28017

ENTRYPOINT [ "/root/run.sh" ]
CMD ["mongod"]

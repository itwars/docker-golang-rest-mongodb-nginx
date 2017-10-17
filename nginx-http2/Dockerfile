FROM alpine:edge
RUN	set -x \
	&& apk update \
	&& apk upgrade \
	&& apk add --no-cache nginx inotify-tools openssl 
RUN     openssl req 	-x509 -nodes \
			-days 365 \
			-newkey rsa:2048 \
			-keyout server.key \
			-out server.crt  \
			-subj "/C=FR/ST=Aquitaine/L=Bordeaux/O=MeMyselfAndI/OU=IT Department/CN=webapp.local" \
	&& openssl dhparam -out dhparam.pem 2048 \
	&& mkdir /etc/certs \
	&& mv server.* /etc/certs \
	&& mv dhparam.pem /etc/certs \
	&& apk del openssl \
	&& rm -rf /var/cache/apk/* 
COPY nginx.conf /etc/nginx/nginx.conf
COPY reload.sh /
RUN chmod +x reload.sh
RUN mkdir -p /run/nginx
EXPOSE 80 443
CMD ["/reload.sh"]


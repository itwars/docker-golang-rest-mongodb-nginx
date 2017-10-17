#!/bin/sh

nginx -g "daemon off;" &

while true
do
  inotifywait -e create -e modify /etc/certs /etc/nginx/conf.d/
  nginx -t
    if [ $? -eq 0 ]
      then
  	echo "Reloading Nginx Configuration"
  	nginx -s reload
      fi
done

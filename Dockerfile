FROM golang:1.16.4-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
LABEL maintainer="Viktor Nikitash <nikitashvictor@gmail.com>"
RUN mkdir /app
COPY . ./app

EXPOSE 9000

ADD start.sh /
RUN chmod +x /start.sh

CMD ["/start.sh"]
#CMD ["go","run","cmd/*.go"]
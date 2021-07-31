FROM alpine:3.14.0

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git img

CMD ["/bin/sh"]

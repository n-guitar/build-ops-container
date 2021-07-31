FROM alpine:3.14.0
LABEL maintainer="github.com/n-guitar"

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git img

CMD ["/bin/sh"]

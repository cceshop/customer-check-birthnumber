FROM golang:1.15.2-alpine3.12

LABEL cce.image.vendor="Crypto Currency Exchange s.r.o."
LABEL cce.image.maintainer="Crypto Currency Exchange s.r.o."
LABEL cce.git.repository="https://github.com/cceshop/customer-check-birthnumber.git"
LABEL cce.quay.repository="https://quay.io/repository/cceshop/customer-check-birthnumber"

ONBUILD USER root
RUN adduser --disabled-password --gecos '' -s /bin/bash cce \
    && mkdir -p /app
COPY ./entrypoint.sh /app
COPY ./api.go /app
RUN chmod +x /app/entrypoint.sh \ 
    && apk add git tini \
    && go get github.com/gorilla/mux \
    && chown -R cce:cce /app 

USER cce
WORKDIR /app

EXPOSE 8080
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/entrypoint.sh"]

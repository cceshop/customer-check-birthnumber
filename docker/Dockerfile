FROM golang:1.12.4-alpine3.9

LABEL cce.image.vendor="Crypto Currency Exchange s.r.o."
LABEL cce.image.maintainer="Crypto Currency Exchange s.r.o."
LABEL cce.git.repository="https://github.com/cceshop/customer-check-cardID.git"
LABEL cce.quay.repository="https://quay.io/repository/cceshop/customer-check-cardid"

USER root
RUN mkdir -p /app
WORKDIR /app
COPY ./entrypoint.sh /app
COPY ./api_check_cardID.go /app
RUN chmod +x /app/entrypoint.sh \ 
    && apk add git \
    && go get github.com/gorilla/mux

ENTRYPOINT ["/app/entrypoint.sh"]

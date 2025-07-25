FROM alpine:3.20

RUN apk add --no-cache \
    iputils \
    net-tools \
    curl \
    tree \
    bash \
    libc6-compat

WORKDIR /app

COPY orion /usr/local/bin/orion

VOLUME ["/storage"]

EXPOSE 8080

ENTRYPOINT ["orion"]
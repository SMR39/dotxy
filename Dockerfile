
FROM alpine:3.6

RUN apk add --update ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY dist/linux_amd64/dotxy /usr/local/bin/dotxy
ENV DNS_RESOLVER_ADDR=1.1.1.1:853
ENV LISTEN_ADDR=:53

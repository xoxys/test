FROM arm32v7/alpine:3.16@sha256:0c673ee68853a04014c0c623ba5ee21ee700e1d71f7ac1160ddb2e31c6fdbb18

ARG BUILD_DATE

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"
LABEL test="$BUILD_DATE"

RUN apk --no-cache add netcat-openbsd

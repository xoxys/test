FROM arm32v7/alpine:3.13@sha256:5fae3e61d4571e541775abf1098bdb9912e0122a1007852a30e9212c5e85472f

ARG BUILD_DATE

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"
LABEL test="$BUILD_DATE"

RUN apk --no-cache add netcat-openbsd

FROM --platform=linux/amd64 alpine:3.17

ARG BUILD_DATE

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"
LABEL test="$BUILD_DATE"

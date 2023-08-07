FROM --platform=linux/arm64/v8 alpine:3.17@sha256:f71a5f071694a785e064f05fed657bf8277f1b2113a8ed70c90ad486d6ee54dc

ARG BUILD_DATE

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"
LABEL test="$BUILD_DATE"

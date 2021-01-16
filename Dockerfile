FROM arm32v7/alpine:3.13

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"

RUN apk --no-cache add netcat-openbsd

FROM arm32v7/alpine:3.16

# renovate: depName=https://github.com/xoxys/test-drone depBranch=main
ENV GIT_REF_VERSION="${GIT_REF_VERSION:-50bd841b6ff4afca276eec5a11cc863e6f51d404}"

ARG BUILD_DATE

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"
LABEL test="$BUILD_DATE"

RUN apk --no-cache add netcat-openbsd

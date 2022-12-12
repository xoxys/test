FROM arm32v7/alpine:3.16

# renovate: depName=https://github.com/xoxys/test-drone depBranch=main
ENV GIT_REF_VERSION="${GIT_REF_VERSION:-ecd414ca4c2bdc1a45fc8dd0c4dd31ad8713a039}"

ARG BUILD_DATE

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"
LABEL test="$BUILD_DATE"

RUN apk --no-cache add netcat-openbsd

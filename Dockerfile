FROM arm32v7/alpine:3.17@sha256:68a5b7d32422e42b98bedfe2aef4d0b3445f69f0efe390ba2204427d80179a92

# renovate: depName=https://github.com/xoxys/test-drone depBranch=main
ENV GIT_REF_VERSION="${GIT_REF_VERSION:-ecd414ca4c2bdc1a45fc8dd0c4dd31ad8713a039}"

ARG BUILD_DATE

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"
LABEL test="$BUILD_DATE"

RUN apk --no-cache add netcat-openbsd

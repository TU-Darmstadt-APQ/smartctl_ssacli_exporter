ARG SSACLI_VERSION=6.15-11.0

FROM golang:1.21.1-bookworm as builder

ARG GIT_REPOSITORY
ARG SSH_DEPLOY_KEY

RUN \
  apt-get update && \
  apt-get upgrade -y && \
  wget https://downloads.linux.hpe.com/SDR/repo/mcp/Debian/pool/non-free/ssacli-${SSACLI_VERSION}_amd64.deb && \
  mkdir /root/.ssh/ && \
  echo "${SSH_DEPLOY_KEY}" > /root/.ssh/id_rsa && \
  chmod 600 /root/.ssh/id_rsa && \
  ssh-keyscan github.com >> /root/.ssh/known_hosts && \
  git clone git@github.com:${GIT_REPOSITORY}.git app && \
  cd app && \
  go mod init smartctl_ssacli_exporter && \
  go get && \
  go build -o smartctl_ssacli_exporter

FROM debian:12.1-slim
LABEL maintainer="Patrick Baus <patrick.baus@physik.tu-darmstadt.de>"
ARG SSACLI_VERSION

COPY --from=builder /go/app/smartctl_ssacli_exporter /sbin/smartctl_ssacli_exporter
COPY --from=builder /go/ssacli-${SSACLI_VERSION}_amd64.deb /

# Upgrade installed packages
RUN apt-get update && \
  apt-get upgrade -y && \
  apt-get install -y --no-install-recommends  \
    smartmontools \
    procps && \
  dpkg -i ssacli-${SSACLI_VERSION}_amd64.deb && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* && \
  rm ssacli-${SSACLI_VERSION}_amd64.deb

ENTRYPOINT ["/sbin/smartctl_ssacli_exporter"]

FROM golang:1.19 AS LOCAL-BUILD

WORKDIR /go/app

ARG SKAFFOLD_GO_GCFLAGS
ARG BUILD_DONT_STRIP

COPY Makefile go.mod ./
RUN go mod download

RUN --mount=type=ssh \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    make install-deps

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=ssh \
    make api-server-amd64

FROM alpine:3.16 AS RUNTIME

RUN \
   apk add --no-cache \
     curl

# UID and GID must match those defined security context of the deployment, cron job, etc.
ARG UID=1000

ENV HOME=/home/shady
RUN \
  adduser -h ${HOME} --uid ${UID} -D shady
ENV USER=shady

COPY --from=LOCAL-BUILD --chown=shady:shady /go/app/api-server-amd64 /opt/shady/shady-api/bin/api-server

RUN \
  mkdir -p /opt/shady && \
  chown -R shady:shady /opt/shady

CMD ["/opt/shady/shady-api/bin/api-server"]
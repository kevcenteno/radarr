FROM golang:1.13
WORKDIR /go/src/github.com/SkYNewZ/radarr

COPY go.* ./
RUN go mod download

COPY . .
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    export COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null) && \
    export VERSION=$(git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.0-${COMMIT_HASH}") && \
    go build -ldflags "-X main.version=${VERSION}" -a -installsuffix cgo -o /radarr ./cmd/radarr


FROM scratch
LABEL description="Radarr command-line utility"
LABEL maintainer="Quentin Lemaire <quentin@lemairepro.fr>"

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /radarr /radarr
ENTRYPOINT ["/radarr"]
CMD ["--help"]

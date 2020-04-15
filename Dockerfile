FROM golang:1.13
WORKDIR /go/src/github.com/SkYNewZ/radarr

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /radarr ./cmd/radarr


FROM scratch

LABEL description="Radarr command-line utility"
LABEL maintainer="Quentin Lemaire <quentin@lemairepro.fr>"

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /radarr /radarr
ENTRYPOINT ["/radarr"]
CMD ["--help"]

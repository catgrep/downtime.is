FROM golang:1.24.2-alpine AS build

ARG VERSION=dev
ENV CGO_ENABLED=0
WORKDIR /go/src/server
COPY . /go/src/server
RUN <<EOF
go build -ldflags="-X main.version=${VERSION}" -o /usr/bin/server .
server -version
EOF

FROM alpine:3.21.3

COPY --from=build /usr/bin/server /usr/bin/server
EXPOSE 8080
CMD ["server", "-port", "8080"]
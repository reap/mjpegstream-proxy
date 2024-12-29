FROM golang:1.22 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 go build -v -o ./proxy

FROM scratch
COPY --from=builder /build/proxy /usr/local/bin/proxy

ENV STREAM_URL="http://www.example.com/stream"
ENV STREAM_USER=""
ENV STREAM_PASSWORD=""

EXPOSE 8080

ENTRYPOINT [ "/usr/local/bin/proxy" ]
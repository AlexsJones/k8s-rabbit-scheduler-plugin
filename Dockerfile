# syntax = docker/dockerfile:experimental

FROM golang:1.14-alpine3.11 as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /bin/rabbit-scheduler

FROM alpine:3.11

COPY --from=builder /bin/rabbit-scheduler /bin/rabbit-scheduler

CMD ["/bin/rabbit-scheduler"]

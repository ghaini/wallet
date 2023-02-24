FROM golang:1.19-alpine AS builder
RUN apk --no-cache add ca-certificates git
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOSUMDB=off

WORKDIR /go/src/wallet

COPY . .
RUN go mod download
RUN go mod tidy
RUN  go build -o /go/bin/wallet .

FROM alpine:latest
RUN apk add --no-cache bash
COPY --from=builder /go/bin/wallet .
ENTRYPOINT ["./wallet"]

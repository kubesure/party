FROM golang:1.12-alpine AS builder
RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig
RUN apk update && apk add --no-cache git
WORKDIR /go/src/github.com/kubesure/party
COPY . .
RUN CGO_ENABLED=0 go install

FROM scratch
WORKDIR /opt
COPY --from=builder /go/bin/party .
EXPOSE 50051
CMD ["/opt/party"]
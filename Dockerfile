FROM golang:1.12.7-alpine AS builder
ENV GOBIN=/go/bin
WORKDIR /go/src/github.com/kubesure/party
COPY . .
RUN CGO_ENABLED=0 go install party.go
COPY vendor /go/src
RUN CGO_ENABLED=0 go install github.com/grpc-ecosystem/grpc-health-probe

FROM scratch
WORKDIR /opt
COPY --from=builder /go/bin/. .
EXPOSE 50051
CMD ["/opt/party"]
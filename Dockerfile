FROM golang:1.12.7-alpine AS builder
ENV GOBIN=/go/bin
WORKDIR /go/src/github.com/kubesure/party
COPY . .
RUN CGO_ENABLED=0 go install party.go

FROM scratch
WORKDIR /opt
COPY --from=builder /go/bin/party .
COPY --from=builder /go/src/github.com/kubesure/party/grpc-health-probe .
EXPOSE 50051
CMD ["/opt/party"]
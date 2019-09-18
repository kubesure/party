package main

import (
	"context"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc/health"

	api "github.com/kubesure/party/api/v1"
	service "github.com/kubesure/party/service/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
}

//starts a party service server and an health check server for k8s readiness probes.
func main() {
	log.Info("party server on...")
	ctx := context.Background()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	svc := &service.PartyService{}
	api.RegisterPartyServiceServer(s, svc)
	reflection.Register(s)

	h := health.NewServer()
	h.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, h)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Info("shutting down party server...")
			s.GracefulStop()
			<-ctx.Done()
		}
	}()
	s.Serve(lis)
}

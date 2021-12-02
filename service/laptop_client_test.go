package service_test

import (
	"context"
	"grpc/pb"
	"grpc/sample"
	"grpc/serializer"
	"grpc/service"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := service.NewInMemoryLaptopStore()
	serverAddress := startTestLaptopServer(t, laptopStore)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, err)
	require.Equal(t, expectedID, res.Id)

	other, err := laptopStore.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, err)

	requireSameLaptop(t, laptop, other)
}

func startTestLaptopServer(t *testing.T, l service.LaptopStore) string {
	laptopServer := service.NewLaptopServer(l)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, s string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(s, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, l1 *pb.Laptop, l2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(l1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(l2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}

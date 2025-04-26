package grpcport

import (
	"context"
	"time"

	pb "github.com/GabrielEValenzuela/RefuCare/internal/ports/grpc/proto"

	"google.golang.org/grpc"
)

type PatientRecordsClient struct {
	client pb.PatientRecordsClient
}

func NewPatientRecordsClient(address string) (*PatientRecordsClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	return &PatientRecordsClient{
		client: pb.NewPatientRecordsClient(conn),
	}, nil
}

func (c *PatientRecordsClient) GetPatientInfo(id string) (*pb.PatientInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.client.GetPatientInfo(ctx, &pb.PatientId{Id: id})
}

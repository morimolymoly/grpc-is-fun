package main

import (
	"context"
	"errors"

	pb "github.com/morimolymoly/grpc-is-fun/api"
)

/*
import (
	"fmt"
	pb "github.com/morimolymoly/grpc-is-fun/api"
)*/

// MyAnimalServiceServer ... animalapiのサービス
type MyAnimalServiceServer struct {
}

// GetAllAnimals ... animalsを返す
func (s *MyAnimalServiceServer) GetAllAnimals(ctx context.Context, req *pb.GetAllAnimalsRequest) (*pb.Animals, error) {
	return nil, errors.New("a")
}

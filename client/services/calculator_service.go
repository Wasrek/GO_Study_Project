package services

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CalcultorService interface {
	Hello(name string) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalcultorService {
	return calculatorService{calculatorClient}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name:        name,
		CreatedDate: timestamppb.Now(),
	}

	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	fmt.Printf("Service : Hello\n")
	fmt.Printf("Request : %v\n", req.Name)
	fmt.Printf("Response: %v\n", res.Result)
	return nil
}

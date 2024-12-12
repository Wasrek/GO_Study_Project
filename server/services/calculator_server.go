package services

import (
	context "context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

//plug and adapter concept return interface

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		//"google.golang.org/grpc/status"
		return nil, status.Errorf(
			codes.InvalidArgument,
			"name is required",
		)
	}

	result := fmt.Sprintf("Hello %v at %v", req.Name, req.CreatedDate.AsTime().Local())
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

func (calculatorServer) Fibonacci(req *FibonacciRequest, stream Calculator_FibonacciServer) error {
	for n := uint32(0); n <= req.N; n++ {
		result := fibo(n)
		res := FibonacciResponse{
			Result: result,
		}
		stream.Send(&res)
		time.Sleep(time.Second)
	}
	return nil
}

func fibo(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibo(n-1) + fibo(n-2)
	}
}

func (calculatorServer) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}
	res := AverageResponse{
		Result: sum / count,
	}
	return stream.SendAndClose(&res)
}

func (calculatorServer) Sum(stream Calculator_SumServer) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		res := SumResponse{
			Result: sum,
		}
		err = stream.Send(&res)
		if err != nil {
			return err 
		}
	}
	return nil
}

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/RaminCH/go3_grpc/task2/server/proto/consignment"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Solve(*pb.Solution) (*pb.Solution, error)
	GetAll() []*pb.Solution
}

//Repository ... Наша база данных
type Repository struct {
	solutions []*pb.Solution
}

//Equation ...
func Equation(a int, b int, c int) int {
  
    var d = b*b - 4*a*c; 
    sqrtVal := sqrt(abs(d))
  
    if d > 0 { 
        x1 := (double)(-b + sqrtVal)/(2*a) 
		v2 := (double)(-b - sqrtVal)/(2*a)
		n_roots := append(count, x1, x2)
		// fmt.Printf("%d", len(count))
    } else if d == 0 { 
		x := b / (2*a)
		n_roots = append(count, x)
		// fmt.Printf("%d", len(count))
	} else // d < 0 
	{ 
        x3 := b / (2*a), sqrtVal 
		x4 := b / (2*a), sqrtVal
		n_roots := append(count, x3, x4)
		// fmt.Printf("%d", len(count))
	} 
	solution := len(n_roots)
	return solution
}

//Solve ....
func (r *Repository) Solve(solution *pb.Solution) (*pb.Solution, error) {
	updatedSolution := append(r.solutions, solution)
	r.solutions = updatedSolution
	return solution, nil
}

//GetAll ...
func (r *Repository) GetAll() []*pb.Solution {
	return r.solutions
}

type service struct {
	repo repository
}

func (s *service) Solve(ctx context.Context, req *pb.) (*pb.Response, error) {
	command, err := s.repo.Solve(req)
	if err != nil {
		return nil, err
	}
	log.Printf("The equation has %v roots", solution)
	return &pb.Response{Solve: true, Solution: solution}, nil
}

//GetAllCommands ...
func (s *service) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	solutions := s.repo.GetAll()
	return &pb.Response{Solutions: solutions}, nil
}

func main() {
	repo := &Repository{}

	//Настройка gRPC сервера
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen port: %v", err)
	}

	server := grpc.NewServer()

	//Регистрируем наш сервис для сервера
	ourService := &service{repo}
	pb.RegisterSolverServer(server, ourService)
	//Чтобы выходные параметры сервера сохранялись в go-runtime
	reflection.Register(server)

	log.Println("gRPC server running on port:", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server from port: %v", err)
	}
}
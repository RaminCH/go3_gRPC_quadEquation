package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/RaminCH/go3_grpc/task2/server/proto/consignment"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"
	"math"
)

const (
	address         = "localhost:50051"
	defaultFilename1 = "test1.json"
	defaultFilename2 = "test2.json"
	defaultFilename3 = "test3.json"
)

func parseJSON(file string) (*pb.Solution, error) {
	var solution *pb.Solution
	fileBody, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(fileBody, &solution)
	return solution, err
}

//Equation ...
func Equation(a int, b int, c int) int {
  
    var d := b*b - 4*a*c; 
    double sqrt_val := sqrt(abs(d)); 
  
    if d > 0 { 
        x1 := (double)(-b + sqrt_val)/(2*a) 
		v2 := (double)(-b - sqrt_val)/(2*a))
		n_roots := append(count, x1, x2)
		// fmt.Printf("%d", len(count))
    } 
    else if d == 0
    { 
		x := (double)b / (2*a))
		n_roots = append(count, x)
		// fmt.Printf("%d", len(count))
    } 
    else // d < 0 
    { 
        x3 := (double)b / (2*a),sqrt_val 
		x4 := (double)b / (2*a), sqrt_val)
		n_roots := append(count, x3, x4)
		// fmt.Printf("%d", len(count))
	} 
	return len(n_roots)
}

func main() {

	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect to port: %v", err)
	}
	defer connection.Close()

	client := pb.NewSolverClient(connection)

	command, err := parseJSON(defaultFilename1)
	if err != nil {
		log.Fatalf("can not parse .json file: %v", err)
	}
	resp, err := client.Solve(context.Background(), solution)
	if err != nil {
		log.Fatalf("can not get response: %v", err)
	}
	log.Printf("Solve: %t", Equation(resp.Solve))
	log.Printf("Solution: %v", resp.Solution)

	getAll, err := client.GetAll(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("can not get response: %v", err)
	}
	for _, v := range getAll.Solutions {
		log.Println(v)
	}

}

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
	log.Printf("Solve: %t", resp.Solve)
	log.Printf("Solution: %v", resp.Solution)

	getAll, err := client.GetAll(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("can not get response: %v", err)
	}
	for _, v := range getAll.Solutions {
		log.Println(v)
	}

}

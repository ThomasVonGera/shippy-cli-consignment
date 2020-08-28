package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/ThomasVonGera/shippy-service-consignment/proto/consignment"
	micro "github.com/micro/go-micro/v2"
)

const (
	adress           = "localhost:50051"
	defaultFilename  = "consignment.json"
	microservicename = "shippy.cli.consignment"
)

func parseFile(filename string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err

}

func main() {
	service := micro.NewService(micro.Name(microservicename))
	service.Init()

	client := pb.NewShippingService("shippy.service.consignment", service.Client())

	filename := defaultFilename

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	consignment, err := parseFile(filename)
	if err != nil {
		log.Fatalf("Datei konnte nicht geparst werden: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Konnte nicht grüßen: %v", err)
	}

	log.Printf("Erzeugt: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Konnte Lieferungen nicht auflisten: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

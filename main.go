package main

import(
	"log"
	"json"
	"context"
	"google.golang.org/grpc"
	pb "github.com/ThomasVonGera/shippy-service-consignment/proto/consignment"
)

const (
	adress = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(filename string) (*pb.Consignment, error){
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err

}

func main() {
	grpcConnection, err := grpc.Dial(adress, grpc.WithInsecure)
	if err != nil{
		log.Fatalf("Verbindung zu %s konnte nicht aufgebaut werden. Fehlermeldung: %v",[adress,err])

	}

	defer grpcConnection.Close()

	client := pb.NewShippingServiceClient(grpcConnection)

	filename := defaultFilename

	if len(os.Args) > 1 {
		filename := os.Args[1]
	}

	consignment, err := parsefile(filename)
	if err != nil{
		log.Fatalf("Datei konnte nicht geparst werden: %v",err)
	}

	r,err := client.CreateConsignment(context.Background(), consignment)
	if err != nil{
		log.Fatalf("Konnte nicht grüßen: %v",err)
	}

	log.Printf("Erzeugt: %t", r.Created)
}
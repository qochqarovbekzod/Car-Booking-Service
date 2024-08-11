package main

import (
	"booking/config"
	"booking/generated/booking"
	"booking/logs"
	"booking/service"
	"booking/storage"
	"booking/storage/mongo"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	logs.InitLogger()

	db, err := mongo.ConnectMongo()
	if err != nil {
		logs.Logger.Error("Failed to connect to MongoDB: ", err)
		log.Fatal(err)
	}

	cfg := config.Load()

	listen, err := net.Listen("tcp", cfg.GRPC_PORT)
	if err!= nil {
        logs.Logger.Error("Failed to listen on port: ", cfg.GRPC_PORT, err)
        log.Fatal(err)
    }
	st := storage.NewProductImpl(db, logs.Logger)
	s := grpc.NewServer()

	booking.RegisterBookingServiceServer(s, service.NewService(logs.Logger, st))
	log.Println("server is running on :8082 ...")
	

}

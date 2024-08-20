package main

import (
	"booking/config"
	"booking/generated/booking"
	"booking/kafka/concumer"
	"booking/logs"
	"booking/service"
	"booking/storage"
	"booking/storage/mongo"
	"booking/storage/redis"
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

	rdb := redis.ConnectR()

	cfg := config.Load()

	listen, err := net.Listen("tcp", cfg.GRPC_PORT)
	if err != nil {
		logs.Logger.Error("Failed to listen on port: ", cfg.GRPC_PORT, err)
		log.Fatal(err)
	}
	st := storage.NewProductImpl(db, logs.Logger, rdb)
	s := grpc.NewServer()

	k := concumer.KafkaStorages{
		Str: service.Service{
			Logger:  logs.Logger,
			Storage: st,
		},
	}
	go func() {
		log.Println("kafka consumer")
		consumer := concumer.NewKafkaConsumer([]string{cfg.KAFKA_BROKERS}, "create-booking", logs.Logger)
		err := consumer.ConsumerMasagae(k.ComsumeMessageCreate)
		if err != nil {
			logs.Logger.Error("Failed to start kafka consumer: ", err)
			log.Fatal(err)
		}
	}()

	go func() {
		log.Println("kafka consumer")
		consumer := concumer.NewKafkaConsumer([]string{cfg.KAFKA_BROKERS}, "update-booking", logs.Logger)
		err := consumer.ConsumerMasagae(k.ComsumeMessageUpdate)
		if err != nil {
			logs.Logger.Error("Failed to start kafka consumer: ", err)
			log.Fatal(err)
		}
	}()

	go func() {
		log.Println("kafka consumer")
		consumer := concumer.NewKafkaConsumer([]string{cfg.KAFKA_BROKERS}, "delete-boking", logs.Logger)
		defer consumer.Close()
		err := consumer.ConsumerMasagae(k.ComsumeMessageDelete)
		if err != nil {
			logs.Logger.Error("Failed to start kafka consumer: ", err)
			log.Fatal(err)
		}
	}()

	go func() {
		log.Println("kafka consumer")
		consumer := concumer.NewKafkaConsumer([]string{cfg.KAFKA_BROKERS}, "create-review", logs.Logger)
		err := consumer.ConsumerMasagae(k.ComsumeMessageCreateReview)
		if err != nil {
			logs.Logger.Error("Failed to start kafka consumer: ", err)
			log.Fatal(err)
		}
	}()

	booking.RegisterBookingServiceServer(s, service.NewService(logs.Logger, st))
	log.Println("server is running on :8082 ...")

	s.Serve(listen)
}

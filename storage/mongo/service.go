package mongo

import (
	pb "booking/generated/booking"
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceRepository interface {
	CreateServices(ctx context.Context, in *pb.CreateServiceRequest) (*pb.Void, error)
	UpdateServices(ctx context.Context, in *pb.UpdateServiceRequest) (*pb.Service, error)
	DeleteServices(ctx context.Context, in *pb.Id) (*pb.Void, error)
	GetAllServices(ctx context.Context, in *pb.GetAllServicesRequest) (*pb.GetAllServicesResponse, error)
	GetByIdServices(ctx context.Context, in *pb.Id) (*pb.Service, error)
}

type ServiceRepo struct {
	DB     *mongo.Database
	Logger *slog.Logger
}

func NewServiceRepo(db *mongo.Database, log *slog.Logger) ServiceRepository {
	return &ServiceRepo{DB: db, Logger: log}
}

func (s *ServiceRepo) CreateServices(ctx context.Context, in *pb.CreateServiceRequest) (*pb.Void, error) {
	collection := s.DB.Collection("services")

	created_at := time.Now().Format("2006/01/02")
	updated_at := time.Now().Format("2006/01/02")
	id := uuid.NewString()
	_, err := collection.InsertOne(ctx, bson.M{
		"_id":         id,
		"name":        in.Name,
		"description": in.Description,
		"price":       in.Price,
		"duration":    in.Duration,
		"created_at":  created_at,
		"updated_at":  updated_at,
		"deleted_at":  0,
	})

	if err != nil {
		s.Logger.Error("Error creating service", err)
		return nil, err
	}
	s.Logger.Info("Service created")
	return &pb.Void{}, nil
}

func (s *ServiceRepo) UpdateServices(ctx context.Context, in *pb.UpdateServiceRequest) (*pb.Service, error) {
	s.Logger.Info("Service updated")
	collection := s.DB.Collection("services")
	updated_at := time.Now().Format("2006/01/02")
	var service pb.Service
	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}

	update := bson.M{"$set": bson.M{
		"name":        in.Name,
		"description": in.Description,
		"price":       in.Price,
		"duration":    in.Duration,
		"updated_at":  updated_at,
	},
	}

	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&service)

	if err != nil {
		s.Logger.Error("Error updating service", err)
		return nil, err
	}

	return &service, nil
}

func (s *ServiceRepo) DeleteServices(ctx context.Context, in *pb.Id) (*pb.Void, error) {
	s.Logger.Info("Deleting services")
	collection := s.DB.Collection("services")
	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}

	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"deleted_at": time.Now().Unix()}})
	if err != nil {
		s.Logger.Error("Error deleting service", err)
		return nil, err
	}

	s.Logger.Info("Service deleted")
	return nil, nil

}

func (s *ServiceRepo) GetByIdServices(ctx context.Context, in *pb.Id) (*pb.Service, error) {
	s.Logger.Info("GetByIdServices method called with ")
	collection := s.DB.Collection("services")
	var service pb.Service
	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}

	err := collection.FindOne(ctx, filter).Decode(&service)
	if err != nil {
		s.Logger.Error("Error getting service by id", err)
		return nil, err
	}
	s.Logger.Info("Service found")
	return &service, nil
}

func (s *ServiceRepo) GetAllServices(ctx context.Context, in *pb.GetAllServicesRequest) (*pb.GetAllServicesResponse, error) {
	s.Logger.Info("GetAllServices request")
	collection := s.DB.Collection("services")

	optionsFind := options.Find()

	optionsFind.SetLimit(int64(in.Limit))
	optionsFind.SetSkip(int64(in.Page))

	fil := bson.D{}

	if in.Name != "" {
		fil = append(fil, bson.E{Key: "name", Value: in.Name})
	}

	if in.Description != "" {
		fil = append(fil, bson.E{Key: "description", Value: in.Description})
	}

	if in.Price != 0 {
		fil = append(fil, bson.E{Key: "price", Value: in.Price})
	}
	if in.Duration != 0 {
		fil = append(fil, bson.E{Key: "duration", Value: in.Duration})
	}

	filter := bson.M{"$and": fil}
	if len(fil) == 0 {
        filter = bson.M{"deleted_at": 0}
    }

	cur, err := collection.Find(ctx, filter, optionsFind)
	if err != nil {
		s.Logger.Error("Error fetching services", err)
		return nil, err
	}

	var services []*pb.Service
	for cur.Next(ctx) {
		var service pb.Service
		err := cur.Decode(&service)
		if err != nil {
			s.Logger.Error("Error decoding service", err)
			return nil, err
		}
		services = append(services, &service)
	}
	return &pb.GetAllServicesResponse{
		Services: services,
		Limit:    in.Limit,
		Page:     int32(in.Page),
	}, nil
}

package mongo

import (
	pb "booking/generated/booking"
	"booking/model"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProviderRepository interface {
	CreateProviders(ctx context.Context, in *pb.CreateProvidersRequest) (*pb.Void, error)
	GetByIdProvider(ctx context.Context, id string) (*pb.Providers, error)
	GetAllProviders(ctx context.Context, in *pb.GetAllProvidersRequest) (*pb.GetAllProviderssResponse, error)
	UpdateProvider(ctx context.Context, in *pb.UpdateProvidersRequest) (*pb.Providers, error)
	DeleteProvider(ctx context.Context, id string) (*pb.Void, error)
}

type ProviderRepo struct {
	DB     *mongo.Database
	Logger *slog.Logger
}

func NewProviderRepo(db *mongo.Database, log *slog.Logger) ProviderRepository {
	return &ProviderRepo{DB: db, Logger: log}
}
func (p ProviderRepo) CreateProviders(ctx context.Context, in *pb.CreateProvidersRequest) (*pb.Void, error) {
	p.Logger.Info("provider create request")
	collection := p.DB.Collection("providers")
	created_at := time.Now().Format("2006/01/02")
	updated_at := time.Now().Format("2006/01/02")
	id := uuid.NewString()
	_, err := collection.InsertOne(ctx, bson.M{
		"_id":            id,
		"user_id":        in.UserId,
		"company_name":   in.CompanyName,
		"description":    in.Description,
		"services":       in.Services,
		"availability":   in.Availability,
		"average_rating": in.AverageRating,
		"location":       in.Location,
		"created_at":     created_at,
		"updated_at":     updated_at,
		"deleted_at": 0,
	})
	if err != nil {
		p.Logger.Error("Error creating provider", err)
		return nil, err
	}
	return &pb.Void{}, nil

}

func (p ProviderRepo) GetByIdProvider(ctx context.Context, id string) (*pb.Providers, error) {
	p.Logger.Info("GetByIdProvider request")
	collection := p.DB.Collection("providers")
	var provider pb.Providers
	filter := bson.M{"$and": []bson.M{
		{"_id": id},
	}}

	err := collection.FindOne(ctx, filter).Decode(&provider)
	if err != nil {
		fmt.Println(id)
		p.Logger.Error("Error getting provider by id", err)
		return nil, err
	}
	return &provider, nil
}

func (p *ProviderRepo) GetAllProviders(ctx context.Context, in *pb.GetAllProvidersRequest) (*pb.GetAllProviderssResponse, error) {
	p.Logger.Info("GetAllProviders request received", "request", in)
	collection := p.DB.Collection("providers")

	findOptions := options.Find()
	findOptions.SetLimit(int64(in.Limit))
	findOptions.SetSkip(int64((in.Page-1) * int64(in.Limit))) 

	filter := bson.M{"deleted_at": 0}

	if in.UserId != "" {
		filter["user_id"] = in.UserId
	}
	if in.CompanyName != "" {
		filter["company_name"] = in.CompanyName
	}
	if in.Description != "" {
		filter["description"] = in.Description
	}
	if in.AverageRating != 0 {
		filter["average_rating"] = in.AverageRating
	}
	if in.Location != "" {
		filter["location"] = in.Location
	}



	var providers []*pb.Providers 

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		p.Logger.Error("Error fetching providers", "error", err)
		return nil, err
	}
	defer cur.Close(ctx) 
	for cur.Next(ctx) {

		var provider model.Provider
		err := cur.Decode(&provider)
		if err != nil {
			p.Logger.Error("Error decoding provider", "error", err)
			return nil, err
		}
		providers = append(providers, convertToProtoProvider(provider))
	}

	if err := cur.Err(); err != nil {
		p.Logger.Error("Cursor error", "error", err)
		return nil, err
	}

	p.Logger.Info("Providers retrieved successfully")
	return &pb.GetAllProviderssResponse{
		Providers: providers,
		Limit:     in.Limit,
		Page:      int32(in.Page),
	}, nil
}

func convertToProtoProvider(provider model.Provider) *pb.Providers {
	return &pb.Providers{
		Id:            provider.Id,
		UserId:        provider.UserId,
		CompanyName:   provider.CompanyName,
		Description:   provider.Description,
		AverageRating: provider.AverageRating,
		Location:      provider.Location,
		CreatedAt:     provider.CreatedAt,
		UpdatedAt:     provider.UpdatedAt,
	}
}
func (p ProviderRepo) UpdateProvider(ctx context.Context, in *pb.UpdateProvidersRequest) (*pb.Providers, error) {
	p.Logger.Info("updating provider")
	collection := p.DB.Collection("providers")
	var provider pb.Providers

	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
	}}

	update := bson.M{
		"$set": bson.M{
			"user_id":        in.UserId,
			"company_name":   in.CompanyName,
			"description":    in.Description,
			"services":       in.Services,
			"availability":   in.Availability,
			"average_rating": in.AverageRating,
			"location":       in.Location,
			"updated_at":     time.Now().Format("2006/01/02"),
		},
	}

	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&provider)

	if err != nil {
		p.Logger.Error("Error updating provider", err)
		return nil, err
	}
	return &provider, nil
}

func (p *ProviderRepo) DeleteProvider(ctx context.Context, id string) (*pb.Void, error) {
	p.Logger.Info("Deleting provider")
	collection := p.DB.Collection("providers")
	filter := bson.M{"$and": []bson.M{
		{"_id": id},
	},
	}
	row, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"deleted_at": time.Now().Unix()}})
	if err != nil {
		p.Logger.Error("Error deleting provider", err)
		return nil, err
	}
	if row.ModifiedCount == 0 {
		return nil, fmt.Errorf("provider not found")
	}
	return &pb.Void{}, nil
}

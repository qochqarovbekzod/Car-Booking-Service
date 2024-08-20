package mongo

import (
	"booking/generated/booking"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BestRepository interface {
	GetBestProvider(ctx context.Context) (*string, error)
	GetBestProviderWithFilter(ctx context.Context, serviceId string) (*booking.Service, error)
}

type Best struct {
	DB *mongo.Database
}

func NewBest(db *mongo.Database) BestRepository {
	return &Best{DB: db}
}

func (b *Best) GetBestProvider(ctx context.Context) (*string, error) {
	collection := b.DB.Collection("bookings")

	pipeline := mongo.Pipeline{
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$service_id"},
				{Key: "orderCount", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "bookingId", Value: bson.D{{Key: "$first", Value: "$_id"}}},
			}},
		},
		{
			{Key: "$sort", Value: bson.D{{Key: "orderCount", Value: -1}}},
		},
		{
			{Key: "$limit", Value: 1},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "booking"},
				{Key: "localField", Value: "bookingId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "bookingDetails"},
			}},
		},
		{
			{Key: "$unwind", Value: "$bookingDetails"},
		},
		{
			{Key: "$replaceRoot", Value: bson.D{
				{Key: "newRoot", Value: "$bookingDetails"},
			}},
		},
	}
	
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	
	var result bson.M
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		fmt.Println(cursor)
		serviceId, ok := result["_id"].(string)
		fmt.Println(serviceId) 
		if !ok {
			return nil, fmt.Errorf("unexpected type for service_id")
		}
		return &serviceId, nil
	}
	return nil, mongo.ErrNoDocuments
}

func (b *Best) GetBestProviderWithFilter(ctx context.Context, serviceId string) (*booking.Service, error) {
	collection := b.DB.Collection("services")
	filter := bson.M{
		"_id":        serviceId,
		"deleted_at": 0,
	}

	var service booking.Service
	err := collection.FindOne(ctx, filter).Decode(&service)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

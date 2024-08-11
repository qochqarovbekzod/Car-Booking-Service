package mongo

import (
	pb "booking/generated/booking"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, in *pb.CreateReviewRequest) (*pb.Void, error)
	UpdatedReview(ctx context.Context, in *pb.UpadateReviewRequest) (*pb.Review, error)
	DeleteReview(ctx context.Context, in *pb.Id) (*pb.Void, error)
	GetAllReviews(ctx context.Context, in *pb.GetAllReviewsRequest) (*pb.GetAllReviewsResponse, error)
	GetByIdReview(ctx context.Context, in *pb.Id) (*pb.Review, error)
}

type ReviewRepo struct {
	DB     *mongo.Database
	Logger *slog.Logger
}

func NewReviewRepo(db *mongo.Database, log *slog.Logger) ReviewRepository {
	return &ReviewRepo{DB: db, Logger: log}
}

func (r *ReviewRepo) CreateReview(ctx context.Context, in *pb.CreateReviewRequest) (*pb.Void, error) {
	r.Logger.Info("CreateReview request received")
	collection := r.DB.Collection("reviews")

	created_at := time.Now().Format("2006/01/02")
	updated_at := time.Now().Format("2006/01/02")
	id := uuid.NewString()

	_, err := collection.InsertOne(ctx, bson.M{
		"_id":         id,
		"booking_id":  in.BookingId,
		"user_id":     in.UserId,
		"provider_id": in.ProviderId,
		"rating":      in.Rating,
		"comment":     in.Comment,
		"created_at":  created_at,
		"updated_at":  updated_at,
		"deleted_at":  0,
	})

	if err != nil {
		r.Logger.Error("Error creating review", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *ReviewRepo) UpdatedReview(ctx context.Context, in *pb.UpadateReviewRequest) (*pb.Review, error) {
	r.Logger.Info("Updated review")
	collection := r.DB.Collection("reviews")
	updated_at := time.Now().Format("2006/01/02")
	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}
	var review pb.Review
	update := bson.M{"$set": bson.M{
		"booking_id":  in.BookingId,
		"user_id":     in.UserId,
		"provider_id": in.ProviderId,
		"rating":      in.Rating,
		"comment":     in.Comment,
		"updated_at":  updated_at,
	}}

	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&review)
	if err != nil {
		r.Logger.Error("Error updating review", err)
		return nil, err
	}
	r.Logger.Info("Updated review")
	return &review, nil
}

func (r *ReviewRepo) DeleteReview(ctx context.Context, in *pb.Id) (*pb.Void, error) {
	r.Logger.Info("Deleting review")
	collection := r.DB.Collection("reviews")
	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}
	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set":bson.M{"updated_at": time.Now().Unix()}})
	if err != nil {
		r.Logger.Error("Error deleting review", err)
		return nil, err
	}
	if res.ModifiedCount == 0 {
		return nil, fmt.Errorf("review not found")
	}
	return &pb.Void{}, nil
}

func (r *ReviewRepo) GetAllReviews(ctx context.Context, in *pb.GetAllReviewsRequest) (*pb.GetAllReviewsResponse, error) {
	r.Logger.Info("GetAllReviews request received")
	collection := r.DB.Collection("reviews")

	optionsFind := options.Find()
	optionsFind.SetLimit(int64(in.Limit))
	optionsFind.SetSkip(in.Page)

	fil := bson.D{}
	if in.BookingId != "" {
		fil = append(fil, bson.E{Key: "booking_id", Value: in.BookingId})
	}

	if in.ProviderId != "" {
		fil = append(fil, bson.E{Key: "provider_id", Value: in.ProviderId})
	}

	if in.Rating != 0 {
		fil = append(fil, bson.E{Key: "rating", Value: in.Rating})
	}

	if in.Comment != "" {
		fil = append(fil, bson.E{Key: "comment", Value: in.Comment})
	}

	filter := bson.M{"$and": fil}
	if len(fil) == 0 {
		filter = bson.M{"deleted_at": 0}
    }
	cur, err := collection.Find(ctx, filter, optionsFind)
	if err != nil {
		r.Logger.Error("Error fetching reviews", err)
		return nil, err
	}
	var reviews []*pb.Review
	for cur.Next(ctx) {
		var review pb.Review
		err := cur.Decode(&review)
		if err != nil {
			r.Logger.Error("Error decoding review", err)
			return nil, err
		}
		reviews = append(reviews, &review)
	}
	r.Logger.Info("Fetched reviews")
	return &pb.GetAllReviewsResponse{
		Reviews: reviews,
		Limit:   in.Limit,
		Page:    in.Page,
	}, nil
}

func (r *ReviewRepo) GetByIdReview(ctx context.Context, in *pb.Id) (*pb.Review, error) {
	r.Logger.Info("GetByIdReview request received")
	collection := r.DB.Collection("reviews")

	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}
	var review pb.Review
	err := collection.FindOne(ctx, filter).Decode(&review)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("review not found")
	} else if err != nil {
		r.Logger.Error("Error fetching review", err)
		return nil, err
	}
	r.Logger.Info("Fetched review")
	return &review, nil
}

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
	var review model.Review
	var reviewss pb.Review
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
	reviewss.Id = review.Id
	reviewss.BookingId = review.BookingId
	reviewss.UserId = review.UserId
	reviewss.ProviderId = review.ProviderId
	reviewss.Rating = review.Rating
	reviewss.Comment = review.Comment
	reviewss.CreatedAt = review.CreatedAt
	reviewss.UpdatedAt = review.UpdatedAt

	r.Logger.Info("Updated review")
	return &reviewss, nil
}

func (r *ReviewRepo) DeleteReview(ctx context.Context, in *pb.Id) (*pb.Void, error) {
	r.Logger.Info("Deleting review")
	collection := r.DB.Collection("reviews")
	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}
	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"deleted_at": time.Now().Unix()}})
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
	r.Logger.Info("GetAllReviews request received", "request", in)
	collection := r.DB.Collection("reviews")

	findOptions := options.Find()
	findOptions.SetLimit(int64(in.Limit))
	findOptions.SetSkip(int64(in.Page * int64(in.Limit))) // Ensure pagination logic is correct

	filter := bson.M{"deleted_at": 0} // Default filter for non-deleted reviews

	if in.BookingId != "" {
		filter["booking_id"] = in.BookingId
	}

	if in.ProviderId != "" {
		filter["provider_id"] = in.ProviderId
	}

	if in.Rating != 0 {
		filter["rating"] = in.Rating
	}

	if in.Comment != "" {
		filter["comment"] = in.Comment
	}

	var reviews []*pb.Review

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		r.Logger.Error("Error fetching reviews", "error", err)
		return nil, err
	}
	defer cur.Close(ctx) 
	
	for cur.Next(ctx) {
		var review model.Review
		err := cur.Decode(&review)
		if err != nil {
			fmt.Println(err)
			r.Logger.Error("Error decoding review", "error", err)
			return nil, err
		}
		reviews = append(reviews, convertToProtoReview(review))
	}

	if err := cur.Err(); err != nil {
		r.Logger.Error("Cursor error", "error", err)
		return nil, err
	}

	r.Logger.Info("Fetched reviews successfully")
	return &pb.GetAllReviewsResponse{
		Reviews: reviews,
		Limit:   in.Limit,
		Page:    in.Page,
	}, nil
}

func convertToProtoReview(review model.Review) *pb.Review {
	return &pb.Review{
		Id:         review.Id,
		BookingId:  review.BookingId,
		UserId:     review.UserId,
		ProviderId: review.ProviderId,
		Rating:     review.Rating,
		Comment:    review.Comment,
		CreatedAt:  review.CreatedAt,
		UpdatedAt:  review.UpdatedAt,
	}
}

func (r *ReviewRepo) GetByIdReview(ctx context.Context, in *pb.Id) (*pb.Review, error) {
	r.Logger.Info("GetByIdReview request received")
	collection := r.DB.Collection("reviews")

	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}
	var review model.Review
	var reviewss pb.Review
	err := collection.FindOne(ctx, filter).Decode(&review)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("review not found")
	} else if err != nil {
		r.Logger.Error("Error fetching review", err)
		return nil, err
	}
	reviewss.Id = review.Id
	reviewss.BookingId = review.BookingId
	reviewss.UserId = review.UserId
	reviewss.ProviderId = review.ProviderId
	reviewss.Rating = review.Rating
	reviewss.Comment = review.Comment
	reviewss.CreatedAt = review.CreatedAt
	reviewss.UpdatedAt = review.UpdatedAt

	r.Logger.Info("Fetched review")
	return &reviewss, nil
}

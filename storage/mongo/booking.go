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

type BookingRepository interface {
	CreateBooking(ctx context.Context, in *pb.CreateBookingRequest) (*pb.Void, error)
	UpdateBooking(ctx context.Context, in *pb.UpdateBookingRequest) (*pb.Booking, error)
	DeleteBooking(ctx context.Context, in *pb.Id) (*pb.Void, error)
	GetByIdBooking(ctx context.Context, in *pb.Id) (*pb.Booking, error)
	GetAllBookings(ctx context.Context, in *pb.GetAllBookingRequest) (*pb.GetAllBookingsResponse, error)
}

type BookingRepo struct {
	DB     *mongo.Database
	Logger *slog.Logger
}

func NewBookingRepo(db *mongo.Database, log *slog.Logger) BookingRepository {
	return &BookingRepo{DB: db, Logger: log}
}

func (b *BookingRepo) CreateBooking(ctx context.Context, in *pb.CreateBookingRequest) (*pb.Void, error) {
	b.Logger.Info("Creating a new booking")
	collection := b.DB.Collection("bookings")

	created_at := time.Now().Format("2006-01-02T15:04:05Z")
	updated_at := time.Now().Format("2006-01-02T15:04:05Z")
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
		"delete_at":      0,
	})

	if err != nil {
		b.Logger.Error("Error creating booking", err)
		return nil, err
	}

	return &pb.Void{}, nil
}

func (b *BookingRepo) GetByIdBooking(ctx context.Context, in *pb.Id) (*pb.Booking, error) {
	b.Logger.Info("GetById called with id: ")
	collection := b.DB.Collection("bookings")
	var booking pb.Booking

	err := collection.FindOne(ctx, bson.M{"_id": in.Id}).Decode(&booking)

	if err != nil {
		b.Logger.Error("Error getting booking by id", err)
		return nil, err
	}

	b.Logger.Info(" booking found")
	return &booking, nil
}

func (b BookingRepo) UpdateBooking(ctx context.Context, in *pb.UpdateBookingRequest) (*pb.Booking, error) {
	b.Logger.Info(" updating booking")
	collection := b.DB.Collection("bookings")
	var booking pb.Booking

	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"delete_at": 0},
	},
	}
	updated_at := time.Now().Format("2006-01-02T15:04:05Z")

	update := bson.M{
		"$set": bson.M{
			"user_id":        in.UserId,
			"company_name":   in.CompanyName,
			"description":    in.Description,
			"services":       in.Services,
			"availability":   in.Availability,
			"average_rating": in.AverageRating,
			"location":       in.Location,
			"updated_at":     updated_at,
		},
	}

	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&booking)

	if err != nil {
		b.Logger.Error("Error updating booking", err)
		return nil, err
	}

	b.Logger.Info("booking updated")
	return &booking, nil
}

func (b BookingRepo) DeleteBooking(ctx context.Context, in *pb.Id) (*pb.Void, error) {
	b.Logger.Info("Deleting booking")
	collection := b.DB.Collection("bookings")

	rews, err := collection.DeleteOne(ctx, bson.M{"_id": in.Id})
	if err != nil {
		b.Logger.Error("Error deleting booking", err)
		return nil, err
	}
	if rews.DeletedCount == 0 {
		b.Logger.Info("Booking not found")
		return nil, nil
	}
	b.Logger.Info("booking deleted")
	return &pb.Void{}, nil
}

func (b *BookingRepo) GetAllBookings(ctx context.Context, in *pb.GetAllBookingRequest) (*pb.GetAllBookingsResponse, error) {
	b.Logger.Info("Booking list")
	collection := b.DB.Collection("bookings")

	findOptions := options.Find()
	findOptions.SetLimit(in.Limit)
	findOptions.SetSkip(in.Page)

	fil := bson.D{}

	if in.CompanyName != "" {
		fil = append(fil, bson.E{Key: "company_name", Value: in.CompanyName})
	}

	if in.Description!=""{
		fil = append(fil, bson.E{Key: "description", Value: in.Description})
	}

	if in.Services!=""{
        fil = append(fil, bson.E{Key: "services", Value: in.Services})
    }

	if in.Availability!=""{
        fil = append(fil, bson.E{Key: "availability", Value: in.Availability})
    }

	if in.AverageRating!=0{
        fil = append(fil, bson.E{Key: "average_rating", Value: in.AverageRating})
    }
	
	filter:=bson.M{"$and": fil}
	if len(fil)==0{
		filter=bson.M{}
	} 


	var bookings []*pb.Booking

	rows, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		b.Logger.Error("Error fetching bookings", err)
		return nil, err
	}

	for rows.Next(ctx) {
		var booking *pb.Booking
		err := rows.Decode(&booking)
		if err != nil {
			b.Logger.Error("Error decoding booking", err)
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	b.Logger.Info("Bookings retrieved successfully")
	return &pb.GetAllBookingsResponse{
		Bookings: bookings,
		Limit:    int32(in.Limit),
		Page:     int32(in.Page),
	}, nil
}

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

	created_at := time.Now().Format("2006/01/02")
	updated_at := time.Now().Format("2006/01/02")
	id := uuid.NewString()

	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "_id", Value: id},
		{Key: "user_id", Value: in.UserId},
		{Key: "provider_id", Value: in.ProviderId},
		{Key: "service_id", Value: in.ServiceId},
		{Key: "status", Value: in.Status},
		{Key: "scheduled_time", Value: in.ScheduledTime},
		{Key: "location", Value: in.Location},
		{Key: "total_price", Value: in.TotalPrice},
		{Key: "created_at", Value: created_at},
		{Key: "updated_at", Value: updated_at},
		{Key: "deleted_at", Value: 0},
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
	var booking model.Booking
	var bookings pb.Booking
	err := collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: in.Id},
		{Key: "deleted_at", Value: 0},
	}).Decode(&booking)
	if err != nil {
		b.Logger.Error("Error getting booking by id", err)
		return nil, err
	}

	b.Logger.Info(" booking found")
	bookings.Id = booking.Id
	bookings.UserId = booking.UserId
	bookings.ProviderId = booking.ProviderId
	bookings.ServiceId = booking.ServiceId
	bookings.Status = booking.Status
	bookings.ScheduledTime = booking.ScheduledTime
	bookings.Location = booking.Location
	bookings.TotalPrice = booking.TotalPrice
	bookings.CreatedAt = booking.CreatedAt
	bookings.UpdatedAt = booking.UpdatedAt
	return &bookings, nil
}

func (b BookingRepo) UpdateBooking(ctx context.Context, in *pb.UpdateBookingRequest) (*pb.Booking, error) {
	b.Logger.Info(" updating booking")
	collection := b.DB.Collection("bookings")
	var booking model.Booking
	var bookings pb.Booking

	filter := bson.M{"$and": []bson.M{
		{"_id": in.Id},
		{"deleted_at": 0},
	},
	}
	updated_at := time.Now().Format("2006/01/02")

	update := bson.M{
		"$set": bson.M{
			"user_id":        in.UserId,
			"provider_id":    in.ProviderId,
			"service_id":     in.ServiceId,
			"status":         in.Status,
			"scheduled_time": in.ScheduledTime,
			"location":       in.Location,
			"total_price":    in.TotalPrice,
			"updated_at":     updated_at,
		},
	}

	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&booking)

	if err != nil {
		b.Logger.Error("Error updating booking", err)
		return nil, err
	}
	booking.Id = bookings.Id
	bookings.UserId = booking.UserId
	bookings.ProviderId = booking.ProviderId
	bookings.ServiceId = booking.ServiceId
	bookings.Status = booking.Status
	bookings.ScheduledTime = booking.ScheduledTime
	bookings.Location = booking.Location
	bookings.TotalPrice = booking.TotalPrice
	bookings.CreatedAt = booking.CreatedAt
	bookings.UpdatedAt = updated_at

	b.Logger.Info("booking updated")
	return &bookings, nil
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
	findOptions.SetSkip(in.Page * in.Limit)

	filter := bson.M{}

	if in.UserId != "" {
		filter["user_id"] = in.UserId
	}
	if in.ProviderId != "" {
		filter["provide_id"] = in.ProviderId
	}
	if in.ServiceId != "" {
		filter["service_id"] = in.ServiceId
	}
	if in.Status != "" {
		filter["status"] = in.Status
	}
	if in.TotalPrice != 0 {
		filter["total_price"] = in.TotalPrice
	}

	var bookings []*pb.Booking

	rows, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		fmt.Println(err)
		b.Logger.Error("Error fetching bookings", err)
		return nil, err
	}
	defer rows.Close(ctx)

	for rows.Next(ctx) {
		var bo model.Booking
		var booking pb.Booking
		err := rows.Decode(&bo)
		if err != nil {
			b.Logger.Error("Error decoding booking", err)
			return nil, err
		}
		booking = convertToProtoBooking(bo)
		bookings = append(bookings, &booking)
	}

	if err := rows.Err(); err != nil {
		b.Logger.Error("Cursor error", err)
		return nil, err
	}

	b.Logger.Info("Bookings retrieved successfully")
	return &pb.GetAllBookingsResponse{
		Bookings: bookings,
		Limit:    int32(in.Limit),
		Page:     int32(in.Page),
	}, nil
}
func convertToProtoBooking(bo model.Booking) pb.Booking {
	return pb.Booking{
		Id:            bo.Id,
		UserId:        bo.UserId,
		ProviderId:    bo.ProviderId,
		ServiceId:     bo.ServiceId,
		Status:        bo.Status,
		ScheduledTime: bo.ScheduledTime,
		Location:      bo.Location,
		TotalPrice:    bo.TotalPrice,
		CreatedAt:     bo.CreatedAt,
		UpdatedAt:     bo.UpdatedAt,
	}
}

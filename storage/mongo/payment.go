package mongo

import (
	pb "booking/generated/booking"
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, in *pb.CreatePaymentRequest) (*pb.Void, error)
	GetByIdPayment(ctx context.Context, id string) (*pb.Payment, error)
	GetAllPayments(ctx context.Context) (*pb.GetAllPaymentsResponse, error)
}

type PaymentRepo struct {
	DB     *mongo.Database
	Logger *slog.Logger
}

func NewPaymentRepo(db *mongo.Database, log *slog.Logger) PaymentRepository {
	return &PaymentRepo{DB: db, Logger: log}
}

func (p *PaymentRepo) CreatePayment(ctx context.Context, in *pb.CreatePaymentRequest) (*pb.Void, error) {
	p.Logger.Info("CreatePayment request")

	collection := p.DB.Collection("payments")
	created_at := time.Now().Format("2006/01/02")
	updated_at := time.Now().Format("2006/01/02")
	id := uuid.NewString()

	_, err := collection.InsertOne(ctx, bson.M{
		"_id":            id,
		"booking_id":     in.BookingId,
		"amount":         in.Amount,
		"status":         in.Status,
		"payment_method": in.PaymentMethod,
		"transaction_id": in.TransactionId,
		"created_at":     created_at,
		"updated_at":     updated_at,
		"deleted_at":     0,
	})

	if err != nil {
		p.Logger.Error("Failed to create payment", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (p *PaymentRepo) GetByIdPayment(ctx context.Context, id string) (*pb.Payment, error) {
	p.Logger.Info("GetByIdPayment called with")
	collection := p.DB.Collection("payments")
	var payment pb.Payment
	filter := bson.M{"$and": []bson.M{
		{"_id": id},
		{"deleted_at": 0},
	}}
	err := collection.FindOne(ctx, filter).Decode(&payment)
	if err != nil {
		p.Logger.Error("Failed to get payment by id", err)
		return nil, err
	}

	p.Logger.Info("Payment received successfully")
	return &payment, nil
}

func (p *PaymentRepo) GetAllPayments(ctx context.Context) (*pb.GetAllPaymentsResponse, error) {
	p.Logger.Info("Payment received successfully")
	collection := p.DB.Collection("payments")
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		p.Logger.Error("Failed to get all payments", err)
		return nil, err
	}
	var payments []*pb.Payment
	for cur.Next(ctx) {
		var payment pb.Payment
		err := cur.Decode(&payment)
		if err != nil {
			p.Logger.Error("Failed to decode payment", err)
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return &pb.GetAllPaymentsResponse{Payments: payments}, nil
}

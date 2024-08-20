package mongo

import (
	pb "booking/generated/booking"
	"booking/model"
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
		"bookingId":     in.BookingId,
		"amount":         in.Amount,
		"status":         in.Status,
		"paymentMethod": in.PaymentMethod,
		"transactionId": in.TransactionId,
		"createdAt":     created_at,
		"updatedAt":     updated_at,
		"deletedAt":     0,
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
	var payment model.Payment
	var payments pb.Payment
	filter := bson.M{"$and": []bson.M{
		{"_id": id},
		{"deleted_at": 0},
	}}
	err := collection.FindOne(ctx, filter).Decode(&payment)
	if err != nil {
		p.Logger.Error("Failed to get payment by id", err)
		return nil, err
	}

	payments.Id = payment.Id
	payments.BookingId = payment.BookingId
	payments.Amount = payment.Amount
	payments.Status = payment.Status
	payments.PaymentMethod = payment.PaymentMethod
	payments.TransactionId = payment.TransactionId
	payments.CreatedAt = payment.CreatedAt
	payments.UpdatedAt = payment.UpdatedAt
	p.Logger.Info("Payment received successfully")
	return &payments, nil
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
		var payment model.Payment
		var paymen pb.Payment
		err := cur.Decode(&payment)
		if err != nil {
			p.Logger.Error("Failed to decode payment", err)
			return nil, err
		}
		paymen.Id = payment.Id
		paymen.BookingId = payment.BookingId
		paymen.Amount = payment.Amount
		paymen.Status = payment.Status
		paymen.PaymentMethod = payment.PaymentMethod
		paymen.TransactionId = payment.TransactionId
		paymen.CreatedAt = payment.CreatedAt
		paymen.UpdatedAt = payment.UpdatedAt
		payments = append(payments, &paymen)
	}
	return &pb.GetAllPaymentsResponse{Payments: payments}, nil
}

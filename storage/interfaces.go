package storage

import (
	m "booking/storage/mongo"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

type Cars interface {
	Booking() m.BookingRepository
	Review() m.ReviewRepository
	Service() m.ServiceRepository
	Payment() m.PaymentRepository
}

type productImpl struct {
	mDB *mongo.Database
	logger *slog.Logger
}

func NewProductImpl(db *mongo.Database,log *slog.Logger) Cars {
	return &productImpl{mDB: db, logger: log}
}

func (p *productImpl) Booking() m.BookingRepository {
    return m.NewBookingRepo(p.mDB, p.logger)
}

func (p *productImpl) Review() m.ReviewRepository {
    return m.NewReviewRepo(p.mDB, p.logger)
}

func (p *productImpl) Service() m.ServiceRepository {
    return m.NewServiceRepo(p.mDB, p.logger)
}

func (p *productImpl) Payment() m.PaymentRepository {
    return m.NewPaymentRepo(p.mDB, p.logger)
}



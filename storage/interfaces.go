package storage

import (
	m "booking/storage/mongo"
	r "booking/storage/redis"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/redis/go-redis/v9"

)

type Cars interface {
	Booking() m.BookingRepository
	Review() m.ReviewRepository
	Service() m.ServiceRepository
	Payment() m.PaymentRepository
	Provider() m.ProviderRepository
	Best() r.ProviderRepository
	BestRepository() m.BestRepository
}

type productImpl struct {
	mDB    *mongo.Database
	rDB    *redis.Client
	logger *slog.Logger
}

func NewProductImpl(db *mongo.Database, log *slog.Logger, rdb *redis.Client) Cars {
	return &productImpl{mDB: db, logger: log, rDB: rdb}
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

func (p *productImpl) Provider() m.ProviderRepository {
	return m.NewProviderRepo(p.mDB, p.logger)
}

func (p *productImpl) Best() r.ProviderRepository {
	return r.NewRedis(p.rDB)
}


func (p *productImpl) BestRepository() m.BestRepository {
    return m.NewBest(p.mDB)
}
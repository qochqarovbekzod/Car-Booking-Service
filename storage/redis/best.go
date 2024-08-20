package redis

import (
	pb "booking/generated/booking"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type ProviderRepository interface {
	CreateAndGet(ctx context.Context, in *pb.Service) (*pb.Service, error)
}

type Redis struct {
	Rdb *redis.Client
}

func NewRedis(rdb *redis.Client) ProviderRepository {
	return &Redis{
		Rdb: rdb,
	}
}

func (r *Redis) CreateAndGet(ctx context.Context, in *pb.Service) (*pb.Service, error) {
	// Marshal the service to JSON
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	// Set the service data in Redis
	err = r.Rdb.Set(ctx, in.Id, b, 0).Err()
	if err != nil {
		return nil, err
	}

	return in, nil
}

package mongo

import (
	pb "booking/generated/booking"
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	mDB, err := ConnectMongo()
	assert.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fmt.Println(logger, "Creating")
	logger.Info("Creating")
	assert.NoError(t, err)

	repo := NewBookingRepo(mDB, logger)

	req := pb.CreateBookingRequest{
        
	}

	res, err := repo.CreateBooking(context.Background(), &req)
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetByIdBooking(t *testing.T)   {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    fmt.Println(logger, "GetById")
    logger.Info("GetById")
    assert.NoError(t, err)

    repo := NewBookingRepo(mDB, logger)

    req := pb.Id{
        Id: "4f81c7fa-fa2b-47c1-a60a-479f5f6336e8",
    }

    res, err := repo.GetByIdBooking(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

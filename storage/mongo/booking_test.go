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
		UserId:        "fask",
		CompanyName:   "test",
		Description:   "test",
		Services:      "test",
		Availability:  "test",
		AverageRating: 1,
		Location:      "test",
	}

	res, err := repo.CreateBooking(context.Background(), &req)
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetByIdBooking(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewBookingRepo(mDB, logger)

    req := pb.Id{
        Id: "205102d6-8a5d-4041-9d25-0c5ba54d7fc5",
    }

    res, err := repo.GetByIdBooking(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestUpadateBooking(t *testing.T){
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewBookingRepo(mDB, logger)

    req := pb.UpdateBookingRequest{
        Id: "205102d6-8a5d-4041-9d25-0c5ba54d7fc5",
        UserId:        "fask",
        CompanyName:   "test",
        Description:   "test",
        Services:      "test",
        Availability:  "test",
        AverageRating: 1,
        Location:      "test",
    }

    res, err := repo.UpdateBooking(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestDeleteBooking(t *testing.T){
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewBookingRepo(mDB, logger)

    req := pb.Id{
        Id: "205102d6-8a5d-4041-9d25-0c5ba54d7fc5",
    }

    res, err := repo.DeleteBooking(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestGetAllBookings(t *testing.T){
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewBookingRepo(mDB, logger)

    req := pb.GetAllBookingRequest{
		Limit: 2,
		Page: 0,
	}

    res, err := repo.GetAllBookings(context.Background(), &req)
    fmt.Println(res,"nmkjokoinjijijijklujhikijok")
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}
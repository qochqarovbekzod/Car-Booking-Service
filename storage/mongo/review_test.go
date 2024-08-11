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

func TestCreateReview(t *testing.T) {

	mDB, err := ConnectMongo()
	assert.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fmt.Println(logger, "Creating")
	logger.Info("Creating")
	assert.NoError(t, err)

	repo := NewReviewRepo(mDB, logger)

	req := pb.CreateReviewRequest{
		BookingId: "test",
		UserId: "test",
		ProviderId: "test",
        Rating:     1,
        Comment:   "test",
	}

	res, err := repo.CreateReview(context.Background(), &req)
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}


func TestGetByIdReview(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewReviewRepo(mDB, logger)

    req := pb.Id{
        Id: "58e4c774-b1d5-4bb8-8cc8-d7e184d3f95d",
    }

    res, err := repo.GetByIdReview(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestGetAllReviews(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewReviewRepo(mDB, logger)

    req := pb.GetAllReviewsRequest{}

    res, err := repo.GetAllReviews(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestUpdatedReview(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewReviewRepo(mDB, logger)

    req := pb.UpadateReviewRequest{
        Id: "58e4c774-b1d5-4bb8-8cc8-d7e184d3f95d",
        BookingId: "test",
        UserId: "test",
        ProviderId: "test",
        Rating:     1,
        Comment:   "test",
    }

    res, err := repo.UpdatedReview(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestDeleteReview(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewReviewRepo(mDB, logger)

    req := pb.Id{
        Id: "58e4c774-b1d5-4bb8-8cc8-d7e184d3f95d",
    }

    res, err := repo.DeleteReview(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

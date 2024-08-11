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

func TestCreatePayment(t *testing.T) {

	mDB, err := ConnectMongo()
	assert.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fmt.Println(logger, "Creating")
	logger.Info("Creating")
	assert.NoError(t, err)

	repo := NewPaymentRepo(mDB, logger)

	req := pb.CreatePaymentRequest{
		BookingId: "booking",
		Amount:     1000,
        Status:    "SUCCESS",
		PaymentMethod: "CREDIT_CARD",
		TransactionId: "1",

	}

	res, err := repo.CreatePayment(context.Background(), &req)
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}


func TestGetByIdPayment(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewPaymentRepo(mDB, logger)

    req := pb.Id{
        Id: "cf5bbe73-a1d9-4edd-a9ab-d73b41226b5c",
    }

    res, err := repo.GetByIdPayment(context.Background(), req.Id)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestGetAllPayments(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    repo := NewPaymentRepo(mDB, logger)


    res, err := repo.GetAllPayments(context.Background())
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}
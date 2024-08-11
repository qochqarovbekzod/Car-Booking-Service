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

func TestCreateService(t *testing.T) {

	mDB, err := ConnectMongo()
	assert.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fmt.Println(logger, "Creating")
	logger.Info("Creating")
	assert.NoError(t, err)

	repo := NewServiceRepo(mDB, logger)

	req := pb.CreateServiceRequest{
		Name:        "test",
		Description: "test",
		Price:       100,
		Duration:    9090,
	}

	res, err := repo.CreateServices(context.Background(), &req)
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUpdateService(t *testing.T) {
	mDB, err := ConnectMongo()
	assert.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fmt.Println(logger, "Updating")
	logger.Info("Updating")
	assert.NoError(t, err)

	repo := NewServiceRepo(mDB, logger)

	req := pb.UpdateServiceRequest{
		Id:          "27e4df56-2b41-41a1-9d5b-e401036dcf39",
		Name:        "test_updated",
		Description: "test_updated",
		Price:       110,
		Duration:    9090,
	}

	res, err := repo.UpdateServices(context.Background(), &req)
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteService(t *testing.T) {
	mDB, err := ConnectMongo()
	assert.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	fmt.Println(logger, "Deleting")
	logger.Info("Deleting")
	assert.NoError(t, err)

	repo := NewServiceRepo(mDB, logger)

	req := pb.Id{
		Id: "55741c97-384b-46ec-b800-907567d945d8",
	}

	res, err := repo.DeleteServices(context.Background(), &req)
	fmt.Println(res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAllServices(t *testing.T){
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    fmt.Println(logger, "Getting all services")
    logger.Info("Getting all services")
    assert.NoError(t, err)

    repo := NewServiceRepo(mDB, logger)

    req := pb.GetAllServicesRequest{}

    res, err := repo.GetAllServices(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestGetServiceById(t *testing.T){
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    fmt.Println(logger, "Getting service by id")
    logger.Info("Getting service by id")
    assert.NoError(t, err)

    repo := NewServiceRepo(mDB, logger)

    req := pb.Id{
        Id: "55741c97-384b-46ec-b800-907567d945d8",
    }

    res, err := repo.GetByIdServices(context.Background(), &req)
    fmt.Println(res)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}
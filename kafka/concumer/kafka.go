package concumer

import (
	pb "booking/generated/booking"
	"booking/logs"
	"booking/service"
	"context"
	"encoding/json"
	"log/slog"
)

type KafkaStorages struct {
	Str service.Service
}

func (k KafkaStorages) ComsumeMessageCreate(message []byte) {
	logs.InitLogger()
	req := pb.CreateBookingRequest{}

	err := json.Unmarshal(message, &req)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	_, err = k.Str.CreateBooking(context.Background(), &req)
	if err != nil {
		logs.Logger.Error(err.Error())
		return
	}

}

func (k KafkaStorages) ComsumeMessageUpdate(message []byte) {
	logs.InitLogger()

	req := pb.UpdateBookingRequest{}

	err := json.Unmarshal(message, &req)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	_, err = k.Str.UpdateBooking(context.Background(), &req)

	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func (k KafkaStorages) ComsumeMessageDelete(message []byte) {
	logs.InitLogger()

	req := pb.Id{}

	err := json.Unmarshal(message, &req)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	_, err = k.Str.DeleteBooking(context.Background(), &req)

	if err != nil {
		slog.Error(err.Error())
		return
	}

}

func (k KafkaStorages) ComsumeMessageCreateReview(message []byte) {
	logs.InitLogger()

	req := pb.CreateReviewRequest{}

	err := json.Unmarshal(message, &req)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	_, err = k.Str.CreateReview(context.Background(), &req)

	if err != nil {
		slog.Error(err.Error())
		return
	}

}

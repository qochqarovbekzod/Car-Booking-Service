package service

import (
	pb "booking/generated/booking"
	"booking/storage"
	"context"
	
	"log/slog"
)

type Service struct {
	pb.UnimplementedBookingServiceServer
	Logger  *slog.Logger
	Storage storage.Cars
}

func NewService(logger *slog.Logger, storage storage.Cars) *Service {
	return &Service{Logger: logger, Storage: storage}
}

func (s *Service) CreateBooking(ctx context.Context, in *pb.CreateBookingRequest) (*pb.Void, error) {

	s.Logger.Info("CreateBooking called with", "", in)
	resp, err := s.Storage.Booking().CreateBooking(ctx, in)
	if err != nil {
		s.Logger.Error("CreateBooking failed", err)
		return nil, err
	}
	s.Logger.Info("Booking created")
	return resp, nil
}

func (s *Service) GetByIdBooking(ctx context.Context, in *pb.Id) (*pb.Booking, error) {

	s.Logger.Info("GetByIdBooking successful")
	resp, err := s.Storage.Booking().GetByIdBooking(ctx, in)
	if err != nil {
		s.Logger.Error("GetByIdBooking failed", err)
		return nil, err
	}
	s.Logger.Info("Book get by id successfully")
	return resp, nil
}

func (s *Service) GetAllBookings(ctx context.Context, in *pb.GetAllBookingRequest) (*pb.GetAllBookingsResponse, error) {
	s.Logger.Info("GetAllBookings successful")
	resp, err := s.Storage.Booking().GetAllBookings(ctx, in)
	if err != nil {
		s.Logger.Error("GetAllBookings failed", err)
		return nil, err
	}
	s.Logger.Info("All books fetched successfully")
	return resp, nil
}

func (s *Service) UpdateBooking(ctx context.Context, in *pb.UpdateBookingRequest) (*pb.Booking, error) {
	s.Logger.Info("UpdateBooking successful")
	resp, err := s.Storage.Booking().UpdateBooking(ctx, in)
	if err != nil {
		s.Logger.Error("UpdateBooking failed", err)
		return nil, err
	}
	s.Logger.Info("Book updated successfully")
	return resp, nil
}

func (s *Service) DeleteBooking(ctx context.Context, in *pb.Id) (*pb.Void, error) {
	s.Logger.Info("DeleteBooking successful")
	resp, err := s.Storage.Booking().DeleteBooking(ctx, in)
	if err != nil {
		s.Logger.Error("DeleteBooking failed", err)
		return nil, err
	}
	s.Logger.Info("Book deleted successfully")
	return resp, nil
}

func (s *Service) CreateService(ctx context.Context, in *pb.CreateServiceRequest) (*pb.Void, error) {
	s.Logger.Info("CreateService request", "request type ", "service type ")
	resp, err := s.Storage.Service().CreateServices(ctx, in)
	if err != nil {
		s.Logger.Error("CreateService failed", err)
		return nil, err
	}
	s.Logger.Info("Service created")
	return resp, nil
}

func (s *Service) GetByIdService(ctx context.Context, in *pb.Id) (*pb.Service, error) {
	s.Logger.Info("GetByIdService request", "request type ", "service type ")
	resp, err := s.Storage.Service().GetByIdServices(ctx, in)
	if err != nil {
		s.Logger.Error("GetByIdService failed", err)
		return nil, err
	}
	s.Logger.Info("Service fetched by id")
	return resp, nil
}

func (s *Service) GetAllServices(ctx context.Context, in *pb.GetAllServicesRequest) (*pb.GetAllServicesResponse, error) {
	s.Logger.Info("GetAllServices request", "request type ", "service type ")
	resp, err := s.Storage.Service().GetAllServices(ctx, in)
	if err != nil {
		s.Logger.Error("GetAllServices failed", err)
		return nil, err
	}
	s.Logger.Info("All services fetched")
	return resp, nil
}

func (s *Service) UpdateService(ctx context.Context, in *pb.UpdateServiceRequest) (*pb.Service, error) {
	s.Logger.Info("UpdateService request", "request type ", "service type ")
	resp, err := s.Storage.Service().UpdateServices(ctx, in)
	if err != nil {
		s.Logger.Error("UpdateService failed", err)
		return nil, err
	}
	s.Logger.Info("Service updated")
	return resp, nil
}

func (s *Service) DeleteService(ctx context.Context, in *pb.Id) (*pb.Void, error) {
	s.Logger.Info("DeleteService request", "request type ", "service type ")
	resp, err := s.Storage.Service().DeleteServices(ctx, in)
	if err != nil {
		s.Logger.Error("DeleteService failed", err)
		return nil, err
	}
	s.Logger.Info("Service deleted")
	return resp, nil
}

func (s Service) CreatePayment(ctx context.Context, in *pb.CreatePaymentRequest) (*pb.Void, error) {
	s.Logger.Info("CreatePayment request received")
	resp, err := s.Storage.Payment().CreatePayment(ctx, in)
	if err != nil {
		s.Logger.Error("CreatePayment failed", err)
		return nil, err
	}
	s.Logger.Info("Payment created successfully")
	return resp, nil
}

func (s Service) GetByIdPayment(ctx context.Context, in *pb.Id) (*pb.Payment, error) {
	s.Logger.Info("GetByIdPayment request received")
	resp, err := s.Storage.Payment().GetByIdPayment(ctx, in.Id)
	if err != nil {
		s.Logger.Error("GetByIdPayment failed", err)
		return nil, err
	}
	s.Logger.Info("Payment fetched successfully")
	return resp, nil
}

func (s *Service) GetAllPayments(ctx context.Context, in *pb.Void) (*pb.GetAllPaymentsResponse, error) {
	s.Logger.Info("GetAllPayments successful with params ")
	resp, err := s.Storage.Payment().GetAllPayments(ctx)
	if err != nil {
		s.Logger.Error("GetAllPayments failed", err)
		return nil, err
	}
	s.Logger.Info("All payments fetched successfully")
	return resp, nil
}

func (s *Service) CreateReview(ctx context.Context, in *pb.CreateReviewRequest) (*pb.Void, error) {
	s.Logger.Info("CreateReview request received successfully")
	resp, err := s.Storage.Review().CreateReview(ctx, in)
	if err != nil {
		s.Logger.Error("CreateReview failed", err)
		return nil, err
	}
	s.Logger.Info("Review created successfully")
	return resp, nil
}

func (s *Service) GetByIdReview(ctx context.Context, in *pb.Id) (*pb.Review, error) {
	s.Logger.Info("GetByIdReview request received")
	resp, err := s.Storage.Review().GetByIdReview(ctx, in)
	if err != nil {
		s.Logger.Error("GetByIdReview failed", err)
		return nil, err
	}
	s.Logger.Info("Review fetched successfully")
	return resp, nil
}

func (s *Service) GetAllReviews(ctx context.Context, in *pb.GetAllReviewsRequest) (*pb.GetAllReviewsResponse, error) {
	s.Logger.Info("GetAllReviews successful with params")
	resp, err := s.Storage.Review().GetAllReviews(ctx, in)
	if err != nil {
		s.Logger.Error("GetAllReviews failed", err)
		return nil, err
	}
	s.Logger.Info("All reviews fetched successfully")
	return resp, nil
}

func (s *Service) UpdateReview(ctx context.Context, in *pb.UpadateReviewRequest) (*pb.Review, error) {
	s.Logger.Info("UpdateReview request received")
	resp, err := s.Storage.Review().UpdatedReview(ctx, in)
	if err != nil {
		s.Logger.Error("UpdateReview failed", err)
		return nil, err
	}
	s.Logger.Info("Review updated successfully")
	return resp, nil
}

func (s *Service) DeleteReview(ctx context.Context, in *pb.Id) (*pb.Void, error) {
	s.Logger.Info("DeleteReview request received")
	resp, err := s.Storage.Review().DeleteReview(ctx, in)
	if err != nil {
		s.Logger.Error("DeleteReview failed", err)
		return nil, err
	}
	s.Logger.Info("Review deleted successfully")
	return resp, nil
}

func (s *Service) CreateProviders(ctx context.Context, in *pb.CreateProvidersRequest) (*pb.Void, error) {
	s.Logger.Info("CreateProvider request received successfully")
	resp, err := s.Storage.Provider().CreateProviders(ctx, in)
	if err != nil {
		s.Logger.Error("CreateProvider failed", err)
		return nil, err
	}
	s.Logger.Info("Provider created successfully")
	return resp, nil
}

func (s *Service) GetByIdProviders(ctx context.Context, in *pb.Id) (*pb.Providers, error) {
	s.Logger.Info("GetByIdProvider request received")
	resp, err := s.Storage.Provider().GetByIdProvider(ctx, in.Id)
	if err != nil {
		s.Logger.Error("GetByIdProvider failed", err)
		return nil, err
	}
	s.Logger.Info("Provider fetched successfully")
	return resp, nil
}

func (s *Service) GetAllProviderss(ctx context.Context, in *pb.GetAllProvidersRequest) (*pb.GetAllProviderssResponse, error) {
	s.Logger.Info("GetAllProviders successful with params")
	resp, err := s.Storage.Provider().GetAllProviders(ctx, in)
	if err != nil {
		s.Logger.Error("GetAllProviders failed", err)
		return nil, err
	}
	s.Logger.Info("All providers fetched successfully")
	return resp, nil
}

func (s *Service) UpdateProviders(ctx context.Context, in *pb.UpdateProvidersRequest) (*pb.Providers, error) {
	s.Logger.Info("UpdateProvider request received")
	resp, err := s.Storage.Provider().UpdateProvider(ctx, in)
	if err != nil {
		s.Logger.Error("UpdateProvider failed", err)
		return nil, err
	}
	s.Logger.Info("Provider updated successfully")
	return resp, nil
}

func (s *Service) DeleteProviders(ctx context.Context, in *pb.Id) (*pb.Void, error) {
	s.Logger.Info("DeleteProvider request received")
	resp, err := s.Storage.Provider().DeleteProvider(ctx, in.Id)
	if err != nil {
		s.Logger.Error("DeleteProvider failed", err)
		return nil, err
	}
	s.Logger.Info("Provider deleted successfully")
	return resp, nil
}

func (s *Service) CreateGet(ctx context.Context, in *pb.Void) (*pb.Service, error) {
	id, err := s.Storage.BestRepository().GetBestProvider(ctx)
	if err != nil {
		s.Logger.Error("Failed to get best provider", err)
		return nil, err
	}

	services, err := s.Storage.BestRepository().GetBestProviderWithFilter(ctx, *id)
	if err != nil {
		s.Logger.Error("Failed to get services", err)
		return nil, err
	}

	s.Logger.Info(" creating service ")
	resp, err := s.Storage.Best().CreateAndGet(ctx, services)
	if err != nil {
		s.Logger.Error("Failed to create service", err)
		return nil, err
	}
	s.Logger.Info("service created successfully")
	return resp, nil
}

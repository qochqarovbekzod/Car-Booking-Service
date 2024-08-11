// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: booking.proto

package booking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*Void, error)
	GetByIdBooking(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Booking, error)
	UpdateBooking(ctx context.Context, in *UpdateBookingRequest, opts ...grpc.CallOption) (*Booking, error)
	DeleteBooking(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error)
	GetAllBookings(ctx context.Context, in *GetAllBookingRequest, opts ...grpc.CallOption) (*GetAllBookingsResponse, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*Void, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*Service, error)
	DeleteService(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error)
	GetAllServices(ctx context.Context, in *GetAllServicesRequest, opts ...grpc.CallOption) (*GetAllServicesResponse, error)
	GetByIdService(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Service, error)
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*Void, error)
	GetByIdPayment(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Payment, error)
	GetAllPayments(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetAllPaymentsResponse, error)
	CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*Void, error)
	UpadateReview(ctx context.Context, in *UpadateReviewRequest, opts ...grpc.CallOption) (*Review, error)
	DeleteReview(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error)
	GetAllReviews(ctx context.Context, in *GetAllReviewsRequest, opts ...grpc.CallOption) (*GetAllReviewsResponse, error)
	GetByIdReview(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Review, error)
	SearchProviders(ctx context.Context, in *SearchProvidersRequest, opts ...grpc.CallOption) (*SearchProvidersResponse, error)
	SearchServices(ctx context.Context, in *SearchServicesRequest, opts ...grpc.CallOption) (*SearchServicesResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/booking.BookingService/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetByIdBooking(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetByIdBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) UpdateBooking(ctx context.Context, in *UpdateBookingRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/booking.BookingService/UpdateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteBooking(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/booking.BookingService/DeleteBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllBookings(ctx context.Context, in *GetAllBookingRequest, opts ...grpc.CallOption) (*GetAllBookingsResponse, error) {
	out := new(GetAllBookingsResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAllBookings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/booking.BookingService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/booking.BookingService/UpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteService(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/booking.BookingService/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllServices(ctx context.Context, in *GetAllServicesRequest, opts ...grpc.CallOption) (*GetAllServicesResponse, error) {
	out := new(GetAllServicesResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAllServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetByIdService(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetByIdService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/booking.BookingService/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetByIdPayment(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetByIdPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllPayments(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetAllPaymentsResponse, error) {
	out := new(GetAllPaymentsResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAllPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/booking.BookingService/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) UpadateReview(ctx context.Context, in *UpadateReviewRequest, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/booking.BookingService/UpadateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteReview(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/booking.BookingService/DeleteReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllReviews(ctx context.Context, in *GetAllReviewsRequest, opts ...grpc.CallOption) (*GetAllReviewsResponse, error) {
	out := new(GetAllReviewsResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAllReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetByIdReview(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetByIdReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) SearchProviders(ctx context.Context, in *SearchProvidersRequest, opts ...grpc.CallOption) (*SearchProvidersResponse, error) {
	out := new(SearchProvidersResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/SearchProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) SearchServices(ctx context.Context, in *SearchServicesRequest, opts ...grpc.CallOption) (*SearchServicesResponse, error) {
	out := new(SearchServicesResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/SearchServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	CreateBooking(context.Context, *CreateBookingRequest) (*Void, error)
	GetByIdBooking(context.Context, *Id) (*Booking, error)
	UpdateBooking(context.Context, *UpdateBookingRequest) (*Booking, error)
	DeleteBooking(context.Context, *Id) (*Void, error)
	GetAllBookings(context.Context, *GetAllBookingRequest) (*GetAllBookingsResponse, error)
	CreateService(context.Context, *CreateServiceRequest) (*Void, error)
	UpdateService(context.Context, *UpdateServiceRequest) (*Service, error)
	DeleteService(context.Context, *Id) (*Void, error)
	GetAllServices(context.Context, *GetAllServicesRequest) (*GetAllServicesResponse, error)
	GetByIdService(context.Context, *Id) (*Service, error)
	CreatePayment(context.Context, *CreatePaymentRequest) (*Void, error)
	GetByIdPayment(context.Context, *Id) (*Payment, error)
	GetAllPayments(context.Context, *Void) (*GetAllPaymentsResponse, error)
	CreateReview(context.Context, *CreateReviewRequest) (*Void, error)
	UpadateReview(context.Context, *UpadateReviewRequest) (*Review, error)
	DeleteReview(context.Context, *Id) (*Void, error)
	GetAllReviews(context.Context, *GetAllReviewsRequest) (*GetAllReviewsResponse, error)
	GetByIdReview(context.Context, *Id) (*Review, error)
	SearchProviders(context.Context, *SearchProvidersRequest) (*SearchProvidersResponse, error)
	SearchServices(context.Context, *SearchServicesRequest) (*SearchServicesResponse, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetByIdBooking(context.Context, *Id) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdBooking not implemented")
}
func (UnimplementedBookingServiceServer) UpdateBooking(context.Context, *UpdateBookingRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBooking not implemented")
}
func (UnimplementedBookingServiceServer) DeleteBooking(context.Context, *Id) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetAllBookings(context.Context, *GetAllBookingRequest) (*GetAllBookingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBookings not implemented")
}
func (UnimplementedBookingServiceServer) CreateService(context.Context, *CreateServiceRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedBookingServiceServer) UpdateService(context.Context, *UpdateServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedBookingServiceServer) DeleteService(context.Context, *Id) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedBookingServiceServer) GetAllServices(context.Context, *GetAllServicesRequest) (*GetAllServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllServices not implemented")
}
func (UnimplementedBookingServiceServer) GetByIdService(context.Context, *Id) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdService not implemented")
}
func (UnimplementedBookingServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedBookingServiceServer) GetByIdPayment(context.Context, *Id) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdPayment not implemented")
}
func (UnimplementedBookingServiceServer) GetAllPayments(context.Context, *Void) (*GetAllPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPayments not implemented")
}
func (UnimplementedBookingServiceServer) CreateReview(context.Context, *CreateReviewRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedBookingServiceServer) UpadateReview(context.Context, *UpadateReviewRequest) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpadateReview not implemented")
}
func (UnimplementedBookingServiceServer) DeleteReview(context.Context, *Id) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
}
func (UnimplementedBookingServiceServer) GetAllReviews(context.Context, *GetAllReviewsRequest) (*GetAllReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReviews not implemented")
}
func (UnimplementedBookingServiceServer) GetByIdReview(context.Context, *Id) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdReview not implemented")
}
func (UnimplementedBookingServiceServer) SearchProviders(context.Context, *SearchProvidersRequest) (*SearchProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProviders not implemented")
}
func (UnimplementedBookingServiceServer) SearchServices(context.Context, *SearchServicesRequest) (*SearchServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchServices not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetByIdBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetByIdBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetByIdBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetByIdBooking(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_UpdateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).UpdateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/UpdateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).UpdateBooking(ctx, req.(*UpdateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/DeleteBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteBooking(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAllBookings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllBookings(ctx, req.(*GetAllBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/UpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteService(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAllServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllServices(ctx, req.(*GetAllServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetByIdService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetByIdService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetByIdService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetByIdService(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetByIdPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetByIdPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetByIdPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetByIdPayment(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAllPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllPayments(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateReview(ctx, req.(*CreateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_UpadateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpadateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).UpadateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/UpadateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).UpadateReview(ctx, req.(*UpadateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteReview(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAllReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllReviews(ctx, req.(*GetAllReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetByIdReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetByIdReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetByIdReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetByIdReview(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_SearchProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).SearchProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/SearchProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).SearchProviders(ctx, req.(*SearchProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_SearchServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).SearchServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/SearchServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).SearchServices(ctx, req.(*SearchServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GetByIdBooking",
			Handler:    _BookingService_GetByIdBooking_Handler,
		},
		{
			MethodName: "UpdateBooking",
			Handler:    _BookingService_UpdateBooking_Handler,
		},
		{
			MethodName: "DeleteBooking",
			Handler:    _BookingService_DeleteBooking_Handler,
		},
		{
			MethodName: "GetAllBookings",
			Handler:    _BookingService_GetAllBookings_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _BookingService_CreateService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _BookingService_UpdateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _BookingService_DeleteService_Handler,
		},
		{
			MethodName: "GetAllServices",
			Handler:    _BookingService_GetAllServices_Handler,
		},
		{
			MethodName: "GetByIdService",
			Handler:    _BookingService_GetByIdService_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _BookingService_CreatePayment_Handler,
		},
		{
			MethodName: "GetByIdPayment",
			Handler:    _BookingService_GetByIdPayment_Handler,
		},
		{
			MethodName: "GetAllPayments",
			Handler:    _BookingService_GetAllPayments_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _BookingService_CreateReview_Handler,
		},
		{
			MethodName: "UpadateReview",
			Handler:    _BookingService_UpadateReview_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _BookingService_DeleteReview_Handler,
		},
		{
			MethodName: "GetAllReviews",
			Handler:    _BookingService_GetAllReviews_Handler,
		},
		{
			MethodName: "GetByIdReview",
			Handler:    _BookingService_GetByIdReview_Handler,
		},
		{
			MethodName: "SearchProviders",
			Handler:    _BookingService_SearchProviders_Handler,
		},
		{
			MethodName: "SearchServices",
			Handler:    _BookingService_SearchServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}

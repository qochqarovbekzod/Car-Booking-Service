package model

type Booking struct {
	Id            string  `bson:"id"`
	UserId        string  `bson:"user_id"`
	ProviderId    string  `bson:"provider_id"`
	ServiceId     string  `bson:"service_id"`
	Status        string  `bson:"status"`
	ScheduledTime string  `bson:"scheduled_time"`
	Location      string  `bson:"location"`
	TotalPrice    float32 `bson:"total_price"`
	CreatedAt     string  `bson:"created_at"`
	UpdatedAt     string  `bson:"updated_at"`
}

type Payment struct {
	Id            string  `bson:"_id"`
	BookingId     string  `bson:"booking_id"`
	Amount        float32 `bson:"amount"`
	Status        string  `bson:"status"`
	PaymentMethod string  `bson:"payment_method"`
	TransactionId string  `bson:"transaction_id"`
	CreatedAt     string  `bson:"created_at"`
	UpdatedAt     string  `bson:"updated_at"`
}

type Review struct {
	Id         string `bson:"_id"`
	BookingId  string `bson:"booking_id"`
	UserId     string `bson:"user_id"`
	ProviderId string `bson:"provider_id"`
	Rating     int32  `bson:"rating"`
	Comment    string `bson:"comment"`
	CreatedAt  string `bson:"created_at"`
	UpdatedAt  string `bson:"updated_at"`
}

type Service struct {
	Id          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float32 `bson:"price"`
	Duration    int64   `bson:"duration"`
	CreatedAt   string  `bson:"created_at"`
	UpdatedAt   string  `bson:"updated_at"`
}

type Provider struct {
	Id            string  `bson:"_id"`
	UserId        string  `bson:"user_id"`
	CompanyName   string  `bson:"company_name"`
	Description   string  `bson:"description"`
	Services      string  `bson:"services"`
	Availability  string  `bson:"availability"`
	AverageRating float32 `bson:"average_rating"`
	Location      string  `bson:"location"`
	CreatedAt     string  `bson:"created_at"`
	UpdatedAt     string  `bson:"updated_at"`
}

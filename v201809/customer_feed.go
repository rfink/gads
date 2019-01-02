package v201809

type CustomerFeedService struct {
	Auth
}

func NewCustomerFeedService(auth *Auth) *CustomerFeedService {
	return &CustomerFeedService{Auth: *auth}
}

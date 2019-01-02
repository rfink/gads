package v201809

type CustomerSyncService struct {
	Auth
}

func NewCustomerSyncService(auth *Auth) *CustomerSyncService {
	return &CustomerSyncService{Auth: *auth}
}

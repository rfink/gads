package v201802

type CustomerSyncService struct {
	Auth
}

func NewCustomerSyncService(auth *Auth) *CustomerSyncService {
	return &CustomerSyncService{Auth: *auth}
}

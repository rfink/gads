package v201802

type TrialService struct {
	Auth
}

func NewTrialService(auth *Auth) *TrialService {
	return &TrialService{Auth: *auth}
}

package v201708

type DraftService struct {
	Auth
}

func NewDraftService(auth *Auth) *DraftService {
	return &DraftService{Auth: *auth}
}

package v201802

type CampaignSharedSetService struct {
	Auth
}

func NewCampaignSharedSetService(auth *Auth) *CampaignSharedSetService {
	return &CampaignSharedSetService{Auth: *auth}
}

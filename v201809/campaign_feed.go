package v201809

type CampaignFeedService struct {
	Auth
}

func NewCampaignFeedService(auth *Auth) *CampaignFeedService {
	return &CampaignFeedService{Auth: *auth}
}

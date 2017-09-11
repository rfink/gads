package v201708

type FeedMappingService struct {
	Auth
}

func NewFeedMappingService(auth *Auth) *FeedMappingService {
	return &FeedMappingService{Auth: *auth}
}

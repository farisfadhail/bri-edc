package resources

type AuthorizationResponse struct {
	Approved bool   `json:"approved"`
	Message  string `json:"message"`
}

func ToAuthorizationResponse(approved bool, message string) *AuthorizationResponse {
	return &AuthorizationResponse{
		Approved: approved,
		Message:  message,
	}
}

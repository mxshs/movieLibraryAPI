package adaptermodels

type AuthRequest struct {
	token string `json:"access_token"`
}

type TokenResponse struct {
	refreshToken string `json:"refresh_token"`
	accessToken  string `json:"access_token"`
}

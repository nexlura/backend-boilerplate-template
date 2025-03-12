package requests

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPassword struct {
	Token                string `json:"token"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type OAuthPayload struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Provider  string `json:"provider"`
	Avatar    string `json:"avatar"`
}

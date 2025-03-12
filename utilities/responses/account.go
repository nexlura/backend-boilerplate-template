package responses

type CreateAccountResponse struct {
	Identity  string `json:"identity"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

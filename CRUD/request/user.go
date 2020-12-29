package request

type SignUpRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
}

func (request SignUpRequest) Validate() bool {
	return !(request.Email == "" || request.Password == "" || request.Name == "")
}

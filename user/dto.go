package user

type InfoResponseBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   bool   `json:"gender"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
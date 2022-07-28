package auth

type InfoResponseBody struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Gender bool   `json:"gender"`
}

func NewInfoResponseBody(m *Member) *InfoResponseBody {
	return &InfoResponseBody{
		Email:  m.email,
		Name:   m.name,
		Gender: m.gender,
	}
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JoinRequestBody struct {
	Email    string
	Password string
	Name     string
	Gender   bool
}

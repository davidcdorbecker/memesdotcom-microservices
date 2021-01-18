package handlers

type authHandler struct {
}

type AuthHandler interface {
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

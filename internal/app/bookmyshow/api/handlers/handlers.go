package handlers

import "github.com/deexithparand/bookmyshow/internal/db/repositories"

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(ur *repositories.UserRepository) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(ur),
	}
}

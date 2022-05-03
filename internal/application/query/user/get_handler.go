package query

import (
	"github.com/nooderg/pipiSpot/internal/config"
	"github.com/nooderg/pipiSpot/internal/core/domain"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"github.com/nooderg/pipiSpot/internal/infrastructure/persistence"
)

type GetUserQueryHandler struct {
	ur ports.UserRepository
	//jwt
}

func (q GetUserQueryHandler) New() GetUserQueryHandler {
	return GetUserQueryHandler{
		persistence.NewUserRepository(config.GetDBClient()),
	}
}

func (q *GetUserQueryHandler) Handle(query GetUserQuery) (*domain.User, error) {
	return q.ur.GetUserByID(query.ID)
}

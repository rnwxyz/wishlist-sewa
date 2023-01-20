package users

import (
	"context"

	"github.com/rnwxyz/wishlist-sewa/dto/resources"
	"github.com/rnwxyz/wishlist-sewa/internal/repositories/users"
)

type userServiceImpl struct {
	userRepository users.UserRepository
}

// GetUserByID implements UserService
func (s *userServiceImpl) FindUserByID(id *uint, ctx context.Context) (*resources.UserResource, error) {
	user, err := s.userRepository.FindUserByID(id, ctx)
	if err != nil {
		return nil, err
	}
	var result resources.UserResource
	result.FromModel(user)
	return &result, err
}

func NewUserService(userRepository users.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

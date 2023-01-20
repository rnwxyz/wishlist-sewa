package auth

import (
	"context"

	"github.com/rnwxyz/wishlist-sewa/config"
	"github.com/rnwxyz/wishlist-sewa/constans"
	"github.com/rnwxyz/wishlist-sewa/dto/requests"
	"github.com/rnwxyz/wishlist-sewa/internal/repositories/users"
	"github.com/rnwxyz/wishlist-sewa/model"
	"github.com/rnwxyz/wishlist-sewa/utils/jwt"
	"github.com/rnwxyz/wishlist-sewa/utils/myerrors"
	"github.com/rnwxyz/wishlist-sewa/utils/password"
)

type authServiceImpl struct {
	userRepository users.UserRepository
	password       password.PasswordHash
	jwtService     jwt.JWTService
}

// CreateDefaultAccount implements AuthService
func (s *authServiceImpl) CreateDefaultAccount() error {
	hashPassword, err := s.password.HashPassword(config.Env.OWNER_PASSWORD)
	if err != nil {
		return err
	}
	owner := model.Owner{
		User: model.User{
			Email:    config.Env.OWNER_EMAIL,
			Password: hashPassword,
			Name:     config.Env.OWNER_NAME,
			Role:     constans.RoleOwner,
		},
		NoWa:        config.Env.OWNER_NO_WA,
		Description: "none",
	}
	err = s.userRepository.DefaultAccount(&owner)
	if err != nil {
		panic(err)
	}
	return nil
}

// Login implements AuthService
func (s *authServiceImpl) Login(login *requests.Credential, ctx context.Context) (*model.Token, error) {
	user, err := s.userRepository.FindUserByEmail(&login.Email, ctx)
	if err != nil {
		return nil, err
	}
	if !s.password.CheckPasswordHash(login.Password, user.Password) {
		return nil, myerrors.ErrInvalidEmailPassword
	}
	var newToken model.Token
	newToken.Token, err = s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	newToken.Role = user.Role
	return &newToken, err
}

// Register implements AuthService
func (s *authServiceImpl) UserRegister(register *requests.UserRegister, ctx context.Context) error {
	hashPassword, err := s.password.HashPassword(register.Password)
	if err != nil {
		return err
	}
	user := register.ToModel()
	user.Password = hashPassword
	user.Role = constans.RoleUser
	_, err = s.userRepository.CreateUser(user, ctx)
	return err
}

func NewAuthService(userRepository users.UserRepository, password password.PasswordHash, jwtService jwt.JWTService) AuthService {
	return &authServiceImpl{
		userRepository: userRepository,
		password:       password,
		jwtService:     jwtService,
	}
}

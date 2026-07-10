package service

import (
	"context"
	"errors"

	"github.com/GforsZi/gin-api/api/internal/model"
	"github.com/GforsZi/gin-api/api/internal/repository"
	firebaseAuth "firebase.google.com/go/v4/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(name, email, password string) (*model.User, error)
	AuthWithFirebase(ctx context.Context, idToken string, authClient *firebaseAuth.Client) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByFirebaseUID(uid string) (*model.User, error)
	GetAll() ([]model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(name, email, password string) (*model.User, error) {
	existing, _ := s.repo.FindByEmail(email)
	if existing != nil && existing.Id != 0 {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) AuthWithFirebase(ctx context.Context, idToken string, authClient *firebaseAuth.Client) (*model.User, error) {
	token, err := authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, errors.New("invalid firebase token")
	}

	uid := token.UID
	email, _ := token.Claims["email"].(string)
	name, _ := token.Claims["name"].(string)

	if name == "" {
		name = email
	}

	existing, err := s.repo.FindByFirebaseUID(uid)
	if err == nil && existing != nil && existing.Id != 0 {
		return existing, nil
	}

	user := &model.User{
		FirebaseUID: uid,
		Name:        name,
		Email:       email,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetByFirebaseUID(uid string) (*model.User, error) {
	return s.repo.FindByFirebaseUID(uid)
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.repo.FindAll()
}

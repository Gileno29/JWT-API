package service

import (
	"errors"
	"jwt-api/internernal/auth"
	"jwt-api/internernal/models"
	"jwt-api/internernal/repository"
	"os"
	"regexp"
)

type UserService struct {
	repository repository.UserInterface
}

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func validEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func validPass(pass string) bool {
	return len(pass) >= 6
}

func NewUserService(repository repository.UserInterface) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Register(user *models.User) error {

	u, _ := s.repository.GetUserByEmail(user.Email)
	if u != nil {
		return errors.New("email já cadastrado")
	}

	if !validEmail(user.Email) {
		return errors.New("o email utilizado não é um email valido")
	}

	if !validPass(user.Pass) {
		return errors.New("a senha precisa ter mais de 6 caracteres")
	}
	return s.repository.CreateUser(user)
}

func (s *UserService) Login(user *models.User) (*string, error) {
	j := auth.NewJWT([]byte(os.Getenv("SECRET_KEY")))
	if !validEmail(user.Email) {
		return nil, errors.New("o email utilizado não é um email valido")
	}
	if !validPass(user.Pass) {
		return nil, errors.New("a senha precisa ter mais de 6 caracteres")
	}
	if err := s.repository.Login(user.Email, user.Pass); err != nil {
		return nil, err
	}

	token, err := j.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repository.GetAll()
}

func (s *UserService) GetUser(id uint64) (*models.User, error) {
	return s.repository.GetUser(id)
}

func (s *UserService) DeleteUser(user *models.User) error {
	if _, err := s.repository.GetUser(user.ID); err != nil {
		return err
	}

	return s.repository.DeleteUser(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	actualUser, err := s.repository.GetUser(user.ID)
	if err != nil {
		return err
	}
	if user.Email == "" {
		user.Email = actualUser.Email
	}

	if user.Name == "" {
		user.Name = actualUser.Name
	}

	if user.Pass == "" {
		user.Pass = actualUser.Pass
	}

	if !validEmail(user.Email) {
		return errors.New("o email utilizado não é um email valido")
	}

	if !validPass(user.Pass) {
		return errors.New("a senha precisa ter mais de 6 caracteres")
	}

	if user.Email != actualUser.Email {
		if u, err := s.repository.GetUserByEmail(user.Email); err != nil {
			return err
		} else if u.ID != 0 {
			return errors.New("email já cadastrado")
		}
	}

	return s.repository.UpdateUser(user)
}

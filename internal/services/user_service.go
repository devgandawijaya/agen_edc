package services

import (
	"agen_edc/internal/models"
	"agen_edc/internal/repositories"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo  *repositories.UserRepository
	jwtSecret string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"` // default to "user"
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewUserService(userRepo *repositories.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *UserService) Register(req RegisterRequest) (*models.User, error) {
	// Check if user already exists
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("user with this email already exists")
	}
	if _, err := s.userRepo.GetByUsername(req.Username); err == nil {
		return nil, errors.New("user with this username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set default role if not provided
	role := req.Role
	if role == "" {
		role = "user"
	}
	if role != "user" && role != "admin" {
		return nil, errors.New("invalid role")
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(req LoginRequest) (string, *models.User, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := s.generateJWT(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *UserService) ForgotPassword(req ForgotPasswordRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Generate reset token (simple implementation, in production use secure random)
	resetToken := s.generateResetToken(user.ID)

	// In a real app, send email with token
	// For now, just return the token

	return resetToken, nil
}

func (s *UserService) ResetPassword(req ResetPasswordRequest) error {
	// Verify token (simple implementation)
	userID, err := s.verifyResetToken(req.Token)
	if err != nil {
		return err
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.userRepo.Update(user)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) GetUsers(limit, offset int) ([]models.User, int64, error) {
	return s.userRepo.GetAll(limit, offset)
}

func (s *UserService) SearchUsers(query string, limit, offset int) ([]models.User, int64, error) {
	return s.userRepo.Search(query, limit, offset)
}

func (s *UserService) UpdateUser(id uint, req UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != "" {
		if req.Role != "user" && req.Role != "admin" {
			return nil, errors.New("invalid role")
		}
		user.Role = req.Role
	}

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) generateJWT(user *models.User) (string, error) {
	claims := JWTClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *UserService) generateResetToken(userID uint) string {
	// Simple token generation, in production use secure method
	return fmt.Sprintf("reset_%d_%d", userID, time.Now().Unix())
}

func (s *UserService) verifyResetToken(token string) (uint, error) {
	// Simple verification, in production parse securely
	var userID uint
	var timestamp int64
	_, err := fmt.Sscanf(token, "reset_%d_%d", &userID, &timestamp)
	if err != nil {
		return 0, errors.New("invalid token")
	}
	// Check if token is not expired (e.g., 1 hour)
	if time.Now().Unix()-timestamp > 3600 {
		return 0, errors.New("token expired")
	}
	return userID, nil
}

package services

import (
	" hery-ciaputra/assignment-05-golang-backend/config"
	" hery-ciaputra/assignment-05-golang-backend/dto"
	" hery-ciaputra/assignment-05-golang-backend/httperror"
	" hery-ciaputra/assignment-05-golang-backend/models"
	" hery-ciaputra/assignment-05-golang-backend/repositories"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthService interface {
	SignIn(*dto.SignInReq) (*dto.TokenResponse, error)
}

type authService struct {
	userRepository repositories.UserRepository
	appConfig      config.AppConfig
}

type AuthSConfig struct {
	UserRepository repositories.UserRepository
	AppConfig      config.AppConfig
}

func NewAuthService(c *AuthSConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
		appConfig:      c.AppConfig,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	User *models.User `json:"user"`
}

func (a *authService) generateJWTToken(user *models.User) (*dto.TokenResponse, error) {
	var idExp = a.appConfig.JWTExpireInMinutes * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp
	timeExpire := jwt.NumericDate{Time: time.Unix(tokenExp, 0)}
	timeNow := jwt.NumericDate{Time: time.Now()}

	claims := &idTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.appConfig.AppName,
			IssuedAt:  &timeNow,
			ExpiresAt: &timeExpire,
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.appConfig.JWTSecret)

	if err != nil {
		return new(dto.TokenResponse), httperror.BadRequestError("BAD_REQUEST", "")
	}
	return &dto.TokenResponse{IDToken: tokenString}, nil
}

func (a *authService) SignIn(req *dto.SignInReq) (*dto.TokenResponse, error) {
	user, err := a.userRepository.MatchingCredential(req.Email)

	errNotMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if errNotMatch != nil || user == nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED",
			Message:    "Unauthorized",
		}
	}
	token, err := a.generateJWTToken(user)
	return token, err
}

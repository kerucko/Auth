package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kerucko/auth/internal/jwtauth"
	"github.com/kerucko/auth/internal/models"
	"github.com/kerucko/auth/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

var (
	errInvalidEmailOrPassword = errors.New("invalid email or password")
	errUserExists             = errors.New("user exists")
	errInvalidAppId           = errors.New("invalid app id")
)

type Auth struct {
	UserSaver       UserSaver
	UserProvider    UserProvider
	AppProvider     AppProvider
	TokenExpiration time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passwordHash []byte) (int64, error)
}

type UserProvider interface {
	FindUser(ctx context.Context, email string) (models.User, error)
}

type AppProvider interface {
	FindApp(ctx context.Context, appId int) (models.App, error)
}

func NewAuth(userSaver UserSaver, userProvider UserProvider, appProvider AppProvider, tokenExpiration time.Duration) *Auth {
	return &Auth{
		UserSaver:       userSaver,
		UserProvider:    userProvider,
		AppProvider:     appProvider,
		TokenExpiration: tokenExpiration,
	}
}

func (a *Auth) Register(ctx context.Context, email string, password string) (int64, error) {
	op := "auth.Register"

	log.Println("Registering user")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	userId, err := a.UserSaver.SaveUser(ctx, email, passwordHash)
	if err != nil {
		log.Printf("%s: %v", op, err)

		if errors.Is(err, storage.ErrUserExists) {
			return 0, fmt.Errorf("%s: %w", op, errUserExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return userId, nil
}

func (a *Auth) Login(ctx context.Context, email string, password string, appId int) (string, error) {
	op := "auth.Login"

	log.Println("Logging in")

	user, err := a.UserProvider.FindUser(ctx, email)
	if err != nil {
		log.Printf("%s: %v", op, err)

		if errors.Is(err, storage.ErrUserNotFound) {
			return "", fmt.Errorf("%s: %w", op, errInvalidEmailOrPassword)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		log.Printf("%s: %v", op, err)
		return "", fmt.Errorf("%s: %w", op, errInvalidEmailOrPassword)
	}

	app, err := a.AppProvider.FindApp(ctx, appId)
	if err != nil {
		log.Printf("%s: %v", op, err)

		if errors.Is(err, storage.ErrAppNotFound) {
			return "", fmt.Errorf("%s: %w", op, errInvalidAppId)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Printf("%s: user logged succesfully", op)

	token, err := jwtauth.NewToken(user, app, a.TokenExpiration)
	if err != nil {
		log.Printf("%s: failed tp generate token; %s", op, err.Error())
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

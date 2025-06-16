package helper

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/Pratam-Kalligudda/user-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/user-service-go/internal/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	secret string
}

func NewAuthHelper(secret string) Auth {
	return Auth{secret: secret}
}

func (a Auth) GenerateHashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func (a Auth) ComparePassword(hashPass, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass)); err != nil {
		return err
	}
	return nil
}

func (a Auth) Authorize(ctx fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh-token")
	log.Print("refresh token :" + refreshToken)
	if _, err := a.VerifyToken(refreshToken); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or expired refresh token",
		})
	}

	authStr := ctx.GetReqHeaders()["Authorization"][0]
	if len(authStr) < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors.New("authorization failed"),
		})
	}
	auth := strings.Split(authStr, " ")
	if len(auth) != 2 || auth[0] != "Bearer" || len(auth[1]) < 1 {
		return errors.New("authorization failed")
	}

	token := auth[1]

	user, err := a.VerifyToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	ctx.Locals("id", user.ID)
	ctx.Locals("email", user.Email)
	ctx.Locals("role", user.UserType)
	log.Printf("email : %s, role : %s, id : %d", user.Email, user.UserType, user.ID)
	ctx.Next()

	return nil
}

func (a Auth) GetCurrentUser(ctx fiber.Ctx) (domain.User, error) {
	user := domain.User{
		Email:    ctx.Locals("email").(string),
		ID:       ctx.Locals("id").(uint),
		UserType: ctx.Locals("role").(string),
	}
	uPrint, _ := json.MarshalIndent(user, "", "\t")
	log.Print("user from getcurrentUser" + string(uPrint))
	return user, nil
}

func (a Auth) GenerateToken(id uint, role, email string, exp time.Duration) (string, error) {
	if id <= 0 || (role != "buyer" && role != "seller") || len(email) == 0 || exp < 15*time.Minute {
		return "", errors.New("one of the parameters is not as expected : ")

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(exp).Unix(),
		"iss":   "user-service",
		"sub":   id,
		"email": email,
		"role":  role,
	})

	signedToken, _ := token.SignedString([]byte(a.secret))
	return signedToken, nil
}

func (a Auth) Validate(usr dto.SignupDTO) error {
	phoneRegex := "^[0-9]{10}$"
	emailRegex := `^[\w\.-]+@[\w\.-]+\.[a-zA-Z]{2,6}$`

	// Validate email
	ok, err := regexp.MatchString(emailRegex, usr.Email)
	if len(usr.Email) == 0 || !ok || err != nil {
		return errors.New("incorrect email format")
	}

	// Validate password length
	if len(usr.Password) < 6 {
		return errors.New("password should be at least 6 characters long")
	}

	// Validate phone
	ok, err = regexp.MatchString(phoneRegex, usr.Phone)
	if len(usr.Phone) != 10 || !ok || err != nil {
		return errors.New("incorrect phone format")
	}

	return nil
}

func (a Auth) VerifyToken(tokenStr string) (domain.User, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.secret), nil
	})

	if err != nil || !token.Valid {
		return domain.User{}, errors.New("failed to parse token or token is not valid")
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	email, _ := claims["email"].(string)
	role, _ := claims["role"].(string)
	id := (uint)(claims["sub"].(float64))

	return domain.User{
		Email:    email,
		UserType: role,
		ID:       id,
	}, nil
}

func (a Auth) GenerateCode() (int, error) {
	return -1, nil
}

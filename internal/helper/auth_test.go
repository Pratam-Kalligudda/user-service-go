package helper

import (
	"strings"
	"testing"
	"time"

	"github.com/Pratam-Kalligudda/user-service-go/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

func TestGenerateHashPassword(t *testing.T) {
	a := NewAuthHelper("testingPass")

	tests := []struct {
		name    string
		pass    string
		wantErr bool
	}{
		{"valid password", "MySecurePass123!", false},
		{"empty password", "", false}, // bcrypt allows empty string
		{"short password", "a", false},
		{"long password 72 bytes", strings.Repeat("a", 72), false},
		{"long password > 72 bytes", strings.Repeat("b", 100), true}, // bcrypt returns error
		{"special characters", "!@#$%^&*()_+=<>?", false},
		{"whitespace only", "    ", false},
		{"unicode characters", "パスワード123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashPass, err := a.GenerateHashPassword(tt.pass)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateHashPassword() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if len(hashPass) == 0 {
					t.Fatal("expected a non-empty hash")
				}

				err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(tt.pass))
				if err != nil {
					t.Fatalf("CompareHashAndPassword() failed: %v", err)
				}
			}
		})
	}
}

func TestGenerateJWT(t *testing.T) {
	a := NewAuthHelper("testing")
	tests := []struct {
		name    string
		id      uint
		email   string
		role    string
		exp     time.Duration
		wantErr bool
	}{
		{"valid token", 1, "user@example.com", "buyer", time.Hour, false},
		{"empty email", 2, "", "seller", time.Hour, true},                             // email should not be empty
		{"empty role", 3, "admin@example.com", "", time.Hour, true},                   // role must be buyer or seller
		{"invalid role", 6, "test@example.com", "admin", time.Hour, true},             // unsupported role
		{"zero ID", 0, "test@example.com", "seller", time.Hour, true},                 // ID must be > 0
		{"short expiry", 4, "test2@example.com", "buyer", time.Millisecond, true},     // too short expiry
		{"negative expiry", 5, "test3@example.com", "seller", -1 * time.Minute, true}, // negative expiry
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := a.GenerateToken(tt.id, tt.role, tt.email, tt.exp)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateToken() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				_, err := a.VerifyToken(token)
				if err != nil {
					t.Fatalf("VerifyToken() failed for valid token: %v", err)
				}
			}
		})
	}
}

func TestValidate(t *testing.T) {
	a := NewAuthHelper("testing")
	tests := []struct {
		name    string
		user    dto.SignupDTO
		wantErr bool
	}{
		{
			name: "valid input",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "user@example.com",
					Password: "secure123",
				},
				Phone: "9876543210",
			},
			wantErr: false,
		},
		{
			name: "invalid email format",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "invalidemail",
					Password: "secure123",
				},
				Phone: "9876543210",
			},
			wantErr: true,
		},
		{
			name: "empty email",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "",
					Password: "secure123",
				},
				Phone: "9876543210",
			},
			wantErr: true,
		},
		{
			name: "short password",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "user@example.com",
					Password: "123",
				},
				Phone: "9876543210",
			},
			wantErr: true,
		},
		{
			name: "empty password",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "user@example.com",
					Password: "",
				},
				Phone: "9876543210",
			},
			wantErr: true,
		},
		{
			name: "invalid phone characters",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "user@example.com",
					Password: "secure123",
				},
				Phone: "98765abcde",
			},
			wantErr: true,
		},
		{
			name: "phone too short",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "user@example.com",
					Password: "secure123",
				},
				Phone: "12345",
			},
			wantErr: true,
		},
		{
			name: "phone too long",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "user@example.com",
					Password: "secure123",
				},
				Phone: "1234567890123",
			},
			wantErr: true,
		},
		{
			name: "empty phone",
			user: dto.SignupDTO{
				LoginDTO: dto.LoginDTO{
					Email:    "user@example.com",
					Password: "secure123",
				},
				Phone: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.Validate(tt.user)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() failed error = %v wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

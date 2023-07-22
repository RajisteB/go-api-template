package account

import (
	"server/utils"
	"time"
)

type NewAccount struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Account struct {
	ID        string    `json:"id"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Country   string    `json:"country"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthAccountResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Token     string    `json:"token"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateAccountType(acct Account, password string) Account {
	now := time.Now()
	return Account{
		ID:        utils.CreateUUID(),
		Firstname: acct.Firstname,
		Lastname:  acct.Lastname,
		Email:     acct.Email,
		Password:  password,
		Country:   acct.Country,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func CreateAuthAccountResponseType(account Account, tkn string) AuthAccountResponse {
	return AuthAccountResponse{
		ID:        account.ID,
		Firstname: account.Firstname,
		Lastname:  account.Lastname,
		Country:   account.Country,
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		Token:     tkn,
	}
}

func CreateAccountResponseType(account Account) AccountResponse {
	return AccountResponse{
		ID:        account.ID,
		Firstname: account.Firstname,
		Lastname:  account.Lastname,
		Country:   account.Country,
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

package merchant

import (
	"server/utils"
	"time"
)

type MerchantAccount struct {
	ID                  string    `json:"id"`
	MerchantName        string    `json:"merchant_name"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	Description         string    `json:"description"`
	SignedMerchantTerms bool      `json:"signed_merchant_terms"`
	Country             string    `json:"country"`
	Rating              float64   `json:"rating"`
	AvatarURL           string    `json:"avatar_url"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type MerchantAccountResponse struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	MerchantName string    `json:"merchant_name"`
	Description  string    `json:"description"`
	Country      string    `json:"country"`
	Rating       float64   `json:"rating"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AuthMerchantAccountResponse struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	MerchantName string    `json:"merchant_name"`
	Description  string    `json:"description"`
	Country      string    `json:"country"`
	Rating       float64   `json:"rating"`
	Token        string    `json:"token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func CreateMerchantAccountType(macct MerchantAccount, password string) MerchantAccount {
	now := time.Now()
	return MerchantAccount{
		ID:                  utils.CreateUUID(),
		MerchantName:        macct.MerchantName,
		Email:               macct.Email,
		Password:            password,
		Description:         macct.Description,
		SignedMerchantTerms: true,
		Country:             macct.Country,
		Rating:              0.0,
		CreatedAt:           now,
		UpdatedAt:           now,
	}
}

func CreateAuthMerchantAccountResponseType(maccount MerchantAccount, tkn string) AuthMerchantAccountResponse {
	return AuthMerchantAccountResponse{
		ID:           maccount.ID,
		MerchantName: maccount.MerchantName,
		Rating:       maccount.Rating,
		Description:  maccount.Description,
		Country:      maccount.Country,
		Email:        maccount.Email,
		CreatedAt:    maccount.CreatedAt,
		UpdatedAt:    maccount.UpdatedAt,
		Token:        tkn,
	}
}

func CreateMerchantAccountResponseType(maccount MerchantAccount) MerchantAccountResponse {
	return MerchantAccountResponse{
		ID:           maccount.ID,
		MerchantName: maccount.MerchantName,
		Rating:       maccount.Rating,
		Description:  maccount.Description,
		Country:      maccount.Country,
		Email:        maccount.Email,
		CreatedAt:    maccount.CreatedAt,
		UpdatedAt:    maccount.UpdatedAt,
	}
}

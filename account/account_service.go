package account

import (
	"errors"
	"server/utils"
)

var (
	accountRepo *AccountRepo
)

type AccountService struct{}

func (service *AccountService) succeededResponse() map[string]bool {
	return map[string]bool{"success": true}
}

func (service *AccountService) failedResponse() map[string]bool {
	return map[string]bool{"success": false}
}

/* ** - ACCOUNTS - ** */

func (service *AccountService) CreateAccount(acct Account) (AuthAccountResponse, error) {

	pw := []byte(acct.Password)

	hash, err := utils.HashPassword(pw)
	if err != nil {
		return AuthAccountResponse{}, err
	}

	account := CreateAccountType(acct, string(hash))

	error := accountRepo.CreateAccount(account)
	if error != nil {
		return AuthAccountResponse{}, error
	}

	token, err := utils.CreateJWT(account.ID)
	if err != nil {
		return AuthAccountResponse{}, err
	}

	response := CreateAuthAccountResponseType(account, token)
	return response, nil
}

func (service *AccountService) LoginAccount(email string, accountPassword string) (AuthAccountResponse, error) {
	account, err := accountRepo.GetAccountByEmail(email)
	if err != nil {
		return AuthAccountResponse{}, err
	}

	isValidPassword := utils.CheckPasswordHash(account.Password, accountPassword)

	if !isValidPassword {
		return AuthAccountResponse{}, errors.New("password is invalid")
	}

	tkn, err := utils.CreateJWT(account.ID)
	if err != nil {
		return AuthAccountResponse{}, err
	}

	response := CreateAuthAccountResponseType(account, tkn)
	return response, nil
}

func (service *AccountService) GetAccountByID(accountID string) (AccountResponse, error) {
	account, err := accountRepo.GetAccountByID(accountID)
	if err != nil {
		return AccountResponse{}, err
	}

	response := CreateAccountResponseType(account)
	return response, nil
}

func (service *AccountService) UpdateAccount(accountReq Account) (map[string]bool, error) {
	err := accountRepo.UpdateAccount(accountReq)
	if err != nil {
		return service.failedResponse(), err
	}

	return service.succeededResponse(), nil
}

func (service *AccountService) UpdateAccountCredentials(accountID string, newCredential string) (map[string]bool, error) {
	pw := []byte(newCredential)

	hash, err := utils.HashPassword(pw)
	if err != nil {
		return service.failedResponse(), err
	}

	error := accountRepo.UpdateAccountCredentials(accountID, hash)
	if error != nil {
		return service.failedResponse(), error
	}

	return service.succeededResponse(), nil
}

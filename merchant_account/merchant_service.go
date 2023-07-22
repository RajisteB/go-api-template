package merchant

import (
	"errors"
	"server/utils"
)

var (
	merchantAccountRepo *MerchantAccountRepo
)

type MerchantAccountService struct{}

func (service *MerchantAccountService) succeededResponse() map[string]bool {
	return map[string]bool{"success": true}
}

func (service *MerchantAccountService) failedResponse() map[string]bool {
	return map[string]bool{"success": false}
}

/* ** - MERCHANT ACCOUNTS - ** */

func (service *MerchantAccountService) CreateMerchantAccount(macct MerchantAccount) (AuthMerchantAccountResponse, error) {
	pw := []byte(macct.Password)

	hash, err := utils.HashPassword(pw)
	if err != nil {
		return AuthMerchantAccountResponse{}, err
	}

	maccount := CreateMerchantAccountType(macct, string(hash))

	error := merchantAccountRepo.CreateMerchantAccount(maccount)
	if error != nil {
		return AuthMerchantAccountResponse{}, error
	}

	token, err := utils.CreateJWT(maccount.ID)
	if err != nil {
		return AuthMerchantAccountResponse{}, err
	}

	response := CreateAuthMerchantAccountResponseType(maccount, token)
	return response, nil
}

func (service *MerchantAccountService) LoginMerchantAccount(email string, accountPassword string) (AuthMerchantAccountResponse, error) {
	account, err := merchantAccountRepo.GetMerchantAccountByEmail(email)
	if err != nil {
		return AuthMerchantAccountResponse{}, err
	}

	isValidPassword := utils.CheckPasswordHash(account.Password, accountPassword)

	if !isValidPassword {
		return AuthMerchantAccountResponse{}, errors.New("password is invalid")
	}

	tkn, err := utils.CreateJWT(account.ID)
	if err != nil {
		return AuthMerchantAccountResponse{}, err
	}

	response := CreateAuthMerchantAccountResponseType(account, tkn)
	return response, nil
}

func (service *MerchantAccountService) GetMerchantAccountByID(macctID string) (MerchantAccountResponse, error) {
	result, err := merchantAccountRepo.GetMerchantAccountByID(macctID)
	if err != nil {
		return MerchantAccountResponse{}, err
	}

	response := CreateMerchantAccountResponseType(result)
	return response, nil
}

func (service *MerchantAccountService) UpdateMerchantAccount(macct MerchantAccount) (map[string]bool, error) {
	err := merchantAccountRepo.UpdateMerchantAccount(macct)
	if err != nil {
		return service.failedResponse(), err
	}
	return service.succeededResponse(), nil
}

func (service *MerchantAccountService) UpdateMerchantAccountCredentials(macctID string, newCredential string) (map[string]bool, error) {
	pw := []byte(newCredential)

	hash, err := utils.HashPassword(pw)
	if err != nil {
		return service.failedResponse(), err
	}
	error := merchantAccountRepo.UpdateMerchantAccountCredentials(macctID, hash)
	if error != nil {
		return service.failedResponse(), error
	}
	return service.succeededResponse(), nil
}

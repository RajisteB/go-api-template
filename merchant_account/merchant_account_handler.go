package merchant

import (
	"encoding/json"
	"net/http"
	"server/utils"

	"github.com/bxcodec/faker/v3"
	"github.com/gorilla/mux"
)

var (
	merchantAccountService *MerchantAccountService
)

type MerchantAccountHandler struct{}

func (mah *MerchantAccountHandler) HandleError(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}

func (mah *MerchantAccountHandler) HandleErrorResponse(w http.ResponseWriter, e error) error {
	return utils.WriteJSON(w, http.StatusBadRequest, mah.HandleError(e))
}

/* ** - MERCHANT ACCOUNTS - ** */

func (mah *MerchantAccountHandler) CreateMerchantAccount(w http.ResponseWriter, r *http.Request) error {
	accountReq := MerchantAccount{}

	if err := json.NewDecoder(r.Body).Decode(&accountReq); err != nil {
		return err
	}

	result, error := merchantAccountService.CreateMerchantAccount(accountReq)

	if error != nil {
		return mah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusCreated, result)
}

func (mah *MerchantAccountHandler) MerchantAccountLogin(w http.ResponseWriter, r *http.Request) error {
	merchantAccountReq := MerchantAccount{}

	if err := json.NewDecoder(r.Body).Decode(&merchantAccountReq); err != nil {
		return err
	}

	result, error := merchantAccountService.LoginMerchantAccount(merchantAccountReq.Email, merchantAccountReq.Password)

	if error != nil {
		return mah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

func (mah *MerchantAccountHandler) GetMerchantAccountByID(w http.ResponseWriter, r *http.Request) error {
	merchantAccountID := mux.Vars(r)["merchant_id"]
	result, error := merchantAccountService.GetMerchantAccountByID(merchantAccountID)

	if error != nil {
		return mah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

func (mah *MerchantAccountHandler) UpdateMerchantAccount(w http.ResponseWriter, r *http.Request) error {
	merchantAccountID := mux.Vars(r)["merchant_id"]
	merchantAccountReq := MerchantAccount{}

	if err := json.NewDecoder(r.Body).Decode(&merchantAccountReq); err != nil {
		return err
	}
	merchantAccountReq.ID = merchantAccountID

	result, error := merchantAccountService.UpdateMerchantAccount(merchantAccountReq)
	if error != nil {
		return mah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

func (mah *MerchantAccountHandler) UpdateMerchantCredentials(w http.ResponseWriter, r *http.Request) error {
	merchantAccountID := mux.Vars(r)["merchant_id"]
	merchantAccountReq := MerchantAccount{}

	if err := json.NewDecoder(r.Body).Decode(&merchantAccountReq); err != nil {
		return err
	}
	merchantAccountReq.ID = merchantAccountID
	result, error := merchantAccountService.UpdateMerchantAccountCredentials(merchantAccountID, merchantAccountReq.Password)

	if error != nil {
		return mah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

/* ** - PRESETS - ** */

func (mah *MerchantAccountHandler) CreatePresetMerchantAccount(w http.ResponseWriter, r *http.Request) error {
	merchantAccountReq := MerchantAccount{
		MerchantName: faker.Username(),
		Email:        faker.Email(),
		Description:  faker.Paragraph(),
		Password:     "merchant123",
		Country:      "FAKER",
	}

	result, error := merchantAccountService.CreateMerchantAccount(merchantAccountReq)

	if error != nil {
		return mah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

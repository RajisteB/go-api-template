package account

import (
	"encoding/json"
	"net/http"
	"server/utils"

	"github.com/bxcodec/faker/v3"
	"github.com/gorilla/mux"
)

var (
	accountService *AccountService
)

type AccountHandler struct{}

/* ** - ACCOUNTS - ** */

func (ah *AccountHandler) HandleError(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}

func (ah *AccountHandler) HandleErrorResponse(w http.ResponseWriter, e error) error {
	return utils.WriteJSON(w, http.StatusBadRequest, ah.HandleError(e))
}

func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	accountReq := Account{}

	if err := json.NewDecoder(r.Body).Decode(&accountReq); err != nil {
		return err
	}

	result, error := accountService.CreateAccount(accountReq)

	if error != nil {
		return ah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusCreated, result)
}

func (ah *AccountHandler) GetAccountByID(w http.ResponseWriter, r *http.Request) error {
	accountID := mux.Vars(r)["id"]
	result, error := accountService.GetAccountByID(accountID)

	if error != nil {
		return ah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

func (ah *AccountHandler) AccountLogin(w http.ResponseWriter, r *http.Request) error {
	accountReq := Account{}

	if err := json.NewDecoder(r.Body).Decode(&accountReq); err != nil {
		return err
	}

	result, error := accountService.LoginAccount(accountReq.Email, accountReq.Password)

	if error != nil {
		return ah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

func (ah *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) error {
	accountID := mux.Vars(r)["id"]
	accountReq := Account{}

	if err := json.NewDecoder(r.Body).Decode(&accountReq); err != nil {
		return err
	}
	accountReq.ID = accountID
	result, error := accountService.UpdateAccount(accountReq)

	if error != nil {
		return ah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

func (ah *AccountHandler) UpdateAccountCredentials(w http.ResponseWriter, r *http.Request) error {
	accountID := mux.Vars(r)["id"]
	accountReq := Account{}

	if err := json.NewDecoder(r.Body).Decode(&accountReq); err != nil {
		return err
	}
	accountReq.ID = accountID
	result, error := accountService.UpdateAccountCredentials(accountID, accountReq.Password)

	if error != nil {
		return ah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

/* ** - PRESETS - ** */

func (ah *AccountHandler) CreatePresetAccount(w http.ResponseWriter, r *http.Request) error {
	first := faker.FirstName()
	last := faker.LastName()
	accountReq := Account{
		Firstname: first,
		Lastname:  last,
		Email:     string(first) + "." + string(last[0]) + "@gmail.com",
		Password:  "testing123",
		Country:   "FAKER",
	}

	result, error := accountService.CreateAccount(accountReq)

	if error != nil {
		return ah.HandleErrorResponse(w, error)
	}

	return utils.WriteJSON(w, http.StatusOK, result)
}

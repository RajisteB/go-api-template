package main

import (
	"log"
	"net/http"
	acct "server/account"
	macct "server/merchant_account"
	mw "server/middleware"
	sqldb "server/sqldb"
	"server/utils"

	"github.com/gorilla/mux"
)

func main() {

	var (
		accountHandler         *acct.AccountHandler
		merchantAccountHandler *macct.MerchantAccountHandler
	)

	listenAddr := ":8080"
	router := mux.NewRouter()
	AccountRouter := router.PathPrefix("/accounts").Subrouter()
	PrivateAccountRouter := router.PathPrefix("/accounts").Subrouter()
	MerchantAccountRouter := router.PathPrefix("/accounts/merchants").Subrouter()
	PrivateMerchantAccountRouter := router.PathPrefix("/accounts/merchants").Subrouter()

	router.Use(mw.Logger)

	PrivateAccountRouter.Use(mw.AuthMiddleware)
	PrivateMerchantAccountRouter.Use(mw.AuthMiddleware)

	if err := sqldb.Connect(); err != nil {
		log.Fatal("error connecting to db", err.Error())
	}

	/* ***API ROUTES*** */

	// Account Routes
	AccountRouter.HandleFunc("/new",
		utils.MakeHTTPHandlerFunc(accountHandler.CreateAccount)).Methods("POST")
	AccountRouter.HandleFunc("/login",
		utils.MakeHTTPHandlerFunc(accountHandler.AccountLogin)).Methods("POST")
	AccountRouter.HandleFunc("/{id}",
		utils.MakeHTTPHandlerFunc(accountHandler.GetAccountByID)).Methods("GET")

	// Private Account Routes
	PrivateAccountRouter.HandleFunc("/{id}",
		utils.MakeHTTPHandlerFunc(accountHandler.UpdateAccount)).Methods("PATCH")
	PrivateAccountRouter.HandleFunc("/{id}/credentials",
		utils.MakeHTTPHandlerFunc(accountHandler.UpdateAccountCredentials)).Methods("PATCH")

	// Merchant Account Routes
	MerchantAccountRouter.HandleFunc("/new",
		utils.MakeHTTPHandlerFunc(merchantAccountHandler.CreateMerchantAccount)).Methods("POST")
	MerchantAccountRouter.HandleFunc("/login",
		utils.MakeHTTPHandlerFunc(merchantAccountHandler.MerchantAccountLogin)).Methods("POST")
	MerchantAccountRouter.HandleFunc("/{merchant_id}",
		utils.MakeHTTPHandlerFunc(merchantAccountHandler.GetMerchantAccountByID)).Methods("GET")

	// Private Merchant Account Routes
	PrivateMerchantAccountRouter.HandleFunc("/{merchant_id}",
		utils.MakeHTTPHandlerFunc(merchantAccountHandler.UpdateMerchantAccount)).Methods("PATCH")
	PrivateMerchantAccountRouter.HandleFunc("/{merchant_id}/credentials",
		utils.MakeHTTPHandlerFunc(merchantAccountHandler.UpdateMerchantCredentials)).Methods("PATCH")

	// Preset Routes (Faker)
	router.HandleFunc("/presets/accounts/new",
		utils.MakeHTTPHandlerFunc(accountHandler.CreatePresetAccount)).Methods("POST")
	router.HandleFunc("/presets/accounts/merchants/new",
		utils.MakeHTTPHandlerFunc(merchantAccountHandler.CreatePresetMerchantAccount)).Methods("POST")

	// Listener
	log.Println("API server running on port", listenAddr)
	http.ListenAndServe(listenAddr, router)

}

package merchant

import (
	"database/sql"
	"errors"
	"fmt"
	"server/sqldb"
)

type MerchantAccountRepo struct{}

/* *** ACCOUNT REPO *** */

func (repo *MerchantAccountRepo) checkIfIDExists(tableName string) func(string) (bool, error) {
	return func(ID string) (bool, error) {
		var count int

		query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id = $1", tableName)
		err := sqldb.DB.QueryRow(query, ID).Scan(&count)
		if err != nil {
			return false, err
		}

		return count > 0, nil
	}

}

func (repo *MerchantAccountRepo) CheckIfMerchantAccountIDExists(ID string) error {
	validate := repo.checkIfIDExists("merchants")
	exists, err := validate(ID)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("merchant account with ID %s not found", ID)
	}

	return nil
}

/* *** MERCHANT ACCOUNT REPO *** */

func (repo *MerchantAccountRepo) CreateMerchantAccount(macct MerchantAccount) error {
	query := `INSERT INTO merchants VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, error := sqldb.DB.Exec(
		query,
		macct.ID,
		macct.MerchantName,
		macct.Email,
		macct.Password,
		macct.Description,
		macct.SignedMerchantTerms,
		macct.Country,
		macct.Rating,
		macct.AvatarURL,
		macct.CreatedAt,
		macct.UpdatedAt,
	)

	if error != nil {
		return error
	}
	return nil
}

func (repo *MerchantAccountRepo) GetMerchantAccountByEmail(email string) (MerchantAccount, error) {

	account := MerchantAccount{}
	query := "SELECT * FROM merchants WHERE email = $1"
	row := sqldb.DB.QueryRow(
		query,
		email,
	)

	err := row.Scan(
		&account.ID,
		&account.MerchantName,
		&account.Email,
		&account.Password,
		&account.Description,
		&account.SignedMerchantTerms,
		&account.Country,
		&account.Rating,
		&account.AvatarURL,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return MerchantAccount{}, fmt.Errorf("user with email %s not found", email)
		}
		return MerchantAccount{}, err
	}

	return account, nil
}

func (repo *MerchantAccountRepo) GetMerchantAccountByID(macctID string) (MerchantAccount, error) {
	error := repo.CheckIfMerchantAccountIDExists(macctID)
	if error != nil {
		return MerchantAccount{}, errors.New(error.Error())
	}

	maccount := MerchantAccount{}
	query := "SELECT * FROM merchants WHERE id = $1"
	rows, err := sqldb.DB.Query(
		query,
		macctID,
	)

	if err != nil {
		return MerchantAccount{}, err
	}

	for rows.Next() {
		err = rows.Scan(
			&maccount.ID,
			&maccount.MerchantName,
			&maccount.Email,
			&maccount.Password,
			&maccount.Description,
			&maccount.SignedMerchantTerms,
			&maccount.Country,
			&maccount.Rating,
			&maccount.AvatarURL,
			&maccount.CreatedAt,
			&maccount.UpdatedAt,
		)

		if err != nil {
			return MerchantAccount{}, err
		}
	}

	return maccount, nil

}

func (repo *MerchantAccountRepo) UpdateMerchantAccount(macct MerchantAccount) error {
	error := repo.CheckIfMerchantAccountIDExists(macct.ID)
	if error != nil {
		return errors.New(error.Error())
	}

	fieldsToUpdate := 0
	query := "UPDATE merchants SET "

	if macct.MerchantName != "" {
		query += fmt.Sprintf("merchant_name = '%s', ", macct.MerchantName)
		fieldsToUpdate++
	}

	if macct.Description != "" {
		query += fmt.Sprintf("description = '%s', ", macct.Description)
		fieldsToUpdate++
	}

	if macct.Email != "" {
		query += fmt.Sprintf("email = '%s', ", macct.Email)
		fieldsToUpdate++
	}

	if macct.Country != "" {
		query += fmt.Sprintf("country = '%s', ", macct.Country)
		fieldsToUpdate++
	}

	if fieldsToUpdate == 0 {
		return fmt.Errorf("error: no merchant fields to update")
	}

	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = '%s';", macct.ID)

	result, err := sqldb.DB.Exec(query)

	if result != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *MerchantAccountRepo) UpdateMerchantAccountCredentials(merchantAccountID string, newCredential []byte) error {
	error := repo.CheckIfMerchantAccountIDExists(merchantAccountID)
	if error != nil {
		return errors.New(error.Error())
	}

	query := "UPDATE merchants SET password = $1 WHERE id = $2"

	_, err := sqldb.DB.Exec(
		query,
		newCredential,
		merchantAccountID,
	)

	if err != nil {
		return err
	}

	return nil
}

package account

import (
	"database/sql"
	"errors"
	"fmt"
	"server/sqldb"
)

type AccountRepo struct{}

/* *** ACCOUNT REPO *** */

func (repo *AccountRepo) checkIfIDExists(tableName string) func(string) (bool, error) {
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

func (repo *AccountRepo) CheckIfAccountIDExists(ID string) error {
	validate := repo.checkIfIDExists("accounts")
	exists, err := validate(ID)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("account with ID %s not found", ID)
	}

	return nil
}

func (repo *AccountRepo) CreateAccount(acct Account) error {
	query := `INSERT INTO accounts VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, error := sqldb.DB.Exec(
		query,
		acct.ID,
		acct.Firstname,
		acct.Lastname,
		acct.Email,
		acct.Password,
		acct.Country,
		acct.AvatarURL,
		acct.CreatedAt,
		acct.UpdatedAt,
	)

	if error != nil {
		return error
	}
	return nil
}

func (repo *AccountRepo) GetAccountByEmail(email string) (Account, error) {

	account := Account{}
	query := "SELECT * FROM accounts WHERE email = $1"
	row := sqldb.DB.QueryRow(
		query,
		email,
	)

	err := row.Scan(
		&account.ID,
		&account.Firstname,
		&account.Lastname,
		&account.Email,
		&account.Password,
		&account.Country,
		&account.AvatarURL,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Account{}, fmt.Errorf("user with email %s not found", email)
		}
		return Account{}, err
	}

	return account, nil
}

func (repo *AccountRepo) GetAccountByID(accountID string) (Account, error) {

	account := Account{}
	query := "SELECT * FROM accounts WHERE id = $1"
	rows, err := sqldb.DB.Query(
		query,
		accountID,
	)

	if err != nil {
		return Account{}, err
	}

	for rows.Next() {
		err = rows.Scan(
			&account.ID,
			&account.Firstname,
			&account.Lastname,
			&account.Email,
			&account.Password,
			&account.Country,
			&account.AvatarURL,
			&account.CreatedAt,
			&account.UpdatedAt,
		)

		if err != nil {
			return Account{}, err
		}
	}

	return account, nil

}

func (repo *AccountRepo) UpdateAccount(accountReq Account) error {
	error := repo.CheckIfAccountIDExists(accountReq.ID)
	if error != nil {
		return errors.New(error.Error())
	}

	fieldsToUpdate := 0
	query := "UPDATE accounts SET "

	if accountReq.Firstname != "" {
		query += fmt.Sprintf("first_name = '%s', ", accountReq.Firstname)
		fieldsToUpdate++
	}

	if accountReq.Lastname != "" {
		query += fmt.Sprintf("last_name = '%s', ", accountReq.Lastname)
		fieldsToUpdate++
	}

	if accountReq.Email != "" {
		query += fmt.Sprintf("email = '%s', ", accountReq.Email)
		fieldsToUpdate++
	}

	if accountReq.Country != "" {
		query += fmt.Sprintf("country = '%s', ", accountReq.Country)
		fieldsToUpdate++
	}

	if fieldsToUpdate == 0 {
		return fmt.Errorf("error: no account fields to update")
	}

	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = '%s';", accountReq.ID)

	_, err := sqldb.DB.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (repo *AccountRepo) UpdateAccountCredentials(accountID string, newCredential []byte) error {
	error := repo.CheckIfAccountIDExists(accountID)
	if error != nil {
		return errors.New(error.Error())
	}

	query := "UPDATE accounts SET password = $1 WHERE id = $2"

	_, err := sqldb.DB.Exec(
		query,
		newCredential,
		accountID,
	)

	if err != nil {
		return err
	}

	return nil
}

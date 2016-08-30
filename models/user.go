package models

import (
	"github.com/vikramjakhr/vikramjakhar.com/config"
	"log"
	"github.com/vikramjakhr/vikramjakhar.com/util"
)

type Email string
type UserName string
type Password string

type RegistrationInfo struct {
	Email           Email
	UserName        UserName
	FirstName       string
	LastName        string
	Password        Password
	ConfirmPassword Password
}

func FindByEmail(email string) (User, error) {
	query := "SELECT * FROM user where email=?"
	log.Println(query)
	rows, err := config.DB.Query(query, email)
	user := User{}
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.UserName,
			&user.Email, &user.ImageUrl, &user.Mobile, &user.Address,
			&user.IsBlocked, &user.IsAdmin, &user.IsAccountVerified,
			&user.DateCreated, &user.LastUpdated)
		log.Println(rows.Columns())
		if err != nil {
			return user, err
		}

	}
	return user, nil
}

func (regInfo *RegistrationInfo) Register() {
	//TODO : prepared stmt are for repeatable use, not for a single time use, it reduces the performance because
	//TODO : first it's sending sql to server to prepare for it with placeholder, and server respond with stmt id,
	//TODO : and then executing query, and closing stmt which do 3 round trips to DB.
	//TODO : user simple query instead if you are using it once
	stmt, err := config.DB.Prepare("INSERT user SET email=?,userName=?,firstName=?,lastName=?,password=?")
	checkErr(err)
	rs, err := stmt.Exec(regInfo.Email, regInfo.UserName, regInfo.FirstName, regInfo.LastName, regInfo.Password)
	checkErr(err)
	log.Println(rs.LastInsertId())
	stmt.Close();
}

func (regInfo *RegistrationInfo) Validate() {
	regInfo.Email.Validate()
	if regInfo.Password == regInfo.ConfirmPassword {
		regInfo.Password.Validate()
	}
	regInfo.UserName.Validate()
}

func (e *Email) Validate() bool {
	return util.EmailRegex.MatchString((string)(*e))
}

func (p *Password) Validate() bool {
	return util.PasswordRegex.MatchString((string)(*p))
}

func (u *UserName) Validate() bool {
	return util.UserNameRegex.MatchString((string)(*u))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
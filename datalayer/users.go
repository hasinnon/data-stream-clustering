package datalayer

import (
	"database/sql"
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
)

// UserInfo stores information about a user
type UserInfo struct {
	ID     sql.NullInt32 // int
	Fname  sql.NullString
	Lname  sql.NullString
	Email  sql.NullString
	Credit sql.NullInt32
}

// UserPassword stores password of user
type UserPassword struct {
	Password sql.NullString
}

// UserLogin stores information about user that send to jwt
type UserLogin struct {
	ID       uint
	Fname    string
	Lname    string
	Password string
}

// Is===Valid check rigister information is valid
func (user *UserInfo) IsValid() error {
	err1 := validation.Validate(user.Fname, validation.Required.Error("نام کاربر نمی تواند خالی باشد"))
	if err1 != nil {
		return err1
	}
	err2 := validation.Validate(user.Lname, validation.Required.Error("نام خانوادگی کاربر نمی تواند خالی باشد"))
	if err2 != nil {
		return err2
	}
	err3 := err2 //validation.Validate(user.Email, is.Email)
	if err3 != nil {
		err3.Error()
		return err3
	}
	return nil

}

// GetUserByID returns a user with the given national ID.
func (handler *MyDB) GetUserByID(userID int) (UserInfo, error) {
	row := handler.db.QueryRow("select uid,fname,lname,email,credit from clustream_sru.users where uid=?", userID)

	return getRowDataUser(row)
}

// ***NOT USED***
func (handler *MyDB) GetUsers() ([]UserInfo, error) {
	rows, err := handler.db.Query("select uid,fname,lname,email,credit from clustream_sru.users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []UserInfo{}
	for rows.Next() {
		u, err := getRowsDataUser(rows)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// I++++++++++nsertUser insert a doctor to database and returns err.
func (handler *MyDB) InsertUser(per UserInfo, pass UserPassword) error {
	var err error
	if err != nil {
		return err
	}
	return nil
}

func getRowDataUser(row *sql.Row) (UserInfo, error) {
	u := UserInfo{}

	err := row.Scan(&u.ID, &u.Fname, &u.Lname, &u.Email, &u.Credit)
	if err != nil {
		return u, err
	}
	return u, nil
}

// ***NOT USED***
func getRowsDataUser(row *sql.Rows) (UserInfo, error) {
	u := UserInfo{}

	err := row.Scan(&u.ID, &u.Fname, &u.Lname, &u.Email, &u.Credit)
	if err != nil {
		return u, err
	}
	return u, nil
}

//LOGIN--------------------------------------------------------------

// GetUserLogin returns a Login information with the given email
func (handler *MyDB) GetUserLogin(userEmail string) (UserLogin, error) {
	row := handler.db.QueryRow("select uid,fname,lname,password from clustream_sru.users where uid=?", userEmail)

	return getRowDataUserLogin(row)
}

// GetUserByEmail returns a user with the given Email.
func (handler *MyDB) GetUserByEmail(userEmail string) (UserLogin, error) {
	row := handler.db.QueryRow("select uid,fname,lname,password from clustream_sru.users where email=?", userEmail)

	return getRowDataUserLogin(row)
}

func getRowDataUserLogin(row *sql.Row) (UserLogin, error) {
	u := UserLogin{}

	err := row.Scan(&u.ID, &u.Fname, &u.Lname, &u.Password)
	if err != nil {
		return u, err
	}
	return u, nil
}

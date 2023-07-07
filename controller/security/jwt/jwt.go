package jwt

import (
	"my/ar/399/datastream/datalayer"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
)

var jwtKey = []byte("my_Secret_Key")

// LoginForm stores login form information
type LoginForm struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// UserClaims srores inforamtion about user jwt
type UserClaims struct {
	UID   uint
	Fname string
	Lname string
	jwt.StandardClaims
}

// IsValid check login information is valid
func (l *LoginForm) IsValid() error {
	err1 := validation.Validate(l.Username, validation.Required.Error("نام کاربری نمی تواند خالی باشد"))
	if err1 != nil {
		return err1
	}

	err2 := validation.Validate(l.Password, validation.Required.Error("رمز عبور نمی تواند خالی باشد"))
	if err2 != nil {
		return err2
	}

	return nil

}

// UserLogin returns true if user logined corectlly ***
func UserLogin(w http.ResponseWriter, login LoginForm, dbhandler datalayer.MyDB) bool {
	user, err := dbhandler.GetUserByEmail(login.Username)

	if err != nil {
		return false
	}

	if user.Password != login.Password {
		return false
	}

	expTime := time.Now().Add(30 * time.Minute)

	cliams := &UserClaims{
		UID:   user.ID,
		Fname: user.Fname,
		Lname: user.Lname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	stringToken, err := token.SignedString(jwtKey)

	if err != nil {
		return false
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "JWTToken",
		Expires:  expTime,
		Value:    stringToken,
		HttpOnly: true,
	})
	return true
}

// IsLogedin returns true if user logined ghablan ***
func IsLogedin(r *http.Request) (datalayer.UserLogin, bool) {

	c, err := r.Cookie("JWTToken")
	if err != nil {
		return datalayer.UserLogin{}, false
	}

	tokenString := c.Value

	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		return datalayer.UserLogin{}, false
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return datalayer.UserLogin{}, false
		}
		return datalayer.UserLogin{}, false
	}

	return datalayer.UserLogin{ID: claims.UID, Fname: claims.Fname, Lname: claims.Lname}, true
}

// ServiceClaims srores inforamtion about service jwt
type ServiceClaims struct {
	SID int32
	Key string
	jwt.StandardClaims
}

// IsValidServiceRequest returns true if sevice request is valid
func IsValidServiceRequest(r *http.Request) (datalayer.ServiceLogin, bool) {

	c := r.Header.Get("ApiKey")
	// if err != nil {
	// 	return datalayer.UserLogin{}, false
	// }

	tokenString := c

	claims := &ServiceClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		return datalayer.ServiceLogin{}, false
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return datalayer.ServiceLogin{}, false
		}
		return datalayer.ServiceLogin{}, false
	}

	return datalayer.ServiceLogin{SID: claims.SID, Key: claims.Key}, true
}

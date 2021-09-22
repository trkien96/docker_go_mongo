package controllers

import (
	"encoding/json"
	"errors"
	"go_mongo/constant"
	"go_mongo/models"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

func ShowPage404(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./views/404.html")
	tmpl.Execute(w, nil)
}

func CreatePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./views/create.html")
	tmpl.Execute(w, nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./views/login.html")
	tmpl.Execute(w, nil)
}

func GenToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(3600 * time.Second)
	claims := &Claims{
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func CheckValidToken(tokenHeader string) (*Claims, error) {
	if strings.Contains(tokenHeader, "Bearer ") {
		jwtStrings := strings.Split(tokenHeader, "Bearer ")
		if len(jwtStrings) > 1 {
			tokenHeader = jwtStrings[1]
		}
	}
	if len(tokenHeader) == 0 {
		return nil, errors.New(constant.ERR_TOKEN_IS_INVALID)
	}

	token, err := jwt.ParseWithClaims(tokenHeader, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if ok {
		return claims, nil
	}

	return nil, errors.New(constant.ERR_TOKEN_IS_INVALID)
}

func ResponseErr(w http.ResponseWriter, statusCode int, message string) {
	if message == "" {
		message = http.StatusText(statusCode)
	}

	response, err := json.Marshal(models.Error{
		Status:  http.StatusInternalServerError,
		Message: message,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func ResponseOk(w http.ResponseWriter, data interface{}) {
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

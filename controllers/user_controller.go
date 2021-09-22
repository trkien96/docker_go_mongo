package controllers

import (
	"encoding/json"
	"go_mongo/constant"
	"go_mongo/interfaces"
	"go_mongo/models"
	"html/template"
	"net/http"
)

var (
	jwtKey      = []byte(constant.JWT_KEY)
	user        models.User
	tokenString string
)

type UserControler struct {
	Handler interfaces.UserInterface
}

func (c *UserControler) Create(w http.ResponseWriter, r *http.Request) {
	var regData models.RegistrationData
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		ResponseErr(w, http.StatusBadRequest, "")
		return
	}
	_, err = c.Handler.FindOne(map[string]interface{}{"email": regData.Email})
	if err != nil && err != constant.ERR_USER_NOT_FOUND {
		ResponseErr(w, http.StatusConflict, constant.ERR_USER_IS_EXIST)
		return
	}

	user = models.User{
		Email:    regData.Email,
		Password: regData.Password,
		Name:     regData.Name,
	}
	_, err = c.Handler.Insert(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError, constant.ERR_INSERT_USER)
		return
	}

	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError, constant.ERR_GEN_TOKEN)
		return
	}

	//auth := models.Auth{
	//	Token:   tokenString,
	//	LoginAt: time.Now(),
	//}
	//authController := AuthController{}
	//authController.Handler.Insert(auth)

	ResponseOk(w, models.Success{
		Token:  tokenString,
		Status: http.StatusOK,
		Url:    "/user/login",
	})
}

func (c *UserControler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData models.LoginData
	var err error
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		err = json.NewDecoder(r.Body).Decode(&loginData)
		if err != nil {
			ResponseErr(w, http.StatusBadRequest, "")
			return
		}

		filter := map[string]interface{}{"email": loginData.Email, "password": loginData.Password}
		user, err = c.Handler.FindOne(filter)
		if err != nil {
			ResponseErr(w, http.StatusUnauthorized, constant.ERR_USER_NOT_FOUND.Error())
			return
		}
	} else {
		token, err := CheckValidToken(tokenHeader)
		if err != nil {
			ResponseErr(w, http.StatusInternalServerError, constant.ERR_TOKEN_IS_INVALID)
			return
		}

		user = models.User{
			Email: token.Email,
			Name:  token.Name,
		}
	}

	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError, constant.ERR_GEN_TOKEN)
		return
	}

	ResponseOk(w, models.Success{
		Token:  tokenString,
		Status: http.StatusOK,
		Url:    "/user/index",
	})
}

func (c *UserControler) parseData(w http.ResponseWriter, tmplPath string, data []models.User) {
	tmpl, _ := template.ParseFiles(tmplPath)
	tmpl.Execute(w, data)
}

func (c *UserControler) List(w http.ResponseWriter, r *http.Request) {
	listUser, err := c.Handler.FindMany(nil)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError, constant.ERR_LIST_USER)
		return
	}

	//json.NewEncoder(w).Encode(listUser)
	//log.Println("listUser", listUser)
	c.parseData(w, "./views/index.html", listUser)
}

func (c *UserControler) Show(w http.ResponseWriter, r *http.Request) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		ResponseErr(w, http.StatusForbidden, constant.ERR_TOKEN_IS_REQUIRED)
		return
	}

	token, err := CheckValidToken(tokenHeader)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError, constant.ERR_TOKEN_IS_INVALID)
		return
	}

	ResponseOk(w, token)
}

//func Show(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, ok := vars["id"]
//	if !ok {
//		fmt.Fprintf(w, "id is missing in parameters")
//	}
//
//	ListUser := model.ListStudent()
//	index := Find(ListUser, id)
//	if index == -1 {
//		tmpl, _ := template.ParseFiles("./views/404.html")
//		tmpl.Execute(w, nil)
//	} else {
//		// Use template internal
//		tmpl := `
//		{{define "user"}}
//		<!DOCTYPE html>
//		<html lang="en">
//		<head>
//		<meta charset="UTF-8">
//		<title>Infor UserId: {{ .Id }}</title>
//		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css">
//		</head>
//		<body>
//		</head>
//		<body>
//			<div class="container mt-5">
//				<h3><a href="/user/list">HOME</a> >> Infor UserId: {{ .Id }}</h3>
//				<div class="row">
//				<table class="table table-striped table-bordered">
//				<tr>
//					<th>Id</th>
//					<th>Name</th>
//					<th>Class</th>
//					<th>Sex</th>
//				</tr>
//				<tr>
//					<td><a href="./{{ .Id }}">{{ .Id }}</a></td>
//					<td><a href="./{{ .Id }}">{{ .Name }}</a></td>
//					<td><a href="./{{ .Id }}">{{ .Class }}</a></td>
//					<td><a href="./{{ .Id }}">{{ .Sex }}</a></td>
//				</tr>
//				</table>
//			</div>
//			</div>
//		</body>
//		</html>
//		{{end}}
//		`
//
//		t, _ := template.New("show").Parse(tmpl)
//		t.ExecuteTemplate(w, "user", ListUser[index])
//	}
//}
//

package route

import (
	"go_mongo/controllers"
	"go_mongo/driver"
	"go_mongo/services"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoute() *mux.Router {
	driver.ConnectMongoDB()
	route := mux.NewRouter().StrictSlash(true)
	handler := services.Factory()
	var userController = &controllers.UserControler{
		Handler: handler,
	}

	userRoute := route.PathPrefix("/user").Subrouter()
	userRoute.HandleFunc("/create", controllers.CreatePage).Name("CreatePage").Methods("GET")
	userRoute.HandleFunc("/login", controllers.LoginPage).Name("LoginPage").Methods("GET")
	userRoute.HandleFunc("/create", userController.Create).Name("Create").Methods("POST")
	userRoute.HandleFunc("/login", userController.Login).Name("Login").Methods("POST")
	userRoute.HandleFunc("/index", userController.List).Name("IndexPage").Methods("GET")
	userRoute.HandleFunc("/show", userController.Show).Name("ShowPage").Methods("GET")

	//Custom 404 page
	route.NotFoundHandler = http.HandlerFunc(controllers.ShowPage404)
	return route
}

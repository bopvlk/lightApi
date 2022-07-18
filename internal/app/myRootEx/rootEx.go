package myRootEx

import (
	"log"
	"net/http"
	"os"

	"7.RootExercise/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type RootEX struct {
	router *mux.Router
	rootIn *models.RootIn
	rOut   *models.RootOut
	l      *log.Logger
}

//API constructor: build base API instance
func NewRootEx() *RootEX {
	return &RootEX{
		router: mux.NewRouter(),
		rootIn: &models.RootIn{},
		rOut:   &models.RootOut{},
		l:      log.Default(),
	}
}

//log.Println("Router initalising successfully. Ready to GO!")
func (roo *RootEX) Start() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	log.Println("find port in a .env file")
	port := os.Getenv("app_port")

	roo.l.Println("start server at port:", port)
	//зконфыгуровано маршрутизатор
	roo.resource()

	return http.ListenAndServe(port, roo.router)
}

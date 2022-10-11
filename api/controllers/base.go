package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cigetbudi/fullstack/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPass, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPass, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Gagal konek ke DB %s ", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Sukses konek ke DB %s ", Dbdriver)
		}
	}

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPass)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Gagal konek ke DB %s ", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Sukses konek ke DB  %s ", Dbdriver)
		}
	}

	//migrasi
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	server.Router = mux.NewRouter()

}

func (server *Server) Run(addr string) {
	fmt.Println("Listening ke port ...")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

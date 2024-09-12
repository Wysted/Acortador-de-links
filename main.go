package main

import (
	"log"

	"github.com/Wysted/shortLink/config"
	"github.com/Wysted/shortLink/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar las variables de entorno: ", err)
		return
	}
	db := config.InitDatabase()

	
	r := gin.Default()


	routes.LinkRouter(r,db)
	
	port := ":8081"
	log.Fatal(r.Run(port))
}


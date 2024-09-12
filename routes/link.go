package routes

import (
	"log"
	"net/http"

	"github.com/Wysted/shortLink/models"
	"github.com/Wysted/shortLink/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LinkRouter(r *gin.Engine, db *gorm.DB){
	linkGroup := r.Group("/")

	linkGroup.POST("",func(c *gin.Context){
		var newLink models.Link
		if err:= c.BindJSON(&newLink);err != nil{
			log.Printf("Error al deserializar el cuerpo de la solicitud: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error al deserializar el cuerpo de la solicitud"})
			return
		}

		if err := services.CreateLink(&newLink,db); err != nil{
			log.Printf("Error al crear el registro: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el registro"})
			return
		}
		log.Printf("Registro creado correctamente\n")
		c.JSON(http.StatusOK, gin.H{"message": "Registro creado correctamente"})
	})

	linkGroup.DELETE(":id", func(c *gin.Context) {
		idLink := c.Param("id")

		if err := services.DeleteLink(idLink,db); err != nil{
			log.Printf("Error al eliminar el registro: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el registro"})
			return
		}
		log.Printf("Registro eliminado correctamente\n")
		c.JSON(http.StatusOK, gin.H{"message": "Registro eliminado correctamente"})
	})

	linkGroup.PATCH(":id",func(c *gin.Context) {
		idLink := c.Param("id")
		var newLink models.UpdateLink
		if err:= c.BindJSON(&newLink);err != nil{
			log.Printf("Error al deserializar el cuerpo de la solicitud: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error al deserializar el cuerpo de la solicitud"})
			return
		}
		services.UpdateLink(idLink,&newLink,db)
		log.Printf("Registro actualizado correctamente\n")
		c.JSON(http.StatusOK, gin.H{"message": "Registro actualizado correctamente"})
	})

	linkGroup.GET(":name",func(c *gin.Context) {
		nameLink := c.Param("name")
		link,err := services.GetLinkByName(nameLink,db)
		if err != nil{
			log.Printf("Error al recibir el link por nombre: %v\n", err)
			return
		}
		c.Redirect(http.StatusMovedPermanently,link.URL)
	})

}
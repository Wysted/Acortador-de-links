package services

import (
	"errors"
	"log"

	"github.com/Wysted/shortLink/models"
	"gorm.io/gorm"
)



func GetLinkByID(id string, db *gorm.DB) (*models.Link, error) {
	var link models.Link
	result := db.Model(&link).First(&link,"id = ?",id); 
	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		return nil,errors.New("link no encontrado")
	}
	return &link,nil
}
func GetLinkByName(name string, db *gorm.DB) (*models.Link, error) {
	var link models.Link
	result := db.Model(&link).First(&link,"name = ?",name); 
	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		return nil,errors.New("link no encontrado")
	}
	return &link,nil
}

func CreateLink(link *models.Link,db *gorm.DB) error{
	result := db.Create(&link)	
	if result.Error != nil{
		log.Printf("Error al crear el link")
	}
	return nil
}

func DeleteLink(id string, db *gorm.DB) error {
	var link models.Link

	result := db.Delete(&link,id)
	if result.Error != nil {
		log.Printf("Error al eliminar la ciudad: %v\n", result.Error)
		return result.Error
	}

	log.Printf("Ciudad eliminada correctamente\n")
	return nil
	
}

func UpdateLink(id string,updateLink *models.UpdateLink, db *gorm.DB) error {

	result,err := GetLinkByID(id,db)
	if err != nil{
		log.Printf("Error al recibir el link por id: %v\n", err)
		return err
	}
	if len(updateLink.Name) > 0{
		result.Name = updateLink.Name
	}
	if len(updateLink.URL ) > 0 {
		result.URL = updateLink.URL
	}
	
	db.Save(result)
	return nil
}
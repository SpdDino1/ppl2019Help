package main

import ( 
 	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
 	_ "github.com/jinzhu/gorm/dialects/mysql"
 	"fmt"
 	"models"
)

func main(){
	
	db,err:=gorm.Open("mysql","root:@tcp(127.0.0.1:3306)/gorm")
	if err!=nil{
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to db!");

	db.AutoMigrate(
		&models.Class{},
		&models.Student{})
	
	db.Model(&models.Student{}).AddForeignKey("class_id","classes(class_id)","RESTRICT","RESTRICT")
}
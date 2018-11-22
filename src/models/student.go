package models

import (
 	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct{
	StudentId int `gorm:"primary_key;auto_increment"`
	Class Class
	ClassId int 
	Name string
}
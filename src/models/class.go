package models	

import (
 	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Class struct{
	ClassId int `gorm:"primary_key;auto_increment"`
	Grade int
	Subject string
}
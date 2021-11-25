package model

import "gorm.io/gorm"

type Status int8

var TablePrefix = ""
var Db *gorm.DB

package models

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Entity  string `gorm:"column:entity;type:varchar(64);not null;uniqueIndex:entity_scope"`
	Scope   string `gorm:"column:scope;type:varchar(64);not null;uniqueIndex:entity_scope"`
	Allowed bool   `gorm:"column:allowed;not null;index"`
}

const GlobalEntity = "*"
const RolePrefix = "R$"

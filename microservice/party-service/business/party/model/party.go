package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Admin - represent finxlog_table
type Admin struct {
	ID   string `gorm:"primary_key" json:"id,omitempty"`
	Name string `gorm:"not_null" json:"name,omitempty"`
}

// Kudo represents a oos kudo.
type Party struct {
	ID   int    `gorm:"column:ID" json:"id,omitempty"`
	Name string `gorm:"NAME" json:"name"`
	Size int    `gorm:"SIZE" json:"size"`
}

type Member struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	PartyID    primitive.ObjectID `json:"partyID" bson:"partyID"`
	SizeMember int                `json:"sizeMember" bson:"sizeMember"`
	ListMember []string           `json:"listMember" bson:"listMember"`
}

package model

// Admin - represent finxlog_table
type Admin struct {
	ID   string `gorm:"primary_key" json:"id,omitempty"`
	Name string `gorm:"not_null" json:"name,omitempty"`
}

// Kudo represents a oos kudo.
type Party struct {
	ID   int    `gorm:"column:ID" json:"id,omitempty"`
	Name string `gorm:"column:NAME" json:"name"`
	Size int    `gorm:"column:SIZE" json:"size"`
}

type Member struct {
	ID      int    `gorm:"column:ID" json:"id,omitempty"`
	PartyID int    `gorm:"column:PART_ID" json:"partyID"`
	UserID  string `gorm:"column:USER_ID" json:"userID"`
}

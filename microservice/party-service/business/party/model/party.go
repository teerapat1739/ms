package model

import (
	"math/rand"
	"strings"
	"time"
)

// Admin - represent finxlog_table
type Admin struct {
	ID   string `gorm:"primary_key" json:"id,omitempty"`
	Name string `gorm:"not_null" json:"name,omitempty"`
}

// Kudo represents a oos kudo.
type Party struct {
	ID            int      `gorm:"column:ID" json:"id,omitempty"`
	Name          string   `gorm:"column:NAME" json:"name"`
	Size          int      `gorm:"column:SIZE" json:"size"`
	IsClose       bool     `gorm:"column:IS_CLOSE" json:"IsClose,omitempty"`
	Category      string   `gorm:"-" json:"category,omitempty"`
	Description   string   `gorm:"-" json:"description,omitempty"`
	Featuredimage string   `gorm:"-" json:"featuredImage,omitempty"`
	Images        []string `gorm:"-" json:"images,omitempty"`
	Location      string   `gorm:"-" json:"location,omitempty"`
	Date          string   `gorm:"-" json:"date,omitempty"`
	Time          string   `gorm:"-" json:"time,omitempty"`
}

func (f *Party) FillDefaultParty() {
	now := time.Now()
	f.Description = "Come to our donation drive to help us replenish our stock of pet food, toys, bedding, etc. We will have live bands, games, food trucks, and much more."
	f.Featuredimage = "https://placekitten.com/500/500"
	f.Location = "1234 Dog Alley"
	f.Category = "Fundraising"
	c := rand.Intn(10-1) + 1
	s := strings.Split(now.AddDate(0, 0, c).String(), " ")
	f.Date = s[0]
	f.Time = "12:00"
	f.Images = []string{"https://placekitten.com/500/500", "https://placekitten.com/500/500", "https://placekitten.com/500/500"}
}

type Member struct {
	ID      int    `gorm:"column:ID" json:"id,omitempty"`
	PartyID int    `gorm:"column:PART_ID" json:"partyID"`
	UserID  string `gorm:"column:USER_ID" json:"userID"`
}

type Return struct {
	Data     interface{} `json:"data"`
	ErrorMsg string      `json:"errorMsg"`
}

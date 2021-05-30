package repository

import (
	"log"
	"party-service/business/party"
	"party-service/business/party/model"

	"github.com/globalsign/mgo"
	"gorm.io/gorm"
)

const (
	collectionName = "kudos"
)

func GetCollectionName() string {
	return collectionName
}

type partyRepository struct {
	db        *gorm.DB
	tableName string
	logger    *log.Logger
	session   *mgo.Session
}

// NewRepository - init user repository
func NewRepository(db *gorm.DB, tableName string) party.Repository {
	db.Debug()
	return &partyRepository{
		db:        db,
		tableName: tableName,
	}
}

// Find fetches a kudo from mongo according to the query criteria provided.
func (r *partyRepository) Find(repoID string) (model.Party, error) {

	var kudo model.Party

	return kudo, nil
}

// FindAll fetches all kudos from the database. YES.. ALL! be careful.
func (r *partyRepository) FindAll(selector map[string]interface{}) ([]model.Party, error) {

	var kudos []model.Party

	return kudos, nil
}

// Delete deletes a kudo from mongo according to the query criteria provided.
func (r *partyRepository) Delete(kudo model.Party) error {

	return nil
}

// Update updates an kudo.
func (r *partyRepository) Update(kudo model.Party) error {

	return nil
}

// Create kudos in the database.
func (r *partyRepository) Create(p model.Party) error {
	result := r.db.Table("PARTIES").Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Create kudos in the database.
func (r *partyRepository) JoinParty(m model.Member) error {
	result := r.db.Table("MEMBERS").Create(&m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Count counts documents for a given collection
func (r *partyRepository) Count() (int, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)
	return coll.Count()
}

func (r *partyRepository) GetParty() ([]model.Party, error) {
	// r.db = r.db.Debug()
	var m []model.Party
	result := r.db.Table("PARTIES").Select("*").Where("IS_CLOSE = ?", false).Scan(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}

func (r *partyRepository) CountMemberByPartyID(partyID int) (int64, error) {
	var count int64
	result := r.db.Table("MEMBERS").Where("PART_ID = ?", partyID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (r *partyRepository) GetPartyByID(partyID int) (model.Party, error) {
	var p model.Party
	r.db = r.db.Debug()
	result := r.db.Table("PARTIES").Where("ID = ?", partyID).Scan(&p)
	if result.Error != nil {
		return p, result.Error
	}
	return p, nil
}

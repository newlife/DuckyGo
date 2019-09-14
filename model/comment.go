package model

import (
	"database/sql/driver"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	//	"github.com/jinzhu/gorm"
)

// JSONB mysql json field
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

// Comment ....
type Comment struct {
	ID       uint
	URL      string `gorm:"column:url"`
	name     string
	location string
	Rate     uint
	title    string
	keyword  JSONB `gorm:"column:key_word"`
	cid      uint  `gorm:"column:c_id"`
	vid      uint  `gorm:"column:v_id"`
}

func GetComment(ID interface{}) (Comment, error) {
	var comment Comment
	result := DB.First(&comment, ID)
	return comment, result.Error
}

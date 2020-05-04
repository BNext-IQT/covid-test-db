package poc

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Poc struct {
    Id      uuid.UUID     `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    string        `json:"name" gorm:"column:name; type:string; not_null"` 
    Phone   string        `json:"phone" gorm:"column:phone; type:string;"`
    Email   string        `json:"email" gorm:"column:email; type:string;"`
}

func (Poc) TableName() string {
    return "pocs"
}

func Create(db *gorm.DB, name string, email string, phone string) (*Poc, error) {
    var toInsert = &Poc{
        Name: name,
        Email: email,
        Phone: phone,
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *Poc) (*Poc, error) {
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*Poc, error) {
    result :=  &Poc{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchByNameAndEmail(db *gorm.DB, name string, email string) (*Poc, error) {
    result :=  &Poc{}

    err := db.Where("name = ? AND email = ?", name, email).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]Poc, error) {
    var results []Poc =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}
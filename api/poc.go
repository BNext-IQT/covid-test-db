package poc

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	
	// Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    // Necessary in order to check for transaction retry error codes.
    "github.com/lib/pq"
)

type poc struct {
    id   uuid.UUID `gorm:"type:uuid;primary_key;"`
    Name string NOT NULL,
	phone string NULL,
	email string NULL,
}

// Functions of type `txnFunc` are passed as arguments to our
// `runTransaction` wrapper that handles transaction retries for us
// (see implementation below).
type txnFunc func(*gorm.DB) error

// This function is used for testing the transaction retry loop.  It
// can be deleted from production code.
var forceRetryLoop txnFunc = func(db *gorm.DB) error {

    // The first statement in a transaction can be retried transparently
    // on the server, so we need to add a dummy statement so that our
    // force_retry statement isn't the first one.
    if err := db.Exec("SELECT now()").Error; err != nil {
        return err
    }
    // Used to force a transaction retry.
    if err := db.Exec("SELECT crdb_internal.force_retry('1s'::INTERVAL)").Error; err != nil {
        return err
    }
    return nil
}

func create(db *gorm.DB, name, email, phone) (poc, error) {
	var toInsert poc
	id, err := uuid.NewV4()
	poc.id := id
	poc.name := name
	poc.email := email
	poc.phone := phone

	if err := db.Save(&toInsert).Error; err != nil {
        return err
    }
    return toInsert
}

func fetchById(db *gorm.DB, id) (poc, error) {
    var result poc

    if err := db.First(&result, id).Error; err != nil {
    	result := nil
        return err
    }

    return result
}
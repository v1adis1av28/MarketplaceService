package ad

import (
	"context"
	"fmt"
	"log"
	"mp-service/internal/models"

	"github.com/jackc/pgx/v5"
)

type AdRepository struct {
	db *pgx.Conn
}

func (a *AdRepository) NewAd(ads models.AdsDTO, email string) error {
	tx, err := a.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("Error on openning transaction")
	}
	defer tx.Rollback(context.Background())

	sqlStatement := "INSERT INTO ADVERTISEMENT (HEADER,DESCRIPTION,IMAGE_URL,PRICE,OWNER_ID) SELECT $1, $2,$3,$4, USERS.ID FROM USERS WHERE EMAIL = $5"
	_, err = tx.Exec(context.Background(), sqlStatement, ads.Header, ads.Description, ads.ImageUrl, ads.Price, email)
	if err != nil {
		return fmt.Errorf("Error while inserting new ads in db")
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("error commiting transaction")
	}
	log.Printf("creating new ads successfuly")
	return nil
}

func NewAdRepository(db *pgx.Conn) *AdRepository {
	return &AdRepository{db: db}
}

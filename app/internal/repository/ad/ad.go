package ad

import (
	"context"
	"fmt"
	"log"
	"mp-service/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type AdRepository struct {
	db *pgx.Conn
}

func (a *AdRepository) NewAd(ads models.AdsDTO, email string) (*models.Advertisement, error) {
	tx, err := a.db.Begin(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Error on openning transaction")
	}
	defer tx.Rollback(context.Background())
	var newAd models.Advertisement

	sqlStatement := `
        INSERT INTO ADVERTISEMENT (HEADER, DESCRIPTION, IMAGE_URL, PRICE, OWNER_ID, CREATED_AT)
		SELECT $1, $2, $3, $4, USERS.ID, $5
		FROM USERS
		WHERE EMAIL = $6
		RETURNING HEADER, DESCRIPTION, IMAGE_URL, PRICE, OWNER_ID, CREATED_AT`

	err = tx.QueryRow(context.Background(), sqlStatement,
		ads.Header, ads.Description, ads.ImageUrl, ads.Price, time.Now(), email,
	).Scan(
		&newAd.Header,
		&newAd.Description,
		&newAd.ImageUrl,
		&newAd.Price,
		&newAd.OwnerID,
		&newAd.Created_at,
	)
	if err := tx.Commit(context.Background()); err != nil {
		return nil, fmt.Errorf("error commiting transaction")
	}
	fmt.Println("repo")
	fmt.Println(newAd)
	log.Printf("creating new ads successfuly")
	return &newAd, nil
}

func NewAdRepository(db *pgx.Conn) *AdRepository {
	return &AdRepository{db: db}
}

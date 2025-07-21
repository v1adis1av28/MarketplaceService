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

func (a *AdRepository) FetchAds(sort, order, limit, offset, email string) ([]models.AdFeed, error) {
	validSortFields := map[string]bool{"created_at": true, "price": true}
	if !validSortFields[sort] {
		sort = "created_at"
	}
	if order != "asc" && order != "desc" {
		order = "desc"
	}

	query := fmt.Sprintf(`
		SELECT a.ID, a.HEADER, a.DESCRIPTION, a.IMAGE_URL, a.PRICE, a.CREATED_AT, u.EMAIL, 
		       CASE WHEN u.EMAIL = $1 THEN true ELSE false END as is_owner
		FROM ADVERTISEMENT a
		JOIN USERS u ON u.ID = a.OWNER_ID
		ORDER BY a.%s %s
		LIMIT $2 OFFSET $3`, sort, order)

	rows, err := a.db.Query(context.Background(), query, email, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ads []models.AdFeed
	for rows.Next() {
		var ad models.AdFeed
		err := rows.Scan(&ad.ID, &ad.Header, &ad.Description, &ad.ImageUrl, &ad.Price, &ad.CreatedAt, &ad.OwnerEmail, &ad.IsOwner)
		if err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}
	return ads, nil
}

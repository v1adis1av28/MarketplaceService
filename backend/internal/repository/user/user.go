package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mp-service/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func (u *UserRepository) RegistrateUser(email, hashedPassword string) error {
	tx, err := u.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("Error on openning transaction")
	}
	defer tx.Rollback(context.Background())

	sqlStatement := "INSERT INTO USERS (email,password) VALUES ($1,$2)"
	_, err = tx.Exec(context.Background(), sqlStatement, email, hashedPassword)
	if err != nil {
		return fmt.Errorf("User with that email already exists")
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("error commiting transaction")
	}
	log.Printf("creating new user successfuly")
	return nil
}

func (u *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	sqlStatement := "SELECT u.EMAIL, u.PASSWORD, u.ROLE FROM USERS as u where u.EMAIL = $1"
	var user models.User
	err := u.db.QueryRow(context.Background(), sqlStatement, email).Scan(&user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user with that email not found")
		}
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

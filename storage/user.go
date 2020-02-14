package storage

import (
	"database/sql"

	"github.com/huandu/go-sqlbuilder"
)

type User struct {
	Base

	Name          string
	Discriminator string
	Currency      uint64
	LastTimely    sql.NullTime
}

func (u *User) UserName(maxlength int) string {
	runes := []rune(u.Name)
	if len(runes) > maxlength {
		runes = runes[:maxlength]
		return string(runes) + "..."
	}
	return u.Name + "#" + u.Discriminator
}

type usersTable struct{}

func (*usersTable) name() string {
	return "users"
}

func (ut *usersTable) createTable() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id            VARCHAR(255) PRIMARY KEY ON CONFLICT REPLACE,
			created_at    DATETIME DEFAULT NULL,
			updated_at    DATETIME DEFAULT NULL,
			name          VARCHAR(255) NOT NULL,
			discriminator VARCHAR(255) NOT NULL,
			currency      BIGINT DEFAULT 0,
			last_timely   DATETIME DEFAULT NULL
		);` +
		base(ut.name()))
	return err
}

func (*usersTable) InsertTx(tx *sql.Tx, u *User) error {
	query := sqlbuilder.NewInsertBuilder().
		InsertInto("users").
		Cols("id", "name", "discriminator").
		Values(u.ID, u.Name, u.Discriminator)

	sql, args := query.Build()
	_, err := tx.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (*usersTable) ByID(tx *sql.Tx, id string) (*User, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"id",
		"created_at",
		"updated_at",
		"name",
		"discriminator",
		"currency",
		"last_timely")
	sb.From("users")
	sb.Where(sb.Equal("id", id))

	sql, args := sb.Build()
	row := tx.QueryRow(sql, args...)

	var user User

	err := row.Scan(
		&user.ID,
		&user.InsertedAt,
		&user.UpdatedAt,
		&user.Name,
		&user.Discriminator,
		&user.Currency,
		&user.LastTimely,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

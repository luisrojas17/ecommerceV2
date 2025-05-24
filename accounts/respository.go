package accounts

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// Interface definition
type Respository interface {
	Close()
	Save(ctx context.Context, a Account) error
	GetByID(ctx context.Context, id string) (*Account, error)
	GetAll(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

// Struct defined which it will be who implement the interface
type postgresRespository struct {
	db *sql.DB
}

func NewPostgresRespository(url string) (Respository, error) {

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRespository{db}, nil
}

// To implement the interface you have to implement all methods defined in the interface
// and change the return type by struct that implements the interface.
func (r *postgresRespository) Close() {
	r.db.Close()
}

func (r *postgresRespository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRespository) Save(ctx context.Context, a Account) error {

	var statement = "INSERT INTO ACCOUNTS(ID, NAME) VALUES($1, $2)"

	_, err := r.db.ExecContext(ctx, statement, a.ID, a.Name)

	return err
}

func (r *postgresRespository) GetByID(ctx context.Context, id string) (*Account, error) {

	var statement = "SELECT ID, NAME FROM ACCOUNTS WHERE ID = $1"

	row := r.db.QueryRowContext(ctx, statement, id)

	a := &Account{}

	if err := row.Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *postgresRespository) GetAll(ctx context.Context, skip uint64, take uint64) ([]Account, error) {

	var statement = "SELECT ID, NAME FROM ACCOUNTS ORDER BY ID DESC OFFSET $1 LIMIT $2"

	rows, err := r.db.QueryContext(
		ctx, statement, skip, take,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []Account{}

	for rows.Next() {
		a := &Account{}
		if err = rows.Scan(&a.ID, &a.Name); err == nil {
			accounts = append(accounts, *a)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

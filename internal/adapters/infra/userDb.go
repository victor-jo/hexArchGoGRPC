package infra

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"hex/internal/application/domain"
)

type Adapter struct {
	db *bun.DB
}

func NewUserDbAdapter() (*Adapter, error) {
	dsnFormat := "%s://%s:%s@%s:%d/?sslmode=disable"
	dsn := fmt.Sprintf(dsnFormat, dbtype, user, password, host, port)
	pgconn := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	db := bun.NewDB(sql.OpenDB(pgconn), pgdialect.New())

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Adapter{db: db}, nil
}

func (db Adapter) CreateUser(username string, password string) (uint64, error) {
	// TODO : create user orm
	return 0, nil
}

func (db Adapter) FindUser(id uint64) (domain.UserEntity, error) {
	// TODO : select user orm
	return domain.UserEntity{
		ID:       1,
		Username: "asdf",
		Password: "asdf",
	}, nil
}

package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Database struct {
	DB
	Driver
}

func New() (*Database, error) {
	dbDriver := "postgres"
	DATABASE_URL := os.Getenv("DATABASE_URL")

	database, err := sqlx.Open(dbDriver, DATABASE_URL)
	if err != nil {
		return nil, fmt.Errorf("failed with postgres credentials: %w", err)
	}

	err = database.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	log.Println("Successful connection to database")

	driver := &sqlxDriver{database}
	db := &sqlxDB{database}

	return &Database{
		DB:     db,
		Driver: driver,
	}, nil
}

type Driver interface {
	Querier
	Execer
}

type DB interface {
	Begin(ctx context.Context) (Tx, error)
	Close() error
}

type Tx interface {
	Querier
	Execer
	Commit() error
	Rollback() error
}

type Querier interface {
	QueryRowx(ctx context.Context, query string, args ...any) *sqlx.Row
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedQuery(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type Execer interface {
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type sqlxDriver struct {
	db *sqlx.DB
}

func (d *sqlxDriver) QueryRowx(ctx context.Context, query string, args ...any) *sqlx.Row {
	return d.db.QueryRowxContext(ctx, query, args...)
}

func (d *sqlxDriver) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.SelectContext(ctx, dest, query, args...)
}

func (d *sqlxDriver) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.GetContext(ctx, dest, query, args...)
}

func (d *sqlxDriver) NamedQuery(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	return d.db.NamedQueryContext(ctx, query, arg)
}

func (d *sqlxDriver) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return d.db.NamedExecContext(ctx, query, arg)
}

func (d *sqlxDriver) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

type sqlxDB struct {
	db *sqlx.DB
}

func (d *sqlxDB) Begin(ctx context.Context) (Tx, error) {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &sqlxTX{tx}, nil
}

func (d *sqlxDB) Close() error {
	return d.db.Close()
}

type sqlxTX struct {
	tx *sqlx.Tx
}

func (d *sqlxTX) QueryRowx(ctx context.Context, query string, args ...any) *sqlx.Row {
	return d.tx.QueryRowxContext(ctx, query, args...)
}

func (d *sqlxTX) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.tx.SelectContext(ctx, dest, query, args...)
}

func (d *sqlxTX) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.tx.GetContext(ctx, dest, query, args...)
}

func (d *sqlxTX) NamedQuery(_ context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	return d.tx.NamedQuery(query, arg)
}

func (d *sqlxTX) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return d.tx.NamedExecContext(ctx, query, arg)
}

func (d *sqlxTX) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.tx.ExecContext(ctx, query, args...)
}

func (d *sqlxTX) Commit() error {
	return d.tx.Commit()
}

func (d *sqlxTX) Rollback() error {
	return d.tx.Rollback()
}

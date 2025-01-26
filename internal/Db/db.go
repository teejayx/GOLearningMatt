package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Item struct {	
	Task string
	Status string	
}

type DB struct {
	pool *pgxpool.Pool
}

func New(user, password, dbname, host string, port int) (*DB, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}
	return &DB{pool: pool}, nil
}

func (db *DB) InsertItems(ctx context.Context, item Item) error {
	query := `INSERT INTO todo_items (task, status) VALUES ($1, $2)`
	_, err := db.pool.Exec(ctx, query, item.Task, item.Status)
	return err
}

func (db *DB) GetItems(ctx context.Context) ([]Item, error) {
	query := `SELECT task, status FROM todo_items`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Task, &item.Status); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return items, nil
}

func (db *DB) Close() {
	db.pool.Close()
}	
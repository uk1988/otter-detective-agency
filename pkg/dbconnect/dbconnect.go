package dbconnect

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectWithRetry(connString string, maxRetries int, retryInterval time.Duration) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error
	for i := 0; i < maxRetries; i++ {
		pool, err = pgxpool.Connect(context.Background(), connString)
		if err == nil {
			return pool, nil
		}
		fmt.Printf("Failed to connect to database. Retrying in %s\n", retryInterval)
		time.Sleep(retryInterval)
	}
	return nil, fmt.Errorf("Failed to connect to database after %d retries", maxRetries)
}

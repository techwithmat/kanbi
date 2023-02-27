package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/techwithmat/kanbi/config"
)

func NewDBConnection(ctx context.Context, c *config.Database) (*pgx.Conn, error) {
	// Url Example: postgres://username:password@localhost:5432/database_name
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
	)

	connection, err := pgx.Connect(ctx, dns)

	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	if err = connection.Ping(ctx); err != nil {
		log.Fatal("Unable to connect to database")
		return nil, err
	}

	return connection, nil
}

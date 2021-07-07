package data

import (
	"context"
	"marmota/server/data/ent"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	db *ent.Client
}

func NewData(driver string, dsn string) (*Data, func(), error) {
	client, err := ent.Open(
		driver, dsn,
	)

	if err != nil {
		return nil, nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, nil, err
	}

	d := &Data{db: client}

	return d, func() {
		if err := d.db.Close(); err != nil {
			panic(err)
		}
	}, nil
}

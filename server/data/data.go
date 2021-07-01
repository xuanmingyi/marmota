package data

import (
	"context"
	"marmota/server/data/ent"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	db *ent.Client
}

func NewData() (*Data, func(), error) {
	client, err := ent.Open(
		"mysql", "root:123456@tcp(127.0.0.1:3306)/newdb1?parseTime=true",
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

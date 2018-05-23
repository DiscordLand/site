package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Init initializes the database
func Init() (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "test",
	})

	for _, model := range []interface{}{&DiscordUser{}, &User{}, &Bot{}, &WebSession{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

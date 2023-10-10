package postgres

import (
	"fmt"
	"project-3/stores"
)

type Conn struct {
	db string
}

func New(db string) Conn {
	return Conn{db: db}
}

func (c Conn) Create(u stores.User) error {
	fmt.Println("User Created", c.db)
	return nil
}

func (c Conn) Update(u stores.User) error {
	fmt.Println("User Updated")
	return nil
}

func (c Conn) Delete(u stores.User) error {
	fmt.Println("User Updated")
	return nil
}

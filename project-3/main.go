package main

import (
	"project-3/stores"
	"project-3/stores/mysql"
	"project-3/stores/postgres"
)

func main() {
	u := stores.User{
		Name:  "ajay",
		Email: "ajay@email.com",
	}
	m := mysql.New("mysql")
	// stores.StorerInterface = m
	// stores.StorerInterface.Create(u)

	// u2 := stores.User{
	// 	Name:  "ram",
	// 	Email: "sfe@24",
	// }

	 p := postgres.New("Postgres")
	// stores.StorerInterface = p
	// stores.StorerInterface.Create(u2)

	ms:=stores.NewService(m)
		ms.Create(u)
	ps:=stores.NewService(p)
	ps.Create(u)

}

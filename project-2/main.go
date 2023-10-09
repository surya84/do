package main

import (
	"fmt"
	"project-1/db"
)

func main() {
	a, b := db.NewConn("Mysql")
	fmt.Println(a, b)
}

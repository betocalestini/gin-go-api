package main

import (
	"github.com/betocalestini/gin-go-api/database"
	"github.com/betocalestini/gin-go-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}

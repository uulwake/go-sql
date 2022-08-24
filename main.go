package main

import (
	"database/sql"
	"fmt"
	"go-sql/models"
	"go-sql/repositories"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "secret"
		dbname   = "go-sql"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	ItemRepository := repositories.NewItemRepository(db)

	items := ItemRepository.FetchAll()
	log.Printf("items: %+v", items)

	item := ItemRepository.FetchById(1)
	log.Printf("item: %+v", item)

	item.Name = "new item name"
	ItemRepository.UpdateById(1, item)

	item = ItemRepository.FetchById(1)
	log.Printf("item: %+v", item)

	newItem := models.Item{Name: "itemD", Qty: 4, Weight: 10.123}
	ItemRepository.Create(newItem)

	OrderRepository := repositories.NewOrderRepository(db)

	order := models.Order{RecipientName: "user1", RecipientAddress: "address1", Shipper: "shipper1"}
	OrderRepository.CreateOrder(order, item, 10)

	itemCount := ItemRepository.CountAll()
	fmt.Println(itemCount)
}

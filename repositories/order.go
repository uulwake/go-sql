package repositories

import (
	"database/sql"
	"go-sql/models"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return OrderRepository{db}
}

func (orderRepo OrderRepository) CreateOrder(order models.Order, item models.Item, orderQty int) {
	log.Println("=== CREATE ORDER ===")

	tx, err := orderRepo.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	var orderId int
	err = tx.QueryRow(`
		INSERT INTO orders(recipient_name, recipient_address, shipper)
		VALUES
			($1, $2, $3)
		RETURNING ID
	`, order.RecipientName, order.RecipientAddress, order.Shipper).Scan(&orderId)

	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
		INSERT INTO outbounds(item_id, order_id, qty)
		VALUES
			($1, $2, $3)
	`, item.ID, orderId, orderQty)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

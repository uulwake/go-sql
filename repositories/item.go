package repositories

import (
	"database/sql"
	"go-sql/models"
	"log"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return ItemRepository{db}
}

func (itemRepo ItemRepository) CountItems() int {
	log.Println("=== FETCHING ALL ITEMS ===")

	var counter int
	err := itemRepo.db.QueryRow("SELECT count(*) FROM items").Scan(&counter)
	if err != nil {
		log.Fatal(err)
	}

	return counter
}

func (itemRepo ItemRepository) FetchAllItems() []models.Item {
	log.Println("=== FETCHING ALL ITEMS ===")

	items := []models.Item{}

	rows, err := itemRepo.db.Query("SELECT id, name, qty, weight FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Item{}

		err := rows.Scan(&item.ID, &item.Name, &item.Qty, &item.Weight)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("id=%d, name=%s, qty=%d, weight=%f\n", item.ID, item.Name, item.Qty, item.Weight)
		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return items
}

func (itemRepo ItemRepository) FetchById(itemId int) models.Item {
	log.Println("=== FETCH ITEM BY ID ===")

	item := models.Item{}

	err := itemRepo.db.QueryRow("SELECT id, name, qty, weight FROM items WHERE id = $1", itemId).Scan(&item.ID, &item.Name, &item.Qty, &item.Weight)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("id=%d, name=%s, qty=%d, weight=%f\n", item.ID, item.Name, item.Qty, item.Weight)

	return item
}

func (itemRepo ItemRepository) CreateItem(item models.Item) {
	log.Println("=== CREATE NEW ITEM ===")

	_, err := itemRepo.db.Exec(`
		INSERT INTO items(name, qty, weight) 
		VALUES
			($1, $2, $3) 
		RETURNING id
	`, item.Name, item.Qty, item.Weight)

	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepo ItemRepository) UpdateItemById(id int, item models.Item) {
	log.Println("=== UPDATE ITEM BY ID ===")

	_, err := itemRepo.db.Exec(`
		UPDATE items
		SET
			name = $1,
			qty = $2,
			weight = $3
		WHERE
			id = $4
	`, item.Name, item.Qty, item.Weight, item.ID)

	if err != nil {
		log.Fatal(err)
	}
}

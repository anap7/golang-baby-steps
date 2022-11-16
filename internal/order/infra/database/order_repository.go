package database

import (
	"database/sql"
	"fmt"
	"module/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

/*Método que faz parte do OrderRepository e aguarda a "instância" do struct order*/
func (r *OrderRepository) Save(order *entity.Order) error {
	//Preparando a query de inserção
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("DEU ERRRO MEU NA INSERÇÃO DE DADOS");
		return err
	}
	//Executando a query
	/*O "_" significa que é para ignorar a variável de retorno quando executa a query*/
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select (*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
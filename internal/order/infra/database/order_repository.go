package database

import (
	"database/sql"
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
	stmt, err := r.Db.Prepare("INSERT INTO order (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
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
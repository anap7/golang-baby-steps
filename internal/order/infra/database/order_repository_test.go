package database

import (
	"fmt"
	"testing"
	"database/sql"
	"module/internal/order/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

/*O suite faz parte do pacote de testes */
type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

/*O SetupSuite vai rodar antes de executar um teste. O intuito é gerenciar o processo de conexão, gerenciamento
e encerramento do banco de dados. nesse caso, como o banco está em memória, ao rodar o teste a tabela será criada 
e depois fechad*/
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Print("Erro na conexão!!")
		fmt.Println(db)
	}

	suite.NoError(err)

	sqlStmt := "CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))"
	
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Print("Erro na criação da tabela!!")
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	suite.Db = db
}

/*O TearDownTest vai rodar depois de executar um teste*/
func (suite *OrderRepositoryTestSuite) TearDownTest() {
	//Fechando a conexão
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	//Faz parte do mesmo pacote, por isso a função está sendo chamada assim
	/*Injetando a coneção*/
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	//Variável com  o tipo order
	var orderResult entity.Order
	//Validando o cadastro com o select
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", order.ID).
	Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)
	/*O Scan está sendo utilizado para retornar os valores encontrados e atribuidos ao orderResult*/
	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
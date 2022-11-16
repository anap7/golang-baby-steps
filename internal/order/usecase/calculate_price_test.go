package usecase

import (
	"fmt"
	"testing"
	"database/sql"
	"module/internal/order/entity"
	"module/internal/order/infra/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

/*O suite faz parte do pacote de testes */
type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db *sql.DB
}

/*O SetupSuite vai rodar antes de executar um teste. O intuito é gerenciar o processo de conexão, gerenciamento
e encerramento do banco de dados. nesse caso, como o banco está em memória, ao rodar o teste a tabela será criada 
e depois fechad*/
func (suite *CalculatePriceUseCaseTestSuite) SetupSuite() {
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
	suite.OrderRepository = *database.NewOrderRepository(db)
}

/*O TearDownTest vai rodar depois de executar um teste*/
func (suite *CalculatePriceUseCaseTestSuite) TearDownTest() {
	//Fechando a conexão
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new (CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("1", 10, 2)
	suite.NoError(err)
	order.CalculateFinalPrice()

	//Utilizando o DTO para validar os dados
	calculateFinalPriceInput := OrderInputDTO{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
	}

	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)

	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}
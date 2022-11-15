package entity

import "errors"

/*Struct é uma estrutura de dados com informações com tipos,
é semelhante a uma classe*/
type Order struct {
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}

/*Função: A função possui dois retornos, um tipo Order e um erro*/
func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order :=  &Order{
		ID: id,
		Price: price,
		Tax: tax,
	}
	/*Validando os dados e verificando se retornou algum erro */
	err := order.isValid()
	if err != nil {
		//Retornará nulo e o erro
		return nil, err
	}
	//Retornará a ordem e o erro como vazio
	return order, nil
}

/*Método: Nesse caso, isso não é uma função e sim um método. Pois está atrelada ao
struct, assim como métodos funcionam com suas respectivas classes */
func (o *Order) isValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.isValid()
	if err != nil {
		//Retornará o erro
		return err
	}
	//Retornará nulo se não encontrar um erro
	return nil
}
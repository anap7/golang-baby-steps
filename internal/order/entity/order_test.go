package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.isValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.isValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.isValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	/*Quando usamos ":=" significa que estamos criando uma variável e atribuindo a ela 
	um valor */
	order := Order{
		ID: "123",
		Price: 10.0, 
		Tax: 2.0,
	}

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.isValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	//O erro deve estar nulo
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	//O erro deve estar nulo
	assert.Nil(t, err)
	//Deve retornar nulo
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
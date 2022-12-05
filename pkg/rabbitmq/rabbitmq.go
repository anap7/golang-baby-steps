package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error){
	//Criando uma conexão com o rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
}
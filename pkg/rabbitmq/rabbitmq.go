package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error){
	//Criando uma conexão com o rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

/*Essa função vai ler as mensagens da fila do rabbitmq e pegar essas mensagens 
e jogar em um canal

O primeiro parametro é a conexão do rabbitmq e o segundo será o canala onde
jogaremos as mensagens. O canal vai rodar em uma outra thread*/
func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	//Consumindo as mensagens disponíveis da fila
	msgs, err := ch.Consume(
		"orders",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	//Ler mensagem por mensagem
	for msg := range msgs {
		out <- msg
	}
	return nil
}
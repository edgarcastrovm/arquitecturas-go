package kafka

// Placeholder para producer de Kafka
// referencia :github.com/confluentinc/confluent-kafka-go

type CashInProducer struct{}

func NewCashInProducer() *CashInProducer {
	return &CashInProducer{}
}

func (p *CashInProducer) SendCashInEvent(id string) error {
	// Simula envío a Kafka
	return nil
}

package users

import (
	"github.com/Shopify/sarama"
	"github.com/bugscatcher/users/application"
	"github.com/jackc/pgx"
)

type GRPCHandler struct {
	db            *pgx.ConnPool
	kafkaProducer sarama.SyncProducer
}

func New(app *application.App) *GRPCHandler {
	return &GRPCHandler{
		db:            app.DB,
		kafkaProducer: app.KafkaProducer,
	}
}

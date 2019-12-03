package users

import (
	"github.com/Shopify/sarama"
	"github.com/bugscatcher/users/application"
	"github.com/jackc/pgx"
)

type Handler struct {
	db            *pgx.ConnPool
	kafkaProducer sarama.SyncProducer
}

func New(app *application.App) *Handler {
	return &Handler{
		db:            app.DB,
		kafkaProducer: app.KafkaProducer,
	}
}

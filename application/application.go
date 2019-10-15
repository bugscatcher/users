package application

import (
	"github.com/Shopify/sarama"
	"github.com/bugscatcher/users/config"
	"github.com/bugscatcher/users/postgresql"
	"github.com/jackc/pgx"
	"golang.org/x/xerrors"
)

type App struct {
	DB            *pgx.ConnPool
	Config        *config.Config
	KafkaProducer sarama.SyncProducer
}

func New(conf *config.Config) (*App, error) {
	app := &App{}
	app.Config = conf

	db, err := postgresql.NewConnPool(&conf.PostgreSQL)
	if err != nil {
		return nil, xerrors.Errorf("while creating new connection: %w", err)
	}
	app.DB = db

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Version = sarama.V1_1_1_0
	producer, err := sarama.NewSyncProducer(conf.KafkaConf.Brokers, kafkaConfig)
	if err != nil {
		return nil, xerrors.Errorf("while creating new sync producer: %w", err)
	}
	app.KafkaProducer = producer

	return app, nil
}

func (a *App) Close() {
	_ = a.KafkaProducer.Close()
	a.DB.Close()
}

package users

import (
	"github.com/bugscatcher/users/application"
	"github.com/bugscatcher/users/config"
	"github.com/bugscatcher/users/postgresql"
	"github.com/bugscatcher/users/testutil"
	"github.com/jackc/pgx"
)

const testDBName = "users_test"

type TestHandler struct {
	service *Handler
	kafka   *testutil.KafkaMock
	db      *pgx.ConnPool
}

func newTestHandler(expectedCallCount int) *TestHandler {
	testApp := newTestApp()
	kafkaMock := testutil.MockKafkaProducer(expectedCallCount)
	testApp.KafkaProducer = kafkaMock
	h := New(testApp)
	return &TestHandler{
		service: h,
		kafka:   kafkaMock,
		db:      testApp.DB,
	}
}

func newTestApp() *application.App {
	testConfig, err := config.New()
	if err != nil {
		panic(err)
	}
	testConfig.PostgreSQL.Database = testDBName
	testApp := &application.App{}
	testApp.Config = &testConfig
	db, err := postgresql.NewConnPool(&testConfig.PostgreSQL)
	if err != nil {
		panic(err)
	}
	testApp.DB = db
	return testApp
}

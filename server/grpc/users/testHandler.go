package users

import (
	"context"

	"github.com/bugscatcher/users/application"
	"github.com/bugscatcher/users/config"
	"github.com/bugscatcher/users/headers"
	"github.com/bugscatcher/users/postgresql"
	"github.com/bugscatcher/users/testutil"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

const testDBName = "users_test"

type TestHandler struct {
	service *Handler
	kafka   *testutil.KafkaMock
	db      *pgx.ConnPool
	userID  uuid.UUID
	ctx     context.Context
}

func newTestHandler(expectedCallCount int) *TestHandler {
	testApp := newTestApp()
	kafkaMock := testutil.MockKafkaProducer(expectedCallCount)
	testApp.KafkaProducer = kafkaMock
	h := New(testApp)
	userID := uuid.New()
	ctx := headers.AddUserID(context.Background(), userID.String())
	return &TestHandler{
		service: h,
		kafka:   kafkaMock,
		db:      testApp.DB,
		userID:  userID,
		ctx:     ctx,
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

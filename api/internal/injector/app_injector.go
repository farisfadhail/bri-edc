//go:build wireinject

//go:generate wire

package injector

import (
	"bri-edc/api/internal/handler"
	"bri-edc/api/internal/providers"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type AppContainer struct {
	DB *gorm.DB
	//
	AuthHandler        *handler.AuthHandler
	TransactionHandler *handler.TransactionHandler
}

func NewAppContainer(
	db *gorm.DB,
//
	authHandler *handler.AuthHandler,
	transactionHandler *handler.TransactionHandler,
) (*AppContainer, error) {
	return &AppContainer{
		DB: db,
		//
		AuthHandler:        authHandler,
		TransactionHandler: transactionHandler,
	}, nil
}

func InitializeApp() (*AppContainer, error) {
	wire.Build(
		providers.AppProviderSet,
		NewAppContainer,
	)
	return nil, nil
}

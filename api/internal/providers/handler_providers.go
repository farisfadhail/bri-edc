//go:build wireinject

package providers

import (
	"bri-edc/api/internal/handler"

	"github.com/google/wire"
)

var HandlerProviderSet = wire.NewSet(
	handler.NewAuthHandler,
	handler.NewTransactionHandler,
)

//go:build wireinject

package providers

import (
	"bri-edc/api/internal/services"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	services.NewAuthService,
	services.NewTransactionService,
)

//go:build wireinject

package providers

import (
	"bri-edc/api/internal/repositories"

	"github.com/google/wire"
)

var RepositoryProviderSet = wire.NewSet(
	repositories.NewAuthRepository,
	repositories.NewMerchantRepository,
	repositories.NewTerminalRepository,
	repositories.NewTransactionRepository,
	repositories.NewSettlementRepository,
)

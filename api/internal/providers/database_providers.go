//go:build wireinject

package providers

import (
	"bri-edc/api/config"

	"github.com/google/wire"
)

var DatabaseProviderSet = wire.NewSet(
	config.ConnectGormDB,
)

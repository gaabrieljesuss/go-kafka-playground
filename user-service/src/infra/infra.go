package infra

import (
	"github.com/rs/zerolog"
	"go_kafka_playground/api-gateway/src/utils"
)

func Logger() zerolog.Logger {
	return utils.Logger().With().Str("layer", "infra").Logger()
}

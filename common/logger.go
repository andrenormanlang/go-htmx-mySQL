package common

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// / TODO : take a prefix to know where the logs come
// / from. E..g "[CMSGO]" and "[CMSGO-ADMIN]"
func SetupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Logger created")
}

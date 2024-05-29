package common

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// / TODO : take a prefix to know where the logs come
// / from. E..g "[URCHIN]" and "[URCHIN-ADMIN]"
func SetupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Logger created")
}

// package common

// import (

// 	"encoding/json"
//     "os"

// 	"github.com/rs/zerolog"
// 	"github.com/rs/zerolog/log"
// )

// // / TODO : take a prefix to know where the logs come
// // / from. E..g "[CMSGO]" and "[CMSGO-ADMIN]"
// func SetupLogger() {
// 	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
// 	log.Info().Msg("Logger Created")
// }

// func LoadConfig(filePath string, config interface{}) error {
//     file, err := os.Open(filePath)
//     if err != nil {
//  	   return err
//     }
//     defer file.Close()

//     decoder := json.NewDecoder(file)
//     err = decoder.Decode(config)
//     if err != nil {
//  	   return err
//     }

//     log.Info().Msgf("Config loaded from %s", filePath)
//     return nil
// }

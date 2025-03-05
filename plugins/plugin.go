package plugins

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

// Observer interface
type Plugin struct {
	ScriptName string
	Id         string
}

// returns  the transformed arguments
func (p *Plugin) Update(args []string, lua_states map[string]*lua.LState) []string {
	if handler, found := lua_states[p.ScriptName]; found {

		err := handler.DoString(fmt.Sprintf(`result = HandleShortcode ({%s})`, strings.Join(args, ",")))
		if err != nil {
			// TODO: Do we care if the Lua code fails?
			log.Error().Msgf("could not execute the plugin: %v", err)
			return args
		}

		value := handler.GetGlobal("result")

		// This is definitely going to fail, but waht type should be []string in Lua?
		if ret_type := value.Type().String(); ret_type != "[]string" {
			return args
		}
	}

	// TODO : What to do it we haven't got a plugin registered
	log.Error().Msgf("could not find a registered plugin for '%s'", p.ScriptName)
	return args
}

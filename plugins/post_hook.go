package plugins

import (
	"strconv"

	"github.com/andrenormanlang/common"
	lua "github.com/yuin/gopher-lua"
)

type PostHook struct {
	plugins []Plugin

	latestPost common.Post
}

// This runs at the start up, when you register
// all the plugins you want
func (p *PostHook) Register(plugin Plugin) {
	p.plugins = append(p.plugins, plugin)
}

// When u uninstall a plugin
func (p *PostHook) Deregister(plugin Plugin) {
	// TODO: Implement plugin deregistration
	for i, existingPlugin := range p.plugins {
		if existingPlugin == plugin {
			p.plugins = append(p.plugins[:i], p.plugins[i+1:]...)
			break
		}
	}
}

func (p *PostHook) NotifyAll(lua_states map[string]*lua.LState) []string {
	args := []string{
		strconv.Itoa(p.latestPost.Id),
		p.latestPost.Title,
		p.latestPost.Content,
		p.latestPost.Excerpt,
	}

	for _, plugin := range p.plugins {
		args = plugin.Update(args, lua_states)
	}

	return args
}

func (p *PostHook) UpdatePost(title string, content string, excerpt string, lua_states map[string]*lua.LState) common.Post {
	p.latestPost = common.Post{
		Title:   title,
		Content: content,
		Excerpt: excerpt,
		Id:      -1,
	}

	args := p.NotifyAll(lua_states)

	return common.Post{
		Title:   args[1],
		Content: args[2],
		Excerpt: args[3],
		Id:      -1,
	}
}
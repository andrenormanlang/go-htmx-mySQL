package plugins 

import lua "github.com/yuin/gopher-lua"
type Hook struct {
	Plugins map[string][]*lua.LState
}

// post_save -> plugin_main(post_content, post_title, post_excerpt, post_date)
// render_header -> plugin(html string) // passing everything under/inclusively <header> ... </header>
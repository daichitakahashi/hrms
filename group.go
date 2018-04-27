package hrms

import "path"

// GroupRouter is
type GroupRouter struct {
	router *Router
	base   string
}

// GET is
func (g *GroupRouter) GET(_path string, handle Handle) {
	//g.router.GET(path.Join(g.base, _path), handle)
	g.router.GET(g.base+_path, handle)
}

// POST is
func (g *GroupRouter) POST(_path string, handle Handle) {
	g.router.POST(path.Join(g.base, _path), handle)
}

// Group is
func (g GroupRouter) Group(_path string) *GroupRouter {
	g.base = path.Join(g.base, _path)
	return &g
}

// Group is
func (r *Router) Group(path string) *GroupRouter {
	return &GroupRouter{
		router: r,
		base:   path,
	}
}

package hrms

import (
	"net/http"
	"path"
)

// GroupRouter is
type GroupRouter struct {
	router      *Router
	base        string
	middlewares []func(Handler) Handler
	ancestor    *GroupRouter
}

// Group is
func (g *GroupRouter) Group(_path string) *GroupRouter { // その他見直し==============================
	newRouter := &GroupRouter{
		router:   g.router,
		base:     path.Join(g.base, _path),
		ancestor: g,
	}
	return newRouter
}

// Group is
func (r *Router) Group(path string) *GroupRouter {
	return &GroupRouter{
		router: r,
		base:   path,
	}
}

// Use is
func (g *GroupRouter) Use(middlewares ...func(Handler) Handler) {
	if g.middlewares == nil {
		g.middlewares = middlewares
	} else {
		g.middlewares = append(g.middlewares, middlewares...)
	}
}

func (g *GroupRouter) doHandle(handle Handle) Handle {
	return func(w http.ResponseWriter, r *http.Request, ps Params) {
		if g.ancestor != nil {
			g.applyAncestorsMiddlewares(
				applyMiddlewares(handle, g.middlewares),
			).Handle(w, r, ps)
		} else {
			applyMiddlewares(handle, g.middlewares).Handle(w, r, ps)
		}
	}
}

func (g *GroupRouter) applyAncestorsMiddlewares(endpoint Handler) Handler {
	if g.ancestor != nil {
		return endpoint
	}
	return g.ancestor.applyAncestorsMiddlewares(
		applyMiddlewares(endpoint, g.ancestor.middlewares),
	)
}

// GET is
func (g *GroupRouter) GET(_path string, handle Handle) { // 呼ばれるたびにミドルウェアを縮約するハンドラを生成、登録する
	g.router.GET(path.Join(g.base, _path), handle)
	// g.router.GET(g.base+_path, handle)
}

// POST is
func (g *GroupRouter) POST(_path string, handle Handle) {
	g.router.POST(path.Join(g.base, _path), handle)
}

/*
func (g *GroupRouter) moreMiddlewares() []func(Handler) Handler {
	if g.ancestor != nil {
		return append(g.ancestor.middlewares, g.middlewares)
	}
	return g.middlewares
}*/

/*

func (g *GroupRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	g.router.middlewares = append(g.router.middlewares, g.middlewares...) // これだと、サーバとしての起動時にどんどんミドルウェアが溜まってしまう
	g.router.ServeHTTP(w, r)
}

*/

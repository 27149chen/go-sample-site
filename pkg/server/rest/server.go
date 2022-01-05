package rest

import (
	"net/http"

	"go-sample-site/pkg/log"
	ginmiddleware "go-sample-site/pkg/middleware/gin"

	"github.com/gin-gonic/gin"
)

type engine struct {
	*gin.Engine

	mode string
}

func NewEngine() *engine {
	s := &engine{mode: gin.DebugMode}

	gin.SetMode(s.mode)

	s.injectMiddlewares()
	s.injectRouters()

	return s
}

func (s *engine) injectMiddlewares() {
	g := gin.New()
	defer func() {
		s.Engine = g
	}()

	if s.mode == gin.TestMode {
		return
	}

	g.Use(ginmiddleware.RequestLog(log.NewFileLogger("/tmp/requests.log")))
	g.Use(gin.Recovery())
	g.Use(ginmiddleware.RequestID())
}

func (s *engine) injectRouters() {
	g := s.Engine

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Invalid path: %s", c.Request.URL.Path)
	})
	g.HandleMethodNotAllowed = true
	g.NoMethod(func(c *gin.Context) {
		c.String(http.StatusMethodNotAllowed, "Method not allowed: %s %s", c.Request.Method, c.Request.URL.Path)
	})

	apiRouters := g.Group("")
	s.injectRouterGroup(apiRouters)

	s.Engine = g
}

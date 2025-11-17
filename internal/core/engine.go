package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harluo/gin/internal/config"
)

type Engine struct {
	*shadowEngine
}

func newEngine(config *config.Server) (engine *Engine, err error) {
	handler := gin.Default()

	server := new(http.Server)
	server.Addr = config.Addr()
	server.Handler = handler
	if nil != config.Timeout && 0 != config.Timeout.Read {
		server.WriteTimeout = config.Timeout.Read
	}
	if nil != config.Timeout && 0 != config.Timeout.Write {
		server.WriteTimeout = config.Timeout.Write
	}
	if err = server.ListenAndServe(); nil == err {
		engine = new(Engine)
		engine.shadowEngine = (*shadowEngine)(handler)
	}

	return
}

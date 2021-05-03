package ginx

import (
	"fmt"
	"github.com/Tooooommy/go-one/server/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Default value
var (
	Name = "go-one"
)

type Server struct {
	eng *gin.Engine
	cfg config.HttpConfig
}

type ServerOption func(s *Server)

// NewServer: 实例化Server
func NewServer(options ...ServerOption) *Server {
	s := &Server{
		eng: gin.New(),
		cfg: config.HttpConfig{
			Name: Name,
			Host: "127.0.0.1",
			Port: 9091,
		},
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

// WithGinEngine: 设置GinEngine
func WithGinEngine(eng *gin.Engine) ServerOption {
	return func(s *Server) {
		s.eng = eng
	}
}

// WithConfig: 设置Config
func WithConfig(cfg config.HttpConfig) ServerOption {
	return func(s *Server) {
		s.cfg = cfg
	}
}

// GinEngine: 获取gin.Engine
func (s *Server) GinEngine() *gin.Engine {
	return s.eng
}

// Config: 获取config.HttpConfig配置
func (s *Server) Config() config.HttpConfig {
	return s.cfg
}

// Start: 启动服务
func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	if len(s.cfg.CertFile) == 0 && len(s.cfg.KeyFile) == 0 {
		return s.eng.Run(addr)
	} else {
		return s.eng.RunTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
	}
}

// ServeHTTP: 实现HTTP Serve
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.eng.ServeHTTP(w, r)
}

func wrapHs(handlers ...http.Handler) gin.HandlersChain {
	var chain gin.HandlersChain
	for _, handler := range handlers {
		chain = append(chain, gin.WrapH(handler))
	}
	return chain
}
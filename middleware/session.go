package middleware

import (
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Store memstore.Store

var options = sessions.Options{
	HttpOnly: true,
	MaxAge:   3 * 86400,
	Path:     "/",
	SameSite: http.SameSiteNoneMode,
	Secure:   true,
}

// Session Session配置中间件
func Session() gin.HandlerFunc {
	Store = memstore.NewStore([]byte(conf.SystemConfig.SessionSecret))
	Store.Options(options)
	return sessions.Sessions(conf.DatabaseConfig.TablePrefix+"session", Store)
}

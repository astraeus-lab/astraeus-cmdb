package web

import (
	"net/http"
	"time"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/log"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
	"github.com/astraeus-lab/astraeus-cmdb/src/web/router"

	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(customRecovery(), accessLogWithFormatter())
	router.RegistAllRoute(engine)

	return engine
}

func accessLogWithFormatter() gin.HandlerFunc {

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		latencyTime := time.Now().Sub(startTime).Milliseconds()

		log.Info(c.ClientIP(),
			"Method", c.Request.Method,
			"URI", c.Request.RequestURI,
			"Protocol", c.Request.Proto,
			"Status", c.Writer.Status(),
			"Latency", latencyTime,
			"Host", c.Request.Host,
			"UA", c.Request.UserAgent(),
			"Referer", util.StrWithDefault(c.Request.Referer(), "-"),
			"ResponseSize", c.Writer.Size(),
			"Error", util.StrWithDefault(c.Errors.String(), "-"),
		)
	}
}

func customRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				startTime := time.Now()
				c.Next()
				latencyTime := time.Now().Sub(startTime).Milliseconds()

				log.Error(c.ClientIP(),
					"Method", c.Request.Method,
					"URI", c.Request.RequestURI,
					"Protocol", c.Request.Proto,
					"Status", c.Writer.Status(),
					"Latency", latencyTime,
					"Host", c.Request.Host,
					"UA", c.Request.UserAgent(),
					"Referer", c.Request.Referer(),
					"ResponseSize", c.Writer.Size(),
					"Error", util.StrWithDefault(c.Errors.String(), "-"),
				)

				c.JSON(http.StatusInternalServerError, "")
				c.Abort()
			}
		}()

		c.Next()
	}
}

/*
 * Define routers
 */

package routers

import (
	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/noahzyl/mindmate/config"
	"github.com/noahzyl/mindmate/controllers"
	"github.com/noahzyl/mindmate/pkg/logger"
)

func Router() *gin.Engine {
	r := gin.Default() // Create a router

	// Set a logger
	r.Use(gin.LoggerWithConfig(logger.LogRequest()))
	r.Use(logger.LogError)

	// Set redis and sessions
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", "", []byte("secret"))
	r.Use(sessions.Sessions("mindmate_user_session", store))

	// Set a router group of users (user router)
	user := r.Group("/user")
	{
		// controllers.UserController{} will create an anonymous variable of UserController
		user.POST("/register", (&controllers.UserController{}).Register)
		user.POST("/login", (&controllers.UserController{}).Login)
	}

	return r
}

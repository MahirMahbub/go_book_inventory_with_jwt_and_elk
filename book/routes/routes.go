package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go_practice/book/controllers"
	"go_practice/book/elasticsearch"
	"go_practice/book/logger"
	"go_practice/book/middlewares"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func SetupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin_book.log")
	r := gin.New()
	gin.DefaultWriter = io.MultiWriter(f)
	logger.Logger(f)
	//log.SetOutput(f)
	r.Use(middlewares.CORSMiddleware())
	r.Use(elasticsearch.Client())

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := controllers.NewController()

	v1 := r.Group("/api/v1")
	{
		BookRoute(v1, c)
		ElasticRoute(v1, c)
		AuthorRoute(v1, c)
	}
	return r
}

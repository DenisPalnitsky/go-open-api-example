package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	swagger "github.com/num30/gin-swagger-ui"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

type GetRequest struct {
	Name string `path:"name"`
}

type GetResponse struct {
	Result string
}

func main() {
	// Create gin engine
	g := gin.Default()
	g.Use(gin.Logger())

	// Create a new Fizz instance from the Gin engine.
	f := fizz.NewFromEngine(g)

	// Add Open API description
	infos := &openapi.Info{
		Title:       "Service API",
		Description: "This is my Service API",
		Version:     "0.1",
	}

	// Create an endpoint for openapi.json file
	f.GET("/openapi.json", nil, f.OpenAPI(infos, "json"))

	// Now add a UI handler
	swagger.AddOpenApiUIHandler(g, "swagger", "/openapi.json")

	// Add handler
	// Second parameter is additional open API info. It's not required and can be nil
	// Third parameter is handler function. It should be a tonic.Handler in order for it to appear on OpenAPI spec
	f.GET("/hello/:name", []fizz.OperationOption{fizz.Summary("Get a greeting")}, tonic.Handler(func(c *gin.Context, req *GetRequest) (*GetResponse, error) {
		return &GetResponse{Result: "Hello " + req.Name}, nil
	}, http.StatusOK))

	// run our server
	f.Engine().Run(":8080")
}

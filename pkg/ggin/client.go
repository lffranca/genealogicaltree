package ggin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func New(options Options) (*Client, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	client := new(Client)
	client.common.Client = client
	client.specPath = options.SpecPath
	client.person = (*personService)(&client.common)

	client.app = gin.Default()
	client.routes()

	return client, nil
}

type Client struct {
	common   service
	app      *gin.Engine
	specPath *string
	person   *personService
}

func (pkg *Client) routes() {
	docPath := "/swagger.yaml"

	pkg.app.StaticFile(docPath, *pkg.specPath)

	swaggerOpts := middleware.SwaggerUIOpts{
		Path:    "swagger",
		SpecURL: docPath,
	}
	swaggerM := middleware.SwaggerUI(swaggerOpts, nil)
	pkg.app.GET("/swagger", gin.WrapH(swaggerM))

	opts := middleware.RedocOpts{SpecURL: docPath}
	docM := middleware.Redoc(opts, nil)
	pkg.app.GET("/docs", gin.WrapH(docM))

	v1 := pkg.app.Group("/api/v1")
	{
		person := v1.Group("/person")
		{
			person.GET("", pkg.person.listGET)
		}
	}
}

func (pkg *Client) Run(addr ...string) error {
	return pkg.app.Run(addr...)
}

type service struct {
	Client *Client
}

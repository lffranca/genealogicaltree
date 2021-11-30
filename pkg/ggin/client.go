package ggin

import "github.com/gin-gonic/gin"

func New(options Options) (*Client, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	client := new(Client)
	client.common.Client = client
	client.person = (*personService)(&client.common)

	client.app = gin.Default()
	client.routes()

	return client, nil
}

type Client struct {
	common service
	app    *gin.Engine
	person *personService
}

func (pkg *Client) routes() {
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

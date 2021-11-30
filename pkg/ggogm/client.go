package ggogm

import "github.com/mindstand/gogm/v2"

func New(options Options) (*Client, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	gogmClient, err := connection(
		options.Host,
		options.Port,
		options.Username,
		options.Password,
	)
	if err != nil {
		return nil, err
	}

	client := new(Client)
	client.common.Client = client
	client.gogm = gogmClient
	client.Init = (*initService)(&client.common)
	client.Person = (*PersonService)(&client.common)

	return client, nil
}

type Client struct {
	common service
	gogm   *gogm.Gogm
	Init   *initService
	Person *PersonService
}

func (pkg *Client) Close() error {
	return pkg.gogm.Close()
}

type service struct {
	Client *Client
}

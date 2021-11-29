package ggogm

import (
	"github.com/lffranca/genealogicaltree/pkg/ggogm/model"
	"github.com/mindstand/gogm/v2"
)

func connection(host *string, port *int, username, password *string) (*gogm.Gogm, error) {
	var user string
	if username != nil {
		user = *username
	}

	var pass string
	if password != nil {
		pass = *password
	}

	config := &gogm.Config{
		Host:     *host,
		Port:     *port,
		Username: user,
		Password: pass,
		Protocol: "bolt",
		PoolSize: 50,
		Logger:   gogm.GetDefaultLogger(),
		LogLevel: "DEBUG",
	}

	gogmClient, err := gogm.New(config, gogm.DefaultPrimaryKeyStrategy, &model.Person{})
	if err != nil {
		return nil, err
	}

	return gogmClient, nil
}

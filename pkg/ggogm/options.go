package ggogm

import "errors"

type Options struct {
	Host     *string
	Port     *int
	Username *string
	Password *string
}

func (pkg *Options) validate() error {
	if pkg.Host == nil {
		return errors.New("host param is required")
	}

	if pkg.Port == nil {
		return errors.New("port param is required")
	}

	return nil
}

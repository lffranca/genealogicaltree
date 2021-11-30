package ggin

import "errors"

type Options struct {
	SpecPath         *string
	PersonRepository PersonRepository
}

func (pkg *Options) validate() error {
	if pkg.SpecPath == nil {
		return errors.New("spec path is required param")
	}

	if pkg.PersonRepository == nil {
		return errors.New("PersonRepository is required param")
	}

	return nil
}

package ggin

import "errors"

type Options struct {
	SpecPath *string
}

func (pkg *Options) validate() error {
	if pkg.SpecPath == nil {
		return errors.New("spec path is required param")
	}

	return nil
}

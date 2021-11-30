package presenter

type Relationship struct {
	ID           *int    `json:"id" yaml:"id" xml:"id"`
	Name         *string `json:"name" yaml:"name" xml:"name"`
	Relationship string  `json:"relationship" yaml:"relationship" xml:"relationship"`
}

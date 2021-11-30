package presenter

type Person struct {
	ID            *int            `json:"id" yaml:"id" xml:"id"`
	Name          *string         `json:"name" yaml:"name" xml:"name"`
	Relationships []*Relationship `json:"relationships" yaml:"relationships" xml:"relationships"`
}

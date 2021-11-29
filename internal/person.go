package internal

type Person struct {
	ID            *int             `json:"id" yaml:"id" xml:"id"`
	Name          *string          `json:"name" yaml:"name" xml:"name"`
	Parent        []*Person        `json:"parent,omitempty" yaml:"parent,omitempty" xml:"parent,omitempty"`
	Children      []*Person        `json:"children,omitempty" yaml:"children,omitempty" xml:"children,omitempty"`
	Relationships map[string][]int `json:"relationships,omitempty" yaml:"relationships,omitempty" xml:"relationships,omitempty"`
}

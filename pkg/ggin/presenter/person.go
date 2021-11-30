package presenter

import "github.com/lffranca/genealogicaltree/internal"

func PersonToPresenters(items []*internal.Person) (values []*Person) {
	return
}

func PersonToPresenter(item *internal.Person, relationshipsMap map[int]*internal.Person) *Person {
	if item == nil {
		return nil
	}

	person := new(Person)
	person.ID = item.ID
	person.Name = item.Name

	if len(relationshipsMap) > 0 {
		for key, ids := range item.Relationships {
			for _, id := range ids {
				if familyItem, ok := relationshipsMap[id]; ok {
					person.Relationships = append(person.Relationships, &Relationship{
						ID:           familyItem.ID,
						Name:         familyItem.Name,
						Relationship: &key,
					})
				}
			}
		}
	}

	return person
}

type Person struct {
	ID            *int            `json:"id" yaml:"id" xml:"id"`
	Name          *string         `json:"name" yaml:"name" xml:"name"`
	Relationships []*Relationship `json:"relationships" yaml:"relationships" xml:"relationships"`
}

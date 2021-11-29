package model

import (
	"github.com/lffranca/genealogicaltree/internal"
	"github.com/mindstand/gogm/v2"
)

func PersonToEntities(items []*Person) (values []*internal.Person) {
	if items == nil {
		return nil
	}

	for _, item := range items {
		values = append(values, item.entity())
	}

	return
}

func PersonToModels(items []*internal.Person) (values []*Person) {
	if items == nil {
		return nil
	}

	for _, item := range items {
		values = append(values, PersonToModel(item))
	}

	return
}

func PersonToModel(item *internal.Person) *Person {
	if item == nil {
		return nil
	}

	var id *int64
	if item.ID != nil {
		i := int64(*item.ID)
		id = &i
	}

	person := new(Person)
	person.Id = id
	if item.Name != nil {
		person.Name = *item.Name
	}
	if item.Parent != nil {
		person.Parent = PersonToModels(item.Parent)
	}
	if item.Children != nil {
		person.Children = PersonToModels(item.Children)
	}

	return person
}

type Person struct {
	gogm.BaseNode

	Name     string    `gogm:"name=name;unique" json:"name"`
	Parent   []*Person `gogm:"direction=outgoing;relationship=FAMILY" json:"parent"`
	Children []*Person `gogm:"direction=incoming;relationship=FAMILY" json:"children"`
}

func (item *Person) Entity() *internal.Person {
	return item.entity()
}

func (item *Person) EntityWithFamily() *internal.Person {
	person := new(internal.Person)
	if item.Id != nil {
		i := int(*item.Id)
		person.ID = &i
	}
	person.Name = &item.Name
	if item.Parent != nil {
		for _, item2 := range item.Parent {
			if item2.Id != nil {
				if *item2.Id != 0 {
					person.Parent = append(person.Parent, item2.EntityWithFamily())
				}
			}
		}
	}
	if item.Children != nil {
		for _, item2 := range item.Children {
			if item2.Id != nil {
				if *item2.Id != 0 {
					person.Children = append(person.Children, item2.EntityWithFamily())
				}
			}
		}
	}
	if len(item.LoadMap) > 0 {
		person.Relationships = make(map[string][]int)
		for key, value := range item.LoadMap {
			var ids []int
			for _, id := range value.Ids {
				ids = append(ids, int(id))
			}

			person.Relationships[key] = ids
		}
	}

	return person
}

func (item *Person) entity() *internal.Person {
	person := new(internal.Person)
	if item.Id != nil {
		i := int(*item.Id)
		person.ID = &i
	}
	person.Name = &item.Name
	if len(item.LoadMap) > 0 {
		person.Relationships = make(map[string][]int)
		for key, value := range item.LoadMap {
			var ids []int
			for _, id := range value.Ids {
				ids = append(ids, int(id))
			}

			person.Relationships[key] = ids
		}
	}

	return person
}

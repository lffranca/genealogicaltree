package ggogm

import (
	"context"
	"errors"
	"github.com/lffranca/genealogicaltree/internal"
	"github.com/lffranca/genealogicaltree/pkg/ggogm/model"
	"github.com/mindstand/gogm/v2"
	"log"
)

type PersonService service

func (pkg *PersonService) Save(ctx context.Context, item *internal.Person) (*internal.Person, error) {
	sess, err := pkg.Client.gogm.NewSessionV2(gogm.SessionConfig{
		AccessMode: gogm.AccessModeWrite,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := sess.Close(); err != nil {
			log.Println(err)
		}
	}()

	itemModel := model.PersonToModel(item)

	if err := sess.SaveDepth(ctx, itemModel, 20); err != nil {
		return nil, err
	}

	return itemModel.EntityWithFamily(), nil
}

func (pkg *PersonService) All(ctx context.Context) ([]*internal.Person, error) {
	sess, err := pkg.Client.gogm.NewSessionV2(gogm.SessionConfig{
		AccessMode: gogm.AccessModeRead,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := sess.Close(); err != nil {
			log.Println(err)
		}
	}()

	var items []*model.Person
	if err := sess.LoadAll(ctx, &items); err != nil {
		return nil, err
	}

	return model.PersonToEntities(items), nil
}

func (pkg *PersonService) ByID(ctx context.Context, id *int) (*internal.Person, error) {
	if id == nil {
		return nil, errors.New("id param is required")
	}

	sess, err := pkg.Client.gogm.NewSessionV2(gogm.SessionConfig{
		AccessMode: gogm.AccessModeRead,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := sess.Close(); err != nil {
			log.Println(err)
		}
	}()

	var item model.Person
	if err := sess.LoadDepth(ctx, &item, *id, 20); err != nil {
		return nil, err
	}

	return item.Entity(), nil
}

func (pkg *PersonService) ByName(ctx context.Context, name *string) ([]*internal.Person, error) {
	if name == nil {
		return nil, errors.New("name param is required")
	}

	sess, err := pkg.Client.gogm.NewSessionV2(gogm.SessionConfig{
		AccessMode: gogm.AccessModeRead,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := sess.Close(); err != nil {
			log.Println(err)
		}
	}()

	var items []*model.Person
	query := `
		MATCH p=(:Person)<-[:FAMILY]-(:Person {name: $name})<-[:FAMILY*0..2]-(:Person)-[:FAMILY*0..2]->(:Person)
		RETURN p`
	if err := sess.Query(ctx, query, map[string]interface{}{
		"name": *name,
	}, &items); err != nil {
		return nil, err
	}

	return model.PersonToEntities(items), nil
}

func (pkg *PersonService) ShortestPath(ctx context.Context, name1, name2 *string) ([]*internal.Person, error) {
	if name1 == nil || name2 == nil {
		return nil, errors.New("name1 and name2 is required")
	}

	sess, err := pkg.Client.gogm.NewSessionV2(gogm.SessionConfig{
		AccessMode: gogm.AccessModeRead,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := sess.Close(); err != nil {
			log.Println(err)
		}
	}()

	var items []*model.Person
	query := `
		MATCH p=shortestPath((:Person {name: $name1})-[*]-(:Person {name:$name2}))
		RETURN p`
	if err := sess.Query(ctx, query, map[string]interface{}{
		"name1": *name1,
		"name2": *name2,
	}, &items); err != nil {
		return nil, err
	}

	return model.PersonToEntities(items), nil
}

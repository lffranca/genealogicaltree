package ggogm

import (
	"context"
	"github.com/lffranca/genealogicaltree/pkg/ggogm/model"
	"github.com/mindstand/gogm/v2"
	"log"
)

type initService service

func (pkg *initService) InsertData(ctx context.Context) error {
	sess, err := pkg.Client.gogm.NewSessionV2(gogm.SessionConfig{
		AccessMode: gogm.AccessModeWrite,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err := sess.Close(); err != nil {
			log.Println(err)
		}
	}()

	dunny := &model.Person{Name: "Dunny"}
	mike := &model.Person{Name: "Mike"}
	sonny := &model.Person{Name: "Sonny"}
	phoebe := &model.Person{Name: "Phoebe"}
	martin := &model.Person{Name: "Martin"}
	anastasia := &model.Person{Name: "Anastasia"}
	bruce := &model.Person{Name: "Bruce"}
	ursula := &model.Person{Name: "Ursula"}
	jacqueline := &model.Person{Name: "Jacqueline"}
	ellen := &model.Person{Name: "Ellen"}
	eric := &model.Person{Name: "Eric"}
	oprah := &model.Person{Name: "Oprah"}
	melody := &model.Person{Name: "Melody"}
	ariel := &model.Person{Name: "Ariel"}

	dunny.Children = []*model.Person{mike, phoebe}
	mike.Children = []*model.Person{sonny}
	phoebe.Children = []*model.Person{martin, anastasia}
	bruce.Children = []*model.Person{mike, phoebe}
	jacqueline.Children = []*model.Person{ursula, eric}
	ursula.Children = []*model.Person{martin, anastasia}
	eric.Children = []*model.Person{ellen, oprah}
	melody.Children = []*model.Person{eric, ariel}

	if err := sess.SaveDepth(ctx, dunny, 20); err != nil {
		return err
	}

	if err := sess.SaveDepth(ctx, bruce, 20); err != nil {
		return err
	}

	if err := sess.SaveDepth(ctx, jacqueline, 20); err != nil {
		return err
	}

	if err := sess.SaveDepth(ctx, melody, 20); err != nil {
		return err
	}

	return nil
}

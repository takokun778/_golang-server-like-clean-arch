package gateway

import (
	"context"
	"template/app/domain/template"
	"template/ent"
)

type templateGateway struct {
	db *ent.Database
}

func NewTemplateGateway() template.Repository {
	db := ent.DatabaseConnect()
	gataway := new(templateGateway)
	gataway.db = db
	return gataway
}

func (g *templateGateway) Save(ctx context.Context, template *template.Template) (*template.Template, error) {
	_, err := g.db.Template.
		Create().
		SetID(template.Id().Value()).
		SetSample(template.Sample().Value()).
		SetCreatedAt(template.CreatedAt().Value()).
		SetUpdatedAt(template.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return template, nil
}

func (g *templateGateway) Find(ctx context.Context, id template.Id) (*template.Template, error) {
	res, err := g.db.Template.Get(ctx, id.Value())
	if err != nil {
		return nil, err
	}

	props := template.Props{
		Id:        template.NewId(res.ID),
		Sample:    template.NewSample(res.Sample),
		CreatedAt: template.NewTime(res.CreatedAt),
		UpdatedAt: template.NewTime(res.UpdatedAt),
	}

	return template.New(props), nil
}

func (g *templateGateway) FindAll(ctx context.Context) (*template.TemplateList, error) {
	res, err := g.db.Template.
		Query().All(ctx)

	if err != nil {
		return nil, err
	}

	list := make([]*template.Template, 0)

	for _, r := range res {
		props := template.Props{
			Id:        template.NewId(r.ID),
			Sample:    template.NewSample(r.Sample),
			CreatedAt: template.NewTime(r.CreatedAt),
			UpdatedAt: template.NewTime(r.UpdatedAt),
		}
		list = append(list, template.New(props))
	}

	return template.NewList(list), nil
}

func (g *templateGateway) Update(ctx context.Context, template *template.Template) (*template.Template, error) {
	_, err := g.db.Template.
		Update().
		SetSample(template.Sample().Value()).
		SetUpdatedAt(template.UpdatedAt().Value()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return template, nil
}

func (g *templateGateway) Delete(ctx context.Context, id template.Id) error {
	err := g.db.Template.DeleteOneID(id.Value()).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

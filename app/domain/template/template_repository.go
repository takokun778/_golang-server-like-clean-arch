package template

import "context"

type Repository interface {
	Save(ctx context.Context, template *Template) (*Template, error)
	Find(ctx context.Context, id Id) (*Template, error)
	FindAll(ctx context.Context) (*TemplateList, error)
	Update(ctx context.Context, template *Template) (*Template, error)
	Delete(ctx context.Context, id Id) error
}

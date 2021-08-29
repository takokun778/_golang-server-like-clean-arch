package template

import (
	"context"

	ae "template/app/domain/error"
)

type Usecase interface {
	Create(ctx context.Context, input CreateUsecaseInput) (output *CreateUsecaseOutput, err *ae.Error)
	Get(ctx context.Context, input GetUsecaseInput) (output *GetUsecaseOutput, err *ae.Error)
	GetAll(ctx context.Context, input GetAllUsecaseInput) (output *GetAllUsecaseOutput, err *ae.Error)
	Update(ctx context.Context, input UpdateUsecaseInput) (output *UpdateUsecaseOutput, err *ae.Error)
	Delete(ctx context.Context, input DeleteUsecaseInput) (output *DeleteUsecaseOutput, err *ae.Error)
}

type CreateUsecaseInput struct {
	Sample Sample
}

type CreateUsecaseOutput struct {
	Template *Template
}

type GetUsecaseInput struct {
	Id Id
}

type GetUsecaseOutput struct {
	Template *Template
}

type GetAllUsecaseInput struct {
}

type GetAllUsecaseOutput struct {
	TemplateList *TemplateList
}

type UpdateUsecaseInput struct {
	Id     Id
	Sample Sample
}

type UpdateUsecaseOutput struct {
	Template *Template
}

type DeleteUsecaseInput struct {
	Id Id
}

type DeleteUsecaseOutput struct {
	Template *Template
}

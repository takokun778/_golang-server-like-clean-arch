package usecase

import "template/app/domain/template"

type templateUsecase struct {
	templateRepository template.Repository
}

func NewTemplateUsecase(templateRepository template.Repository) template.Usecase {
	usecase := new(templateUsecase)
	usecase.templateRepository = templateRepository
	return usecase
}

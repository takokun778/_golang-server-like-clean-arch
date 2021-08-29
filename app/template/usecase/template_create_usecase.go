package usecase

import (
	"context"
	ae "template/app/domain/error"
	"template/app/domain/template"
	"time"
)

func (u *templateUsecase) Create(ctx context.Context, input template.CreateUsecaseInput) (output *template.CreateUsecaseOutput, err *ae.Error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	props := template.CreateNewProps(input)
	result := template.CreateNew(props)

	result, repoErr := u.templateRepository.Save(timeOutContext, result)
	if repoErr != nil {
		return nil, ae.NewInternalServerError(repoErr, "")
	}

	output = new(template.CreateUsecaseOutput)
	output.Template = result
	return output, nil
}

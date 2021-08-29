package usecase

import (
	"context"
	ae "template/app/domain/error"
	"template/app/domain/template"
	"time"
)

func (u *templateUsecase) Get(ctx context.Context, input template.GetUsecaseInput) (output *template.GetUsecaseOutput, err *ae.Error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.templateRepository.Find(timeOutContext, input.Id)
	if repoErr != nil {
		return nil, ae.NewInternalServerError(repoErr, "")
	}

	output = new(template.GetUsecaseOutput)
	output.Template = result
	return output, err
}

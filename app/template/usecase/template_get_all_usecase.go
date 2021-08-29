package usecase

import (
	"context"
	ae "template/app/domain/error"
	"template/app/domain/template"
	"time"
)

func (u *templateUsecase) GetAll(ctx context.Context, input template.GetAllUsecaseInput) (output *template.GetAllUsecaseOutput, err *ae.Error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.templateRepository.FindAll(timeOutContext)
	if repoErr != nil {
		return nil, ae.NewInternalServerError(repoErr, "")
	}

	output = new(template.GetAllUsecaseOutput)
	output.TemplateList = result
	return output, nil
}

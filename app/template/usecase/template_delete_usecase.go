package usecase

import (
	"context"
	ae "template/app/domain/error"
	"template/app/domain/template"
	"time"
)

func (u *templateUsecase) Delete(ctx context.Context, input template.DeleteUsecaseInput) (output *template.DeleteUsecaseOutput, err *ae.Error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.templateRepository.Find(timeOutContext, input.Id)
	if repoErr != nil {
		return nil, ae.NewNotFoundError(repoErr, input.Id.Value().String()+" is not found")
	}

	repoErr = u.templateRepository.Delete(timeOutContext, input.Id)
	if err != nil {
		return nil, ae.NewInternalServerError(repoErr, "")
	}

	output = new(template.DeleteUsecaseOutput)
	output.Template = result
	return output, nil
}

package usecase

import (
	"context"
	ae "template/app/domain/error"
	"template/app/domain/template"
	"time"
)

func (u *templateUsecase) Update(ctx context.Context, input template.UpdateUsecaseInput) (output *template.UpdateUsecaseOutput, err *ae.Error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, repoErr := u.templateRepository.Find(timeOutContext, input.Id)
	if repoErr != nil {
		return nil, ae.NewNotFoundError(repoErr, input.Id.Value().String()+" is not found")
	}

	props := template.UpdateProps{
		Sample: input.Sample,
	}

	result.Update(props)

	result, repoErr = u.templateRepository.Update(timeOutContext, result)
	if repoErr != nil {
		return nil, ae.NewInternalServerError(repoErr, "")
	}

	output = new(template.UpdateUsecaseOutput)
	output.Template = result
	return output, nil
}

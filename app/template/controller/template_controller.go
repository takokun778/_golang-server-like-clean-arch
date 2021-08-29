package controller

import (
	"context"
	"template/app/domain/template"
	ec "template/app/error/controller"
	pbTemplate "template/proto/template"
)

type templateController struct {
	templateUsecase template.Usecase
	pbTemplate.UnimplementedTemplateServiceServer
}

func NewTemplateController(templateUsecase template.Usecase) *templateController {
	return &templateController{
		templateUsecase: templateUsecase,
	}
}

func (c *templateController) Create(ctx context.Context, req *pbTemplate.CreateRequest) (res *pbTemplate.CreateResponse, err error) {
	sample, err := TemplateTranslator{}.ToDomainSample(req.Sample)
	if err != nil {
		return nil, err
	}

	input := template.CreateUsecaseInput{
		Sample: sample,
	}

	output, apperr := c.templateUsecase.Create(ctx, input)
	if apperr != nil {
		return nil, ec.ErrorTranslator{}.Translate(apperr)
	}

	res = new(pbTemplate.CreateResponse)
	res.Template = TemplateTranslator{}.ToProto(output.Template)
	return
}

func (c *templateController) Get(ctx context.Context, req *pbTemplate.GetRequest) (res *pbTemplate.GetResponse, err error) {
	id, err := TemplateTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	input := template.GetUsecaseInput{
		Id: id,
	}

	output, apperr := c.templateUsecase.Get(ctx, input)
	if apperr != nil {
		return nil, ec.ErrorTranslator{}.Translate(apperr)
	}

	res = new(pbTemplate.GetResponse)
	res.Template = TemplateTranslator{}.ToProto(output.Template)
	return
}

func (c *templateController) GetAll(ctx context.Context, req *pbTemplate.GetAllRequest) (res *pbTemplate.GetAllResponse, err error) {
	input := template.GetAllUsecaseInput{}

	output, apperr := c.templateUsecase.GetAll(ctx, input)
	if apperr != nil {
		return nil, ec.ErrorTranslator{}.Translate(apperr)
	}

	res = new(pbTemplate.GetAllResponse)
	res.Template = TemplateTranslator{}.ToProtoList(output.TemplateList)
	return
}

func (c *templateController) Update(ctx context.Context, req *pbTemplate.UpdateRequest) (res *pbTemplate.UpdateResponse, err error) {
	id, err := TemplateTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	sample, err := TemplateTranslator{}.ToDomainSample(req.Sample)
	if err != nil {
		return nil, err
	}

	input := template.UpdateUsecaseInput{
		Id:     id,
		Sample: sample,
	}

	output, apperr := c.templateUsecase.Update(ctx, input)
	if apperr != nil {
		return nil, ec.ErrorTranslator{}.Translate(apperr)
	}

	res = new(pbTemplate.UpdateResponse)
	res.Template = TemplateTranslator{}.ToProto(output.Template)
	return
}

func (c *templateController) Delete(ctx context.Context, req *pbTemplate.DeleteRequest) (res *pbTemplate.DeleteResponse, err error) {
	id, err := TemplateTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	input := template.DeleteUsecaseInput{
		Id: id,
	}

	output, apperr := c.templateUsecase.Delete(ctx, input)
	if apperr != nil {
		return nil, ec.ErrorTranslator{}.Translate(apperr)
	}

	res = new(pbTemplate.DeleteResponse)
	res.Template = TemplateTranslator{}.ToProto(output.Template)

	return
}

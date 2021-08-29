package controller

import (
	"errors"
	"template/app/domain/template"
	pbTemplate "template/proto/template"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TemplateTranslator struct{}

func (TemplateTranslator) ToDomainId(value string) (template.Id, error) {
	id, err := template.ParseId(value)
	if err != nil {
		return template.Id{}, err
	}

	return id, nil
}

func (TemplateTranslator) ToDomainSample(value string) (template.Sample, error) {
	if value == "" {
		return template.Sample{}, errors.New("sample is empty")
	}

	return template.NewSample(value), nil
}

func (TemplateTranslator) ToProto(template *template.Template) *pbTemplate.Template {
	proto := new(pbTemplate.Template)
	proto.Id = template.Id().Value().String()
	proto.Sample = template.Sample().Value()
	proto.CreatedAt = timestamppb.New(template.CreatedAt().Value())
	proto.UpdatedAt = timestamppb.New(template.UpdatedAt().Value())
	return proto
}

func (TemplateTranslator) ToProtoList(templateList *template.TemplateList) []*pbTemplate.Template {
	proto := make([]*pbTemplate.Template, 0)
	for _, template := range templateList.Items() {
		proto = append(proto, TemplateTranslator{}.ToProto(template))
	}
	return proto
}

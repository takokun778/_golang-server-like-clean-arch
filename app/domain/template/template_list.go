package template

type TemplateList struct {
	items []*Template
}

func (t *TemplateList) Items() []*Template {
	return t.items
}

func NewList(items []*Template) *TemplateList {
	res := new(TemplateList)
	res.items = items
	return res
}

func (t *TemplateList) Len() int {
	return len(t.items)
}

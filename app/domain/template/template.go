package template

type Template struct {
	id        Id
	sample    Sample
	createdAt Time
	updatedAt Time
}

func (t *Template) Id() Id {
	return t.id
}

func (t *Template) Sample() Sample {
	return t.sample
}

func (t *Template) CreatedAt() Time {
	return t.createdAt
}

func (t *Template) UpdatedAt() Time {
	return t.updatedAt
}

type Props struct {
	Id        Id
	Sample    Sample
	CreatedAt Time
	UpdatedAt Time
}

func New(props Props) *Template {
	template := new(Template)
	template.id = props.Id
	template.sample = props.Sample
	template.createdAt = props.CreatedAt
	template.updatedAt = props.UpdatedAt
	return template
}

type CreateNewProps struct {
	Sample Sample
}

func CreateNew(props CreateNewProps) *Template {
	template := new(Template)
	template.id = Id{}.Create()
	template.sample = props.Sample
	now := Time{}.Now()
	template.createdAt = now
	template.updatedAt = now
	return template
}

type UpdateProps struct {
	Sample Sample
}

func (t *Template) Update(props UpdateProps) {
	t.sample = props.Sample
	t.updatedAt = Time{}.Now()
}

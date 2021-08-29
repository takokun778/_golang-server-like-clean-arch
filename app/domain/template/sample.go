package template

type Sample struct {
	value string
}

func (s Sample) Value() string {
	return s.value
}

func NewSample(value string) Sample {
	res := new(Sample)
	res.value = value
	return *res
}

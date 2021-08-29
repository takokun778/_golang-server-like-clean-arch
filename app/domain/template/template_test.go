package template_test

import (
	"template/app/domain/template"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	id := template.Id{}.Create()
	sample := template.NewSample("new")
	now := template.Time{}.Now()
	createdAt := now
	updatedAt := now

	props := template.Props{
		Id:        id,
		Sample:    sample,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	result := template.New(props)

	assert.Equal(t, result.Id(), id)
	assert.Equal(t, result.Sample(), sample)
	assert.Equal(t, result.CreatedAt(), createdAt)
	assert.Equal(t, result.UpdatedAt(), updatedAt)
}

func TestNewCreate(t *testing.T) {
	sample := template.NewSample("create")
	props := template.CreateNewProps{
		Sample: sample,
	}

	result := template.CreateNew(props)

	assert.Equal(t, result.Sample(), sample)
	assert.Equal(t, result.CreatedAt(), result.UpdatedAt())
}

func TestUpdate(t *testing.T) {
	props := template.CreateNewProps{
		Sample: template.NewSample("new"),
	}

	result := template.CreateNew(props)

	// 更新時間を比較しやすくするため
	time.Sleep(time.Millisecond)

	sample := template.NewSample("update")
	updateProps := template.UpdateProps{
		Sample: sample,
	}
	result.Update(updateProps)

	assert.Equal(t, result.Sample(), sample)
	assert.Greater(t, result.UpdatedAt().Value().UnixNano(), result.CreatedAt().Value().UnixNano())
}

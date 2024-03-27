package toModule

import (
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/model"
)

func CompDoMainToModule(m *model.Comp, d *domain.Comp) {
	m.CreateId = d.CreateId
	m.Types = d.Types
	m.IsList = d.IsList
	m.Title = d.Title
	m.Info = d.Info
	m.Deploy = d.Deploy
	m.Url = d.Url
	m.PreviewUrl = d.PreviewUrl
	m.ID = d.ID
}

package service

import (
	"github.com/PGshen/go-project/blog/internal/model"
	"github.com/PGshen/go-project/blog/pkg/app"
)

type CountTagRequest struct {
	Name string `from:"name" binding:"max=100"`
	State uint8 `from:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name string `from:"name" binding:"max=100"`
	State uint8 `from:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name string `from:"name" binding:"required,min=3,max=100"`
	CreatedBy string `from:"created_by" binding:"required,min=1,max=100"`
	State uint8 `from:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID uint32 `from:"id" binding:"required,gte=1"`
	Name string `from:"name" binding:"required,min=3,max=100"`
	ModifiedBy string `from:"created_by" binding:"required,min=1,max=100"`
	State uint8 `from:"state,default=1" binding:"oneof=0 1"`
}

type DeleteTagRequest struct {
	ID uint32 `from:"id" binding:"required,gte=1"`
}

func (s *Service) CountTag(param *CountTagRequest) (int, error) {
	return s.dao.CountTag(param.Name, param.State)
}

func (s *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (s *Service) CreateTag(param *CreateTagRequest) error {
	return s.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (s *Service) UpdateTag(param *UpdateTagRequest) error {
	return s.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (s *Service) DeleteTag(param *DeleteTagRequest) error {
	return s.dao.DeleteTag(param.ID)
}
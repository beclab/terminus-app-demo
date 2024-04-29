package storage

import "github.com/beclab/terminus-app-demo/pkg/model"

type Operator interface {
	Get(id int) (*model.Note, error)
	List() ([]*model.Note, error)
	Create(data *model.Note) (*model.Note, error)
	Delete(id int) error
	Update(data *model.Note) (*model.Note, error)
}

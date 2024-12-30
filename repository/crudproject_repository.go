package repository

import (
	"context"
	"crud-project/entity"
)

type ProjectcrudRepository interface {
	Insert(ctx context.Context, projectcrud entity.Crudproject) (entity.Crudproject, error)
	FindByid(ctx context.Context, id int32) (entity.Crudproject, error)
	FindAll(ctx context.Context) ([]entity.Crudproject, error)
	Update(ctx context.Context, projectcrud entity.Crudproject) (entity.Crudproject, error)
	Delete(ctx context.Context, id int32) error
}

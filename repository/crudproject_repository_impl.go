package repository

import (
	"context"
	"crud-project/entity"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type crudprojectRepositoryImpl struct {
	DB *sql.DB
}

func NewCrudProjectRepository(db *sql.DB) *crudprojectRepositoryImpl {
	return &crudprojectRepositoryImpl{
		DB: db,
	}
}

func (repository *crudprojectRepositoryImpl) Insert(ctx context.Context, projectcrud entity.Crudproject) (entity.Crudproject, error) {
	script := "INSERT INTO crudproject(name,task) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, projectcrud.Name, projectcrud.Task)
	if err != nil {
		return projectcrud, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return projectcrud, err
	}
	projectcrud.Id = int32(id)
	// projectcrud.Done = false
	return projectcrud, nil
}

func (repository *crudprojectRepositoryImpl) FindByid(ctx context.Context, id int32) (entity.Crudproject, error) {
	script := "SELECT id, name, done, task FROM crudproject WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	projectcrud := entity.Crudproject{}
	if err != nil {
		return projectcrud, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&projectcrud.Id, &projectcrud.Name, &projectcrud.Done, &projectcrud.Task)
		return projectcrud, nil
	} else {
		return projectcrud, errors.New("id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository *crudprojectRepositoryImpl) FindAll(ctx context.Context) ([]entity.Crudproject, error) {
	script := "SELECT id, name , done , task FROM crudproject"
	rows, err := repository.DB.QueryContext(ctx, script)
	var projectcruds []entity.Crudproject
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		projectcrud := entity.Crudproject{}
		rows.Scan(&projectcrud.Id, &projectcrud.Name, &projectcrud.Done, &projectcrud.Task)
		projectcruds = append(projectcruds, projectcrud)
	}
	return projectcruds, nil
}

// func (repository *crudprojectRepositoryImpl) Update(ctx context.Context, projectcrud entity.Crudproject) (entity.Crudproject, error) {
// 	script := "UPDATE projectcrud SET done = ?, task = ? WHERE id = ?"
// 	result, err := repository.DB.ExecContext(ctx, script, projectcrud.Done, projectcrud.Task, projectcrud.Id)
// 	if err != nil {
// 		return entity.Crudproject{}, err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return entity.Crudproject{}, err
// 	}
// 	if rowsAffected == 0 {
// 		return entity.Crudproject{}, errors.New("data not found")
// 	}
// 	return projectcrud, nil
// }
//

func (repository *crudprojectRepositoryImpl) Update(ctx context.Context, projectcrud entity.Crudproject) (entity.Crudproject, error) {
	// Correct table name and include all fields
	script := "UPDATE crudproject SET name = ?, done = ?, task = ? WHERE id = ?"

	// Correct parameter order and include all fields
	result, err := repository.DB.ExecContext(ctx, script,
		projectcrud.Name, // First parameter
		projectcrud.Done, // Second parameter
		projectcrud.Task, // Third parameter
		projectcrud.Id,   // Fourth parameter (WHERE clause)
	)
	if err != nil {
		return entity.Crudproject{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entity.Crudproject{}, err
	}

	if rowsAffected == 0 {
		// Check if record exists
		var count int
		checkQuery := "SELECT COUNT(*) FROM crudproject WHERE id = ?"
		err := repository.DB.QueryRowContext(ctx, checkQuery, projectcrud.Id).Scan(&count)
		if err != nil {
			return entity.Crudproject{}, err
		}

		if count == 0 {
			return entity.Crudproject{}, fmt.Errorf("record with ID %d not found", projectcrud.Id)
		}

		return entity.Crudproject{}, errors.New("no changes made to the record")
	}

	return projectcrud, nil
}

func (repository *crudprojectRepositoryImpl) Delete(ctx context.Context, id int32) error {
	panic("implement me")
}

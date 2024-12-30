package repository

import (
	"context"
	"crud-project"
	"crud-project/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectCrudinsert(t *testing.T) {
	projectCrudRepository := NewCrudProjectRepository(crud.GetConnection())
	ctx := context.Background()
	projectCrud := entity.Crudproject{
		Name: "nabiel 2",
		Task: "belajar kalkulus jam 2.30",
	}
	result, err := projectCrudRepository.Insert(ctx, projectCrud)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestProjectCrudFindById(t *testing.T) {
	ProjectcrudRepository := NewCrudProjectRepository(crud.GetConnection())
	projectcrud, err := ProjectcrudRepository.FindByid(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(projectcrud)
}

func TestProjectCrudFindAll(t *testing.T) {
	ctx := context.Background()
	ProjectcrudRepository := NewCrudProjectRepository(crud.GetConnection())
	projectcruds, err := ProjectcrudRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, projectcrud := range projectcruds {
		fmt.Println(projectcrud)
	}
}

func TestProjectCrudUpdate(t *testing.T) {
	// Persiapan
	RepositoryProjectcrud := NewCrudProjectRepository(crud.GetConnection())
	ctx := context.Background()

	// Siapkan data uji
	projectcrud := entity.Crudproject{
		Id:   8, // Pastikan ID ini ada di database Anda
		Name: "nabiel 2",
		Done: true,
		Task: "uji coba pembaruan jam 10 malam",
	}

	// Lakukan pembaruan
	hasil, err := RepositoryProjectcrud.Update(ctx, projectcrud)

	// Periksa kesalahan
	assert.NoError(t, err, "Pembaruan tidak boleh menghasilkan kesalahan")

	// Validasi hasil yang dikembalikan
	assert.NotNil(t, hasil, "Hasil tidak boleh kosong")
	assert.Equal(t, projectcrud.Id, hasil.Id, "ID harus sesuai")
	assert.Equal(t, projectcrud.Name, hasil.Name, "Nama harus sesuai")
	assert.Equal(t, projectcrud.Done, hasil.Done, "Status selesai harus sesuai")
	assert.Equal(t, projectcrud.Task, hasil.Task, "Tugas harus sesuai")

	// Opsional: Verifikasi pembaruan di database
	proyekTerbarui, err := RepositoryProjectcrud.FindByid(ctx, projectcrud.Id)
	assert.NoError(t, err, "Harus dapat menemukan proyek yang diperbarui")
	assert.Equal(t, projectcrud.Name, proyekTerbarui.Name, "Nama di database harus sesuai dengan nama yang diperbarui")
	assert.Equal(t, projectcrud.Done, proyekTerbarui.Done, "Status selesai di database harus sesuai")
	assert.Equal(t, projectcrud.Task, proyekTerbarui.Task, "Tugas di database harus sesuai")

	// Cetak hasil untuk visibilitas tambahan
	t.Logf("Proyek Diperbarui: %+v", hasil)
}



func TestProjectCrudDelete(t *testing.T)  {
  ctx := context.Background()
  repository := NewCrudProjectRepository(crud.GetConnection())
  err := repository.Delete(ctx, 2)
  if err != nil {
    panic(err)
  }
  fmt.Println("done delete")
}

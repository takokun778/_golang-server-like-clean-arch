package controller_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"template/ent"
	pbTemplate "template/proto/template"
	"testing"

	tc "template/app/template/controller"
	tg "template/app/template/gateway"
	tu "template/app/template/usecase"

	"github.com/stretchr/testify/assert"
)

var ctx = context.TODO()
var gateway = tg.NewTemplateGateway()
var usecase = tu.NewTemplateUsecase(gateway)
var controller = tc.NewTemplateController(usecase)

func TestMain(m *testing.M) {
	// テスト前処理
	createTable()

	// テスト実行
	m.Run()

	// テスト後処理
	defer dropTable()
}

// xxx -> result
// yyy -> comparison
// xxxがyyyと一致する
// assert.Equal(t, xxx, yyy)
// xxxがyyyよりも大きい
// assert.Greater(t, xxx, yyy)

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := new(pbTemplate.CreateRequest)
		req.Sample = "create"

		res, err := controller.Create(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Template.Sample, "create")

		resetTable()
	})
}

func TestGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbTemplate.CreateRequest)
		dataReq.Sample = "get"

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbTemplate.GetRequest)
		req.Id = dataRes.Template.Id

		res, err := controller.Get(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Template.Sample, "get")
		assert.Equal(t, res.Template.Id, req.Id)
		assert.Equal(t, res.Template.CreatedAt, dataRes.Template.CreatedAt)
		assert.Equal(t, res.Template.UpdatedAt, dataRes.Template.UpdatedAt)

		resetTable()
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq1 := new(pbTemplate.CreateRequest)
		dataReq1.Sample = "all"

		_, err := controller.Create(ctx, dataReq1)
		assert.NoError(t, err)

		dataReq2 := new(pbTemplate.CreateRequest)
		dataReq2.Sample = "all"

		_, err = controller.Create(ctx, dataReq2)
		assert.NoError(t, err)

		req := new(pbTemplate.GetAllRequest)

		res, err := controller.GetAll(ctx, req)

		assert.NoError(t, err)
		assert.Len(t, res.Template, 2)

		resetTable()
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbTemplate.CreateRequest)
		dataReq.Sample = "src"

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbTemplate.UpdateRequest)
		req.Id = dataRes.Template.Id
		req.Sample = "dst"

		res, err := controller.Update(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Template.Id, req.Id)
		assert.Equal(t, res.Template.Sample, req.Sample)
		assert.Equal(t, res.Template.CreatedAt, dataRes.Template.CreatedAt)
		assert.Greater(t, res.Template.UpdatedAt.Nanos, dataRes.Template.UpdatedAt.Nanos)

		resetTable()
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dataReq := new(pbTemplate.CreateRequest)
		dataReq.Sample = "src"

		dataRes, err := controller.Create(ctx, dataReq)
		assert.NoError(t, err)

		req := new(pbTemplate.DeleteRequest)
		req.Id = dataRes.Template.Id

		res, err := controller.Delete(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, res.Template.Id, req.Id)

		getReq := new(pbTemplate.GetRequest)
		getReq.Id = res.Template.Id

		getRes, err := controller.Get(ctx, getReq)
		assert.Nil(t, getRes)
		assert.Error(t, err)

		resetTable()
	})
}

func createTable() {
	database := ent.DatabaseConnect()

	ctx := context.TODO()

	if err := database.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("create table completed")
}

func resetTable() {
	datasource := fmt.Sprintf("host=%s port=%s user=%s dbname= %s password=%s sslmode=disable", "localhost", "54321", "test", "test", "test")

	database, err := sql.Open("postgres", datasource)

	if err != nil {
		log.Fatal(err)
	}

	database.Exec("DELETE FROM templates")
}

func dropTable() {
	datasource := fmt.Sprintf("host=%s port=%s user=%s dbname= %s password=%s sslmode=disable", "localhost", "54321", "test", "test", "test")

	database, err := sql.Open("postgres", datasource)

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	database.Exec("DROP TABLE IF EXISTS templates")

	log.Println("drop table completed")
}

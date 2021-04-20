package chiu

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"upper.io/db.v3/lib/sqlbuilder"
)

// CrudTable 自动生成针对指定表的增删查改
func CrudTable(r chi.Router, sess sqlbuilder.Database, table string, prefix string) {
	var controller = CrudController{
		Sess:      sess,
		GetParam:  chi.URLParam,
		TableName: table,
	}
	r.Route(prefix, func(r chi.Router) {
		r.Get("/", controller.Pagination)
		r.Get("/{id}", controller.One)
		r.Post("/", controller.Create)
		r.Put("/", controller.Update)
		r.Delete("/{id}", controller.Delete)
		r.Get("/all", controller.All)
		r.Post("/update", controller.Update)
	})
	logrus.Info(fmt.Sprintf("auto mount: [%s => %s]", table, prefix))
}

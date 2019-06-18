# gocrud

`CrudController`

usage:
```go
var db sqlbuilder.Database // https://upper.io/db.v3
var err error

db, err = mysql.Open(settings)
if err != nil {
    logrus.Fatal("数据库连接失败, ", err)
}

var categoryController = gocrud.CrudController{
    Sess:      db,
    GetParam:  chi.URLParam,
    Columns:   []string{"id", "name"},
    TableName: "gl_category",
}

r.Route("/categories", func(r chi.Router) {
    r.Get("/", categoryController.Pagination)
    r.Get("/{id}", categoryController.One)
    r.Post("/", categoryController.Create)
    r.Put("/", categoryController.Update)
    r.Delete("/{id}", categoryController.Delete)
    r.Get("/all", categoryController.All)
})
```

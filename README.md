# gocrud

`controllers.CrudController`

usage:
```go
var db sqlbuilder.Database

var categoryController = controllers.CrudController{
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

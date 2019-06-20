# gocrud

## Usage:
```go
var r = chi.NewRouter()
var sess sqlbuilder.Database
```

`CrudController`
```go
var categoryController = gocrud.CrudController{
    Sess:      sess,
    GetParam:  chi.URLParam,
    Columns:   []string{"id", "name"},
    TableName: "category",
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

`CrudDb*`
```go
gocrud.CrudDBAll(r, sess)
gocrud.CrudDBOnly(r, sess, []string{"demo"})
gocrud.CrudDBExcept(r, sess, []string{"demo"})
```

`CrudTable`
```go
gocrud.CrudTable(r, sess, "demo", "/hi/demo")
```

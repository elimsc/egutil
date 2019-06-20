# gocrud

## Usage:
```go
var r = chi.NewRouter()
var sess sqlbuilder.Database  // see https://upper.io/db.v3
```

`CrudController`
```go
var categoryController = gocrud.CrudController{
    Sess:      sess,
    GetParam:  chi.URLParam,
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

`CrudTable`
```go
gocrud.CrudTable(r, sess, "demo", "/hi/demo")
// GET /hi/demo       list 
// POST /hi/demo      create
// GET /hi/demo/{id}  one
// PUT /hi/demo      update
// GET /hi/demo/all  all
// POST /hi/update   update
```

`CrudDb*`
```go
gocrud.CrudDBAll(r, sess)
gocrud.CrudDBOnly(r, sess, []string{"demo"})
gocrud.CrudDBExcept(r, sess, []string{"demo"})
// for _, table := range tables {
//    gocrud.CrudTable(r, sess, table, "/"+table)
// }
```



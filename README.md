# egu

## Usage:
```go
var r = chi.NewRouter()
var sess sqlbuilder.Database  // see https://upper.io/db.v3
```

`CrudController`
```go
var categoryController = egu.CrudController{
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
egu.CrudTable(r, sess, "demo", "/hi/demo")
// GET /hi/demo       list 
// POST /hi/demo      create
// GET /hi/demo/{id}  one
// PUT /hi/demo      update
// GET /hi/demo/all  all
// POST /hi/update   update
```

`CrudDb*`
```go
egu.CrudDBAll(r, sess)
egu.CrudDBOnly(r, sess, []string{"demo"})
egu.CrudDBExcept(r, sess, []string{"demo"})
// for _, table := range tables {
//    egu.CrudTable(r, sess, table, "/"+table)
// }
```

`FileServer`
```go
workDir, _ := os.Getwd()
filesDir := filepath.Join(workDir, "files")
egu.FileServer(r, "/files", http.Dir(filesDir))

// same as
egu.ServeStatic(r, "/files", "files")
```



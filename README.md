# goe

> Go extend.

```
$ go get https://github.com/elimsc/goe
```

## chie:
```go
var r = chi.NewRouter()
var sess sqlbuilder.Database  // see https://upper.io/db.v3
```

`chie.CrudController`
```go
var categoryController = chie.CrudController{
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

`chie.CrudTable`
```go
chie.CrudTable(r, sess, "demo", "/hi/demo")
// GET /hi/demo       list 
// POST /hi/demo      create
// GET /hi/demo/{id}  one
// PUT /hi/demo      update
// GET /hi/demo/all  all
// POST /hi/update   update
```

`chie.CrudDb*`
```go
chie.CrudDBAll(r, sess)
chie.CrudDBOnly(r, sess, []string{"demo"})
chie.CrudDBExcept(r, sess, []string{"demo"})
// for _, table := range tables {
//    chiu.CrudTable(r, sess, table, "/"+table)
// }
```

`chie.FileServer`
```go
workDir, _ := os.Getwd()
filesDir := filepath.Join(workDir, "files")
chie.FileServer(r, "/files", http.Dir(filesDir))

// same as
chie.ServeStatic(r, "/files", "files")
```



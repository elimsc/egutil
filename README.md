# egu

```
$ go get https://github.com/elimsc/egu
```

## chiu:
```go
var r = chi.NewRouter()
var sess sqlbuilder.Database  // see https://upper.io/db.v3
```

`chiu.CrudController`
```go
var categoryController = chiu.CrudController{
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

`chiu.CrudTable`
```go
chiu.CrudTable(r, sess, "demo", "/hi/demo")
// GET /hi/demo       list 
// POST /hi/demo      create
// GET /hi/demo/{id}  one
// PUT /hi/demo      update
// GET /hi/demo/all  all
// POST /hi/update   update
```

`chiu.CrudDb*`
```go
chiu.CrudDBAll(r, sess)
chiu.CrudDBOnly(r, sess, []string{"demo"})
chiu.CrudDBExcept(r, sess, []string{"demo"})
// for _, table := range tables {
//    chiu.CrudTable(r, sess, table, "/"+table)
// }
```

`chiu.FileServer`
```go
workDir, _ := os.Getwd()
filesDir := filepath.Join(workDir, "files")
chiu.FileServer(r, "/files", http.Dir(filesDir))

// same as
chiu.ServeStatic(r, "/files", "files")
```



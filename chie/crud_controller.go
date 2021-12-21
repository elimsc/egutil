package chie

import (
	"encoding/json"
	"net/http"
	"strconv"

	"upper.io/db.v3/lib/sqlbuilder"

	db "upper.io/db.v3"
)

// CrudController Crud操作集合
type CrudController struct {
	GetParam  func(*http.Request, string) string
	Sess      sqlbuilder.Database
	TableName string
}

// Table 表操作
func (c *CrudController) Table() db.Collection {
	return c.Sess.Collection(c.TableName)
}

// JSON 返回json结果
func (c *CrudController) JSON(w http.ResponseWriter, data interface{}) {
	jsonvalue, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonvalue)
}

// All 一次性获取所有数据 orderby=-id
func (c *CrudController) All(w http.ResponseWriter, r *http.Request) {
	var err error
	var orderby = "-id"
	if r.FormValue("orderby") != "" {
		orderby = r.FormValue("orderby")
	}

	var result = []map[string]*string{}
	err = c.Table().Find().OrderBy(orderby).All(&result)
	if err != nil {
		panic(err)
	}

	c.JSON(w, result)
}

// Pagination 分页显示 , 默认offset=0, limit=10, where={"key": "value"}, orderby=-id
func (c *CrudController) Pagination(w http.ResponseWriter, r *http.Request) {
	var total uint64
	var err error

	var (
		offset  int
		limit   int
		orderby string
	)

	var result = []map[string]*string{}

	orderby = "-id"
	offset, err = strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		offset = 0
	}
	limit, err = strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		limit = 10
	}

	// 构造条件
	var condition = db.Cond{}
	var where = map[string]interface{}{}
	err = json.Unmarshal([]byte(r.FormValue("where")), &where) // 忽略错误
	for k, v := range where {
		condition[k] = v
	}

	if r.FormValue("orderby") != "" {
		orderby = r.FormValue("orderby")
	}

	err = c.Table().Find().Where(condition).OrderBy(orderby).Limit(limit).Offset(offset).All(&result)
	if err != nil {
		panic(err)
	}
	total, err = c.Table().Find().Where(condition).Count()
	if err != nil {
		panic(err)
	}

	c.JSON(w, map[string]interface{}{
		"total":  total,
		"offset": offset,
		"list":   result,
		"limit":  limit,
	})
}

// One 根据id获取单个实例
func (c *CrudController) One(w http.ResponseWriter, r *http.Request) {
	var err error
	var result = map[string]*string{}

	var id = c.GetParam(r, "id")

	err = c.Table().Find().Where(db.Cond{"id": id}).One(&result)
	if err != nil {
		panic(err)
	}

	c.JSON(w, result)
}

// Create 新增实例
func (c *CrudController) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	var model = map[string]interface{}{}

	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		panic(err)
	}

	_, err = c.Table().Insert(&model)
	if err != nil {
		panic(err)
	}

	c.JSON(w, map[string]interface{}{
		"code": 0,
	})
}

// Update 更新实例
func (c *CrudController) Update(w http.ResponseWriter, r *http.Request) {
	var err error
	var model = map[string]interface{}{}

	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		panic(err)
	}

	var id = model["id"]
	delete(model, "id")

	_, err = c.Sess.Update(c.TableName).Set(model).Where(db.Cond{"id": id}).Exec()
	if err != nil {
		panic(err)
	}

	c.JSON(w, map[string]interface{}{
		"code": 0,
	})
}

// Delete 通过id删除实例
func (c *CrudController) Delete(w http.ResponseWriter, r *http.Request) {
	var err error
	var id = c.GetParam(r, "id")

	_, err = c.Sess.DeleteFrom(c.TableName).Where(db.Cond{"id": id}).Exec()
	if err != nil {
		panic(err)
	}

	c.JSON(w, map[string]interface{}{
		"code": 0,
	})
}

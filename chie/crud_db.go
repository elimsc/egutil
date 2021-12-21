package chie

import (
	"github.com/go-chi/chi"
	"upper.io/db.v3/lib/sqlbuilder"
)

// CrudDBAll 自动生成针对于所有表的增删查改
// `GET /table`  分页结果
// `GET /table/{id}``  根据id获取结果
// `POST /table`  新增记录
// `PUT /table`  更新记录
// `DELETE /table/{id}` 删除记录
// `GET /table/all` 所有结果
func CrudDBAll(r chi.Router, sess sqlbuilder.Database) {
	var tables, err = sess.Collections()
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		CrudTable(r, sess, table, "/"+table)
	}
}

// CrudDBOnly 自动生成针对于指定表的增删查改
func CrudDBOnly(r chi.Router, sess sqlbuilder.Database, only []string) {
	var tables, err = sess.Collections()
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		for _, tableNeed := range only {
			if table == tableNeed {
				CrudTable(r, sess, table, "/"+table)
				continue
			}
		}
	}
}

// CrudDBExcept 自动生成针对于排除指定表外所有表的增删查改
func CrudDBExcept(r chi.Router, sess sqlbuilder.Database, except []string) {
	var tables, err = sess.Collections()
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		for _, tableExcept := range except {
			if table != tableExcept {
				CrudTable(r, sess, table, "/"+table)
				continue
			}
		}
	}
}

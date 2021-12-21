package httpe

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

// BaseController is the base for controller
type BaseController struct{}

// ParseBody JSON解析
func (c *BaseController) ParseBody(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

// MustParseBody if err happens, panic
func (c *BaseController) MustParseBody(r *http.Request, data interface{}) {
	err := c.ParseBody(r, data)
	if err != nil {
		panic(err)
	}
}

// JSON 返回json结果
func (c *BaseController) JSON(w http.ResponseWriter, data interface{}) {
	jsonvalue, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonvalue)
}

// SaveUploadFile 保存上传文件
func (c *BaseController) SaveUploadFile(r *http.Request, reqName string, dir string) (string, error) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(reqName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	now := time.Now()
	// 是否可能存在文件名过长的问题?
	filename := path.Join(dir, fmt.Sprintf("%d-%s", now.Unix(), handler.Filename))
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	return filename, err
}

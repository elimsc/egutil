package fastutil

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasttemplate"
)

// RenderHTML output html to browser
func RenderHTML(ctx *fasthttp.RequestCtx, view string, data map[string]interface{}) {
	var template, err = ioutil.ReadFile(path.Join("views", view))
	if err != nil {
		logrus.Error(err)
		fmt.Fprint(ctx, "404 page not found.")
	}

	var t = fasttemplate.New(string(template), "{{", "}}")
	var html = t.ExecuteString(data)

	ctx.SetContentType("html")
	ctx.SetBodyString(html)
}

package fastutil

import (
	"github.com/valyala/fasthttp"
)

// RenderJSON output html to browser
func RenderJSON(ctx *fasthttp.RequestCtx, data interface{}) {
	var json = ""

	ctx.SetContentType("application/json")
	ctx.SetBodyString(json)
}

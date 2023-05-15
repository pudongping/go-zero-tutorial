package httpRest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func NotFoundHandler() http.Handler {
	return http.HandlerFunc(NotFound)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	// 避免请求网站图标出现 404
	if strings.HasPrefix(r.URL.Path, "/favicon.ico") {
		w.WriteHeader(http.StatusOK)
		return
	}

	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("页面无法找到 ...(｡•ˇ‸ˇ•｡) ..."))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	resp := map[string]interface{}{
		"code": http.StatusNotFound,
		"msg":  fmt.Sprintf("路由[%s]未定义", r.URL.Path),
	}
	bs, _ := json.Marshal(resp)
	w.Write(bs)
}

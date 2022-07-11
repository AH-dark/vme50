package controller

import (
	"github.com/AH-dark/random-donate/bootstrap"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"
)

func FrontendHandler() gin.HandlerFunc {
	ignoreFunc := func(c *gin.Context) {
		c.Next()
	}
	if bootstrap.StaticFS == nil {
		return ignoreFunc
	}

	// Check /index.html
	file, err := bootstrap.StaticFS.Open("/index.html")
	if err != nil {
		utils.Log().Warning("静态文件[index.html]不存在，可能会影响首页展示")
		return ignoreFunc
	}

	fileContentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		utils.Log().Warning("静态文件[index.html]读取失败，可能会影响首页展示")
		return ignoreFunc
	}
	indexFileContent := string(fileContentBytes)
	fileServer := http.FileServer(bootstrap.StaticFS)

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// API 跳过
		if strings.HasPrefix(path, "/api") || path == "/manifest.json" {
			c.Next()
			return
		}

		// index.html
		if (path == "/") || !bootstrap.StaticFS.Exists("/", path) {
			c.Header("Content-Type", "text/html")
			c.String(200, indexFileContent)
			c.Abort()
			return
		}

		// react statics
		if strings.HasPrefix(path, "/static/") {
			file, err := bootstrap.StaticFS.Open(path)
			if err != nil {
				c.String(500, "server error")
			}

			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				c.String(500, "server error")
			}

			arr := strings.Split(path, ".")
			ext := arr[len(arr)-1]

			c.Header("Content-Type", mime.TypeByExtension("."+ext))
			c.String(200, string(bytes))
			c.Abort()
			return
		}

		// next.js translations
		if strings.HasPrefix(path, "/locales/") {
			file, err := bootstrap.StaticFS.Open(strings.ToLower(path))
			if err != nil {
				c.String(500, "server error")
			}

			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				c.String(500, "server error")
			}

			c.Header("Content-Type", "application/json")
			c.String(200, string(bytes))
			c.Abort()
			return
		}

		// 存在的静态文件
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

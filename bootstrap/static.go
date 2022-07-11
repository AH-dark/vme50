package bootstrap

import (
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/gin-contrib/static"
	"io/fs"
	"net/http"
)

var StaticFS static.ServeFileSystem

type GinFS struct {
	FS http.FileSystem
}

// Open 打开文件
func (b *GinFS) Open(name string) (http.File, error) {
	return b.FS.Open(name)
}

// Exists 文件是否存在
func (b *GinFS) Exists(prefix string, filepath string) bool {
	if _, err := b.FS.Open(filepath); err != nil {
		return false
	}
	return true
}

// InitStatic 初始化静态文件
func InitStatic(static fs.FS) {
	embedFS, err := fs.Sub(static, "build")
	if err != nil {
		utils.Log().Panic("无法初始化静态资源, %s", err)
	}

	StaticFS = &GinFS{
		FS: http.FS(embedFS),
	}
}

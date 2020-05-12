package guard

import (
	"github.com/HongJaison/filemanager/modules/constant"
	errors "github.com/HongJaison/filemanager/modules/error"
	"github.com/HongJaison/filemanager/modules/permission"
	"github.com/HongJaison/filemanager/modules/root"
	"github.com/HongJaison/filemanager/modules/util"
	"github.com/HongJaison/go-admin/context"
	"github.com/HongJaison/go-admin/modules/db"
	"net/url"
	"path/filepath"
	"strings"
)

type Guardian struct {
	conn        db.Connection
	roots       root.Roots
	permissions permission.Permission
}

func New(r root.Roots, c db.Connection, p permission.Permission) *Guardian {
	return &Guardian{
		roots:       r,
		conn:        c,
		permissions: p,
	}
}

const (
	filesParamKey     = "files_param"
	uploadParamKey    = "upload_param"
	createDirParamKey = "create_dir_param"
	deleteParamKey    = "delete_param"
	renameParamKey    = "rename_param"
	previewParamKey   = "preview_param"
)

type Base struct {
	Path     string
	Prefix   string
	FullPath string
	Error    error
}

func (g *Guardian) GetPrefix(ctx *context.Context) string {
	prefix := ctx.Query(constant.PrefixKey)
	if prefix == "" {
		return "def"
	}
	return prefix
}

func (g *Guardian) getPaths(ctx *context.Context) (string, string, error) {
	var (
		err error

		relativePath, _ = url.QueryUnescape(ctx.Query("path"))
		path            = filepath.Join(g.roots.GetPathFromPrefix(ctx), relativePath)
	)
	if !strings.Contains(path, g.roots.GetPathFromPrefix(ctx)) {
		err = errors.DirIsNotExist
	}

	if !util.FileExist(path) {
		err = errors.DirIsNotExist
	}

	return relativePath, path, err
}

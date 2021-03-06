package guard

import (
	errors "github.com/HongJaison/filemanager/modules/error"
	"github.com/HongJaison/filemanager/modules/util"
	"github.com/HongJaison/go-admin/context"
	"path/filepath"
)

type RenameParam struct {
	Src    string
	Dist   string
	Error  error
	Prefix string
}

func (g *Guardian) Rename(ctx *context.Context) {

	distName := ctx.FormValue("name")
	src := ctx.FormValue("path")

	if src == "" || src == "/" || distName == "" || distName == "/" {
		ctx.SetUserValue(renameParamKey, &RenameParam{Error: errors.EmptyName})
		ctx.Next()
		return
	}

	if filepath.Ext(distName) == "" && util.IsFile(g.roots.GetPathFromPrefix(ctx)+src) {
		distName += filepath.Ext(src)
	}

	ctx.SetUserValue(renameParamKey, &RenameParam{
		Src:    g.roots.GetPathFromPrefix(ctx) + src,
		Dist:   g.roots.GetPathFromPrefix(ctx) + filepath.Dir(src) + "/" + distName,
		Prefix: g.GetPrefix(ctx),
	})
	ctx.Next()
}

func GetRenameParam(ctx *context.Context) *RenameParam {
	return ctx.UserValue[renameParamKey].(*RenameParam)
}

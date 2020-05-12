package controller

import (
	"github.com/HongJaison/filemanager/guard"
	"github.com/HongJaison/filemanager/models"
	"github.com/HongJaison/filemanager/modules/util"
	"github.com/HongJaison/go-admin/context"
	"io/ioutil"
	"mime"
	"path/filepath"
	"strings"
)

func (h *Handler) ListFiles(ctx *context.Context) {

	var (
		param      = guard.GetFilesParam(ctx)
		filesOfDir = make(models.Files, 0)
	)

	if param.Error != nil {
		h.table(ctx, filesOfDir, param.Error)
		return
	}

	fileInfos, err := ioutil.ReadDir(filepath.FromSlash(param.FullPath))

	if err != nil {
		h.table(ctx, filesOfDir, err)
		return
	}

	for _, fileInfo := range fileInfos {

		if util.IsHiddenFile(fileInfo.Name()) {
			continue
		}

		file := models.File{
			IsDirectory:  fileInfo.IsDir(),
			Name:         fileInfo.Name(),
			Size:         int(fileInfo.Size()),
			Extension:    strings.TrimLeft(filepath.Ext(fileInfo.Name()), "."),
			Path:         filepath.Join(param.Path, fileInfo.Name()),
			Mime:         mime.TypeByExtension(filepath.Ext(fileInfo.Name())),
			LastModified: fileInfo.ModTime().Unix(),
		}

		filesOfDir = append(filesOfDir, file)
	}

	h.table(ctx, filesOfDir, nil)
	return
}

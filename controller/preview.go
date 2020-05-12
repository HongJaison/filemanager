package controller

import (
	"github.com/HongJaison/filemanager/guard"
	"github.com/HongJaison/filemanager/previewer"
	"github.com/HongJaison/go-admin/context"
)

func (h *Handler) Preview(ctx *context.Context) {
	param := guard.GetPreviewParam(ctx)
	if param.Error != nil {
		h.preview(ctx, "", param.Path, param.FullPath, param.Error)
		return
	}
	content, err := previewer.Preview(param.FullPath)
	h.preview(ctx, content, param.Path, param.FullPath, err)
}

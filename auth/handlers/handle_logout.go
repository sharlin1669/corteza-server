package handlers

import (
	"github.com/cortezaproject/corteza-server/auth/request"
	"github.com/markbates/goth/gothic"
)

func (h *AuthHandlers) logoutProc(req *request.AuthReq) (err error) {
	req.Session.Options.MaxAge = -1
	if err = req.Session.Save(req.Request, req.Response); err != nil {
		return
	}

	if err = gothic.Logout(req.Response, req.Request); err != nil {
		return
	}

	// Prevent these two to be rendered by in the template
	req.AuthUser = nil
	req.Client = nil
	h.Log.Info("logout successful")

	req.Template = TmplLogout

	req.Data["link"] = GetLinks().Login

	if bl := req.Request.FormValue("back"); bl != "" {
		req.Data["link"] = sanitizeLink(bl)
	}

	return
}

package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
)

func (h *handler) CreateService(r *restful.Request, w *restful.Response) {
	req := service.NewCreateServiceRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	// 补充用户信息
	tk := token.GetTokenFromRequest(r)
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace

	set, err := h.service.CreateService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (u *handler) QueryService(r *restful.Request, w *restful.Response) {
	req := service.NewQueryServiceRequestFromHTTP(r.Request)
	set, err := h.service.QueryService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (u *handler) DescribeService(r *restful.Request, w *restful.Response) {
	req := service.NewDescribeServiceRequest(r.PathParameter("id"))
	ins, err := h.service.DescribeService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (u *handler) UpdateService(r *restful.Request, w *restful.Response) {
	req := service.NewPutServiceRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (u *handler) PatchService(r *restful.Request, w *restful.Response) {
	req := service.NewPatchServiceRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (u *handler) DeleteService(r *restful.Request, w *restful.Response) {
	req := service.NewDeleteServiceRequestWithID(r.PathParameter("id"))
	set, err := h.service.DeleteService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

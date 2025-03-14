package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/user"
)

func (s *impl) CreatePolicy(ctx context.Context, req *policy.CreatePolicyRequest) (
	*policy.Policy, error) {
	ins, err := policy.New(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u, err := s.CheckDependence(ctx, ins)
	if err != nil {
		return nil, err
	}
	s.log.Debugf("user: %s", u.Spec.Username)

	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted policy(%s) document error, %s",
			ins.Meta.Id, err)
	}

	return ins, nil
}

// CheckDependence todo
func (i *impl) CheckDependence(ctx context.Context, p *policy.Policy) (*user.User, error) {
	account, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestByName(p.Spec.Username))
	if err != nil {
		return nil, fmt.Errorf("check user error, %s", err)
	}

	_, err = i.role.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(p.Spec.RoleId))
	if err != nil {
		return nil, fmt.Errorf("check role error, %s", err)
	}

	if !p.IsAllNamespace() {
		_, err = i.namespace.DescribeNamespace(ctx, namespace.NewDescriptNamespaceRequest(p.Spec.Domain, p.Spec.Namespace))
		if err != nil {
			return nil, fmt.Errorf("check namespace error, %s", err)
		}
	}

	return account, nil
}

func (s *impl) QueryPolicy(ctx context.Context, req *policy.QueryPolicyRequest) (
	*policy.PolicySet, error) {
	r, err := newQueryPolicyRequest(req)
	if err != nil {
		return nil, err
	}

	if req.Namespace != "" {
		_, err := s.namespace.DescribeNamespace(ctx, namespace.NewDescriptNamespaceRequest(req.Domain, req.Namespace))
		if err != nil {
			return nil, err
		}
	}

	s.log.Debugf("query policy filter: %s", r.FindFilter())
	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find policy error, error is %s", err)
	}

	set := policy.NewPolicySet()
	// 循环
	for resp.Next(ctx) {
		ins := policy.NewDefaultPolicy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode policy error, error is %s", err)
		}

		// 补充关联的角色信息
		if req.WithRole {
			descRole := role.NewDescribeRoleRequestWithID(ins.Spec.RoleId)
			ins.Role, err = s.role.DescribeRole(ctx, descRole)
			if err != nil {
				return nil, err
			}
		}

		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get policy count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *impl) DescribePolicy(ctx context.Context, req *policy.DescribePolicyRequest) (
	*policy.Policy, error) {
	r, err := newDescribePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	ins := policy.NewDefaultPolicy()
	s.log.Debugf("describe policy filter: %s", r.FindFilter())
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("policy %s not found", req)
		}

		return nil, exception.NewInternalServerError("find policy %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *impl) DeletePolicy(ctx context.Context, req *policy.DeletePolicyRequest) (*policy.Policy, error) {
	descReq := policy.NewDescriptPolicyRequest()
	descReq.Id = req.Id
	p, err := s.DescribePolicy(ctx, descReq)
	if err != nil {
		return nil, err
	}

	r, err := newDeletePolicyRequest(req)
	if err != nil {
		return nil, err
	}

	_, err = s.col.DeleteOne(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("delete policy(%s) error, %s", req.Id, err)
	}

	return p, nil
}

func (s *impl) CheckPermission(ctx context.Context, req *policy.CheckPermissionRequest) (
	*role.Permission, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	// 判断是否是空间所有者
	ns, err := s.namespace.DescribeNamespace(ctx, namespace.NewDescriptNamespaceRequest(req.Domain, req.Namespace))
	if err != nil {
		return nil, err
	}
	if ns.IsOwner(req.Username) {
		return role.OwnerAdminPermssion(), nil
	}

	// 查询用户权限策略
	pReq := policy.NewQueryPolicyRequest()
	pReq.Username = req.Username
	pReq.Domain = req.Domain
	pReq.Namespace = req.Namespace
	pReq.WithRole = true
	ps, err := s.QueryPolicy(ctx, pReq)
	if err != nil {
		return nil, err
	}

	// 无用户相关权限策略设置
	if ps.Len() == 0 {
		return nil, exception.NewPermissionDeny("no permission")
	}

	// 查询用户需要鉴权的功能
	fn := endpoint.NewDescribeEndpointRequestWithID(endpoint.GenHashID(req.ServiceId, req.Path))
	ep, err := s.endpoint.DescribeEndpoint(ctx, fn)
	if err != nil {
		return nil, err
	}

	// 判断改功能是否需要鉴权
	if !ep.Entry.PermissionEnable {
		return role.NewSkipPermission("endpoint not enable permission check, allow all access"), nil
	}

	// 判断策略是否允许
	var perm *role.Permission
	for i := range ps.Items {
		p := ps.Items[i]
		permOk, ok, err := p.Role.HasPermission(ep)
		if err != nil {
			return nil, err
		}
		if ok {
			perm = permOk
			perm.Scope = p.Spec.Scope
			s.log.Debugf("check roles %s has permission access endpoint [%s]", p.Role.Spec.Name, ep.Entry)
		}
	}

	if perm == nil {
		return nil, exception.NewPermissionDeny("in namespace %s, role %s has no permission access endpoint: %s",
			req.Namespace,
			ps.RoleNames(),
			ep.Entry.Path,
		)
	}

	return perm, nil
}

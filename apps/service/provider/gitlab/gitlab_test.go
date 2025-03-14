package gitlab_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
	"github.com/infraboard/mcenter/test/tools"
)

var (
	v4  *gitlab.GitlabV4
	ctx = context.Background()

	ProjectID string = "29032549"
)

func TestListProject(t *testing.T) {
	req := gitlab.NewListProjectRequest()
	req.PageSize = 20
	req.PageNumer = 1
	req.Keywords = ""
	set, err := v4.Project().ListProjects(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestListProjectLanguage(t *testing.T) {
	set, err := v4.Project().ListProjectLanguage(ctx, ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set.Primary())
}

func TestListProjectBranch(t *testing.T) {
	req := gitlab.NewListProjectBranchRequest()
	req.ProjectId = ProjectID
	req.PageSize = 2
	req.Keywords = ""
	set, err := v4.Project().ListProjectBranch(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestGetProjectBranch(t *testing.T) {
	req := gitlab.NewGetProjectBranchRequest()
	req.ProjectId = ProjectID
	req.Branch = "main"
	set, err := v4.Project().GetProjectBranch(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestAddProjectHook(t *testing.T) {
	hook := gitlab.NewGitLabWebHook("9999")
	req := gitlab.NewAddProjectHookRequest(ProjectID, hook)

	resp, err := v4.Project().AddProjectHook(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestDeleteProjectHook(t *testing.T) {
	req := gitlab.NewDeleteProjectHookReqeust(ProjectID, "19001857")
	err := v4.Project().DeleteProjectHook(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	tools.DevelopmentSetup()
	conf, err := gitlab.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	v4 = gitlab.NewGitlabV4(conf)
}

package routers

import (
	"github.com/alanbui/train-ticket/internal/routers/manager"
	"github.com/alanbui/train-ticket/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}
